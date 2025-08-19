package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type InfoResponse struct {
	IPInfo     IPInfo
	HeaderInfo map[string][]string
}

// IPInfo represents the structure of IP lookup response
type IPInfo struct {
	Query         string  `json:"query"`
	Status        string  `json:"status"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Region        string  `json:"region"`
	RegionName    string  `json:"regionName"`
	City          string  `json:"city"`
	District      string  `json:"district"`
	Zip           string  `json:"zip"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	Timezone      string  `json:"timezone"`
	Offset        int     `json:"offset"`
	Currency      string  `json:"currency"`
	ISP           string  `json:"isp"`
	Org           string  `json:"org"`
	AS            string  `json:"as"`
	ASName        string  `json:"asname"`
	Mobile        bool    `json:"mobile"`
	Proxy         bool    `json:"proxy"`
	Hosting       bool    `json:"hosting"`
}

// ErrorResponse represents error response structure
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Set up routes
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/api/lookup/", handleIPLookup)
	http.HandleFunc("/api/my-ip", handleMyIP)

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

// CORS middleware
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// Handle root and static files
func handleRoot(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == "OPTIONS" {
		return
	}

	// Try to serve from embedded or on-disk frontendFS
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" {
		path = "index.html"
	}

	// API routes should 404 here
	if strings.HasPrefix(path, "api/") {
		http.NotFound(w, r)
		return
	}

	log.Println(r.URL.Path)
	if strings.Contains(r.Header.Get("Accept"), "html") {
		// Manual user request
		writeHtmlPage(path, w)
		return
	} else {
		// Probably normal curl request, so let return `/api/my-ip`
		ip := getClientIP(r)
		ipInfo, err := getIPInfo(ip)
		if err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(ipInfo.Query))
		return
	}

}

func writeHtmlPage(path string, w http.ResponseWriter) {
	fullPath := filepath.Join("frontend/dist", path)
	data, err := fs.ReadFile(frontendFS, fullPath)
	if err != nil {
		// serve index.html for SPA routing
		data, err = fs.ReadFile(frontendFS, "frontend/dist/index.html")
		if err != nil {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>IP Lookup API</title>
</head>
<body>
    <h1>IP Lookup API</h1>
    <p>API Endpoints:</p>
    <ul>
        <li><a href="/api/my-ip">/api/my-ip</a> - Get your current IP info</li>
        <li>/api/lookup/{ip} - Get info for specific IP</li>
    </ul>
</body>
</html>`)
			return
		}
		w.Header().Set("Content-Type", "text/html")
	} else {
		// Set content-type by extension
		switch ext := filepath.Ext(path); ext {
		case ".html":
			w.Header().Set("Content-Type", "text/html")
		case ".css":
			w.Header().Set("Content-Type", "text/css")
		case ".js":
			w.Header().Set("Content-Type", "application/javascript")
		case ".json":
			w.Header().Set("Content-Type", "application/json")
		}
	}

	w.Write(data)
}

// Handle IP lookup requests
func handleIPLookup(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == "OPTIONS" {
		return
	}

	ip := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/api/lookup/"))
	if ip == "" {
		sendErrorResponse(w, "IP address is required", http.StatusBadRequest)
		return
	}
	if net.ParseIP(ip) == nil {
		sendErrorResponse(w, "Invalid IP address format", http.StatusBadRequest)
		return
	}

	ipInfo, err := getIPInfo(ip)
	if err != nil {
		log.Printf("Error getting IP info for %s: %v", ip, err)
		sendErrorResponse(w, "Failed to lookup IP information", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	ir := InfoResponse{
		IPInfo:     *ipInfo,
		HeaderInfo: r.Header,
	}

	json.NewEncoder(w).Encode(ir)
}

// Handle current IP lookup
func handleMyIP(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == "OPTIONS" {
		return
	}

	clientIP := getClientIP(r)
	ipInfo, err := getIPInfo(clientIP)
	if err != nil {
		log.Printf("Error getting IP info for %s: %v", clientIP, err)
		sendErrorResponse(w, "Failed to lookup IP information", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	ir := InfoResponse{
		IPInfo:     *ipInfo,
		HeaderInfo: r.Header,
	}

	json.NewEncoder(w).Encode(ir)
}

// Get client IP address
func getClientIP(r *http.Request) string {
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		return strings.TrimSpace(strings.Split(forwarded, ",")[0])
	}
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if strings.Contains(ip, "::1") || strings.Contains(ip, "localhost") {
		log.Printf("Have got local ip of [%s]", ip)
		return ""
	}
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

// Get IP information from external API
func getIPInfo(ip string) (*IPInfo, error) {
	url := fmt.Sprintf(
		"http://ip-api.com/json/%s?fields=status,message,continent,continentCode,"+
			"country,countryCode,region,regionName,city,district,zip,lat,lon,"+
			"timezone,offset,currency,isp,org,as,asname,mobile,proxy,hosting,query",
		ip,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var ipInfo IPInfo
	if err := json.Unmarshal(body, &ipInfo); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}
	if ipInfo.Status != "success" {
		return nil, fmt.Errorf("API returned error status: %s", ipInfo.Status)
	}

	return &ipInfo, nil
}

// Send error response
func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Status:  "error",
		Message: message,
	})
}
