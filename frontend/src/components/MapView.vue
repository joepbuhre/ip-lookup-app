<template>
    <div class="w-full">
        <div
            ref="mapContainer"
            class="rounded-md h-64 w-full overflow-hidden border border-slate-200 dark:border-slate-700"
            style="min-height: 356px"
        ></div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, onUnmounted } from "vue";
import L from "leaflet";
import "leaflet/dist/leaflet.css";

interface Props {
    lat: number;
    lon: number;
    city: string;
    country: string;
    query: string;
}

const props = defineProps<Props>();

const mapContainer = ref<HTMLElement>();
let map: L.Map | null = null;
let marker: L.Marker | null = null;

// Fix Leaflet default markers
delete (L.Icon.Default.prototype as any)._getIconUrl;
L.Icon.Default.mergeOptions({
    iconRetinaUrl:
        "https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.9.4/images/marker-icon-2x.png",
    iconUrl:
        "https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.9.4/images/marker-icon.png",
    shadowUrl:
        "https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.9.4/images/marker-shadow.png",
});

const initMap = () => {
    if (!mapContainer.value) return;

    // Create map
    map = L.map(mapContainer.value, {
        zoomControl: false,
    }).setView([props.lat, props.lon], 10);

    // Add custom zoom control
    L.control
        .zoom({
            position: "bottomright",
        })
        .addTo(map);

    // Use OpenStreetMap tiles
    L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
        attribution:
            '© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
        maxZoom: 19,
    }).addTo(map);

    // Create custom marker icon
    const customIcon = L.divIcon({
        html: `
      <div class="relative">
        <div class="w-4 h-4 bg-slate-900 rounded-full border-2 border-white shadow-lg"></div>
        <div class="absolute -top-1 -left-1 w-6 h-6 bg-slate-900/20 rounded-full animate-ping"></div>
      </div>
    `,
        className: "custom-marker",
        iconSize: [16, 16],
        iconAnchor: [8, 8],
    });

    // Add marker with custom icon
    marker = L.marker([props.lat, props.lon], { icon: customIcon }).addTo(map);

    // Create popup content
    const popupContent = `
    <div class="text-center p-2 bg-white rounded-lg border border-slate-200 shadow-sm">
      <div class="text-sm font-semibold text-slate-900 mb-1">${
          props.query
      }</div>
      <div class="text-slate-600 text-xs">${props.city}, ${props.country}</div>
      <div class="text-slate-500 text-xs mt-1">${props.lat.toFixed(
          4
      )}, ${props.lon.toFixed(4)}</div>
    </div>
  `;

    marker
        .bindPopup(popupContent, {
            className: "custom-popup",
        })
        .openPopup();

    // Add a subtle circle around the location
    L.circle([props.lat, props.lon], {
        color: "#1e293b",
        fillColor: "#1e293b",
        fillOpacity: 0.1,
        radius: 5000,
        weight: 1,
    }).addTo(map);
};

const updateMap = () => {
    if (!map || !marker) return;

    const newLatLng = L.latLng(props.lat, props.lon);

    // Update map view with smooth animation
    map.flyTo(newLatLng, 10, {
        duration: 1.5,
    });

    // Update marker position
    marker.setLatLng(newLatLng);

    // Update popup content
    const popupContent = `
    <div class="text-center p-2 bg-white rounded-lg border border-slate-200 shadow-sm">
      <div class="text-sm font-semibold text-slate-900 mb-1">${
          props.query
      }</div>
      <div class="text-slate-600 text-xs">${props.city}, ${props.country}</div>
      <div class="text-slate-500 text-xs mt-1">${props.lat.toFixed(
          4
      )}, ${props.lon.toFixed(4)}</div>
    </div>
  `;

    marker
        .bindPopup(popupContent, {
            className: "custom-popup",
        })
        .openPopup();
};

onMounted(() => {
    // Small delay to ensure DOM is ready
    setTimeout(initMap, 100);
});

onUnmounted(() => {
    if (map) {
        map.remove();
        map = null;
        marker = null;
    }
});

// Watch for prop changes and update map
watch(
    [
        () => props.lat,
        () => props.lon,
        () => props.city,
        () => props.country,
        () => props.query,
    ],
    () => {
        if (map) {
            updateMap();
        }
    }
);
</script>

<style scoped>
/* Custom map styling */
:deep(.leaflet-container) {
    background: #f8fafc;
    border-radius: 0.375rem;
}

:deep(.leaflet-control-zoom) {
    border: none !important;
    border-radius: 0.375rem !important;
    box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06) !important;
    background: white !important;
}

:deep(.leaflet-control-zoom a) {
    background: white !important;
    color: #1e293b !important;
    border: 1px solid #e2e8f0 !important;
    border-radius: 0.375rem !important;
    transition: all 0.2s ease !important;
    width: 28px !important;
    height: 28px !important;
    line-height: 26px !important;
}

:deep(.leaflet-control-zoom a:hover) {
    background: #f8fafc !important;
    border-color: #cbd5e1 !important;
}

:deep(.custom-popup .leaflet-popup-content-wrapper) {
    background: white !important;
    border-radius: 0.5rem !important;
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1),
        0 2px 4px -1px rgba(0, 0, 0, 0.06) !important;
    border: 1px solid #e2e8f0 !important;
    padding: 0 !important;
}

:deep(.custom-popup .leaflet-popup-tip) {
    background: white !important;
    border: 1px solid #e2e8f0 !important;
}

:deep(.leaflet-popup-close-button) {
    color: #64748b !important;
    font-size: 16px !important;
    padding: 4px 8px !important;
}

:deep(.custom-marker) {
    background: transparent !important;
    border: none !important;
}
</style>
