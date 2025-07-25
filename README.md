# IP Lookup Tool

A modern, responsive IP lookup application with a Go backend and Vue.js frontend, featuring interactive maps and comprehensive IP information display.

## Features

- **IP Geolocation**: Get detailed information about any IP address including location, ISP, and network details
- **Interactive Map**: Visual location display using Leaflet maps with OpenStreetMap tiles
- **Current IP Detection**: Automatically detect and lookup your current IP address
- **Mobile Responsive**: Fully optimized for mobile and desktop devices
- **Single Binary**: Complete application bundled into a single executable
- **Subpath Compatible**: Works correctly when deployed on subpaths
- **Modern UI**: Clean, professional interface built with Vue.js and Tailwind CSS

## API Endpoints

### Get Current IP Information
```
GET /api/my-ip
```
Returns information about the client's current IP address.

### Lookup Specific IP
```
GET /api/lookup/{ip}
```
Returns information about the specified IP address.

**Example Response:**
```json
{
    "query": "8.8.8.8",
    "status": "success",
    "continent": "North America",
    "continentCode": "NA",
    "country": "United States",
    "countryCode": "US",
    "region": "VA",
    "regionName": "Virginia",
    "city": "Ashburn",
    "district": "",
    "zip": "20149",
    "lat": 39.03,
    "lon": -77.5,
    "timezone": "America/New_York",
    "offset": -14400,
    "currency": "USD",
    "isp": "Google LLC",
    "org": "Google Public DNS",
    "as": "AS15169 Google LLC",
    "asname": "GOOGLE",
    "mobile": false,
    "proxy": false,
    "hosting": true
}
```

## Quick Start

### Running the Application

1. **Download the binary** (or build from source)
2. **Run the application:**
   ```bash
   ./ip-lookup-app
   ```
3. **Open your browser** and navigate to `http://localhost:8080`

### Custom Port

Set a custom port using the `PORT` environment variable:
```bash
PORT=3000 ./ip-lookup-app
```

## Building from Source

### Prerequisites

- Go 1.21 or later
- Node.js 20 or later
- npm or yarn

### Build Steps

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd ip-lookup-app
   ```

2. **Build the frontend:**
   ```bash
   cd frontend
   npm install
   npm run build
   cd ..
   ```

3. **Build the Go binary:**
   ```bash
   go build -o ip-lookup-app main.go
   ```

## Architecture

### Backend (Go)
- **HTTP Server**: Serves both API endpoints and static frontend files
- **IP Geolocation**: Uses ip-api.com for IP information lookup
- **CORS Support**: Enables cross-origin requests
- **Embedded Assets**: Frontend files are embedded using Go's embed directive
- **Subpath Routing**: Supports deployment on subpaths with SPA routing

### Frontend (Vue.js + TypeScript)
- **Vue 3**: Modern reactive framework with Composition API
- **TypeScript**: Type-safe development
- **Tailwind CSS**: Utility-first CSS framework for responsive design
- **Leaflet Maps**: Interactive maps with OpenStreetMap tiles
- **Responsive Design**: Mobile-first approach with breakpoint optimization

## Deployment

### Standalone Binary
The application is designed to be deployed as a single binary:

```bash
# Copy the binary to your server
scp ip-lookup-app user@server:/path/to/deployment/

# Run on the server
./ip-lookup-app
```

### Docker (Optional)
Create a `Dockerfile`:
```dockerfile
FROM scratch
COPY ip-lookup-app /
EXPOSE 8080
CMD ["/ip-lookup-app"]
```

### Reverse Proxy
For production deployment behind a reverse proxy (nginx, Apache, etc.), ensure:
- Proper forwarding of client IP headers (`X-Forwarded-For`, `X-Real-IP`)
- Static asset caching for optimal performance
- HTTPS termination if needed

## Configuration

### Environment Variables

- `PORT`: Server port (default: 8080)

### Customization

The application can be customized by modifying:
- **Frontend styling**: Edit Tailwind classes in Vue components
- **Map provider**: Change tile layer in `MapView.vue`
- **API provider**: Modify IP lookup service in `main.go`

## Browser Support

- Chrome/Chromium 90+
- Firefox 88+
- Safari 14+
- Edge 90+

## License

This project is open source and available under the MIT License.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## Support

For issues and questions, please create an issue in the repository or contact the development team.

