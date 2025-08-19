<template>
    <div class="min-h-screen bg-slate-50 dark:bg-slate-950">
        <!-- Main content -->
        <div class="container mx-auto px-4 py-6 max-w-8xl">
            <!-- Header -->
            <header class="text-center mb-8 flex items-center gap-10">
                <div
                    class="inline-flex items-center justify-center w-12 h-12 bg-slate-900 dark:bg-slate-100 rounded-lg mb-4"
                >
                    <svg
                        class="w-6 h-6 text-slate-50 dark:text-slate-900"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9v-9m0-9v9"
                        ></path>
                    </svg>
                </div>
                <h1
                    class="text-3xl font-bold text-slate-900 dark:text-slate-100 mb-2"
                >
                    IP Lookup
                </h1>
                <p class="text-sm text-slate-600 dark:text-slate-400">
                    Discover your digital footprint with comprehensive network
                    insights
                </p>
            </header>

            <!-- Loading state -->
            <div v-if="loading" class="max-w-2xl mx-auto">
                <div
                    class="bg-white dark:bg-slate-900 rounded-lg border border-slate-200 dark:border-slate-800 p-6 shadow-sm"
                >
                    <div class="flex items-center space-x-3">
                        <div
                            class="w-5 h-5 border-2 border-slate-300 border-t-slate-900 dark:border-slate-600 dark:border-t-slate-100 rounded-full animate-spin"
                        ></div>
                        <div>
                            <h3
                                class="text-sm font-medium text-slate-900 dark:text-slate-100"
                            >
                                Analyzing Connection
                            </h3>
                            <p
                                class="text-xs text-slate-600 dark:text-slate-400"
                            >
                                Gathering network information...
                            </p>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Error state -->
            <div v-if="error" class="max-w-2xl mx-auto mb-6">
                <div
                    class="bg-red-50 dark:bg-red-950/50 border border-red-200 dark:border-red-800 rounded-lg p-4"
                >
                    <div class="flex items-center">
                        <div class="flex-shrink-0">
                            <svg
                                class="h-5 w-5 text-red-400"
                                viewBox="0 0 20 20"
                                fill="currentColor"
                            >
                                <path
                                    fill-rule="evenodd"
                                    d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                                    clip-rule="evenodd"
                                />
                            </svg>
                        </div>
                        <div class="ml-3">
                            <h3
                                class="text-sm font-medium text-red-800 dark:text-red-200"
                            >
                                Connection Error
                            </h3>
                            <p class="text-sm text-red-700 dark:text-red-300">
                                {{ error }}
                            </p>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Results -->
            <div v-if="ipInfo && !loading" class="space-y-6">
                <!-- IP Address Card with Copy Functionality -->
                <div
                    class="bg-white dark:bg-slate-900 rounded-lg border border-slate-200 dark:border-slate-800 p-6 shadow-sm"
                >
                    <div class="flex items-center justify-center gap-5 mb-4">
                        <h2
                            class="text-lg font-semibold text-slate-900 dark:text-slate-100"
                        >
                            Your IP Address
                        </h2>
                        <div class="flex items-center space-x-2">
                            <div
                                class="w-2 h-2 bg-green-500 rounded-full"
                            ></div>
                            <span
                                class="text-xs text-slate-600 dark:text-slate-400 uppercase tracking-wide"
                                >{{ ipInfo.status }}</span
                            >
                        </div>
                    </div>

                    <div
                        class="bg-slate-50 dark:bg-slate-800 rounded-md p-3 border border-slate-200 dark:border-slate-700 flex items-center gap-2 flex-col"
                    >
                        <div
                            class="font-mono text-lg font-semibold text-slate-900 dark:text-slate-100"
                        >
                            {{ ipInfo.query }}
                        </div>
                        <div>
                            <button
                                @click="copyToClipboard(ipInfo.query)"
                                class="inline-flex items-center px-3 py-1.5 text-xs font-medium text-slate-700 dark:text-slate-300 bg-white dark:bg-slate-700 border border-slate-300 dark:border-slate-600 rounded-md hover:bg-slate-50 dark:hover:bg-slate-600 focus:outline-none focus:ring-2 focus:ring-slate-500 focus:ring-offset-2 dark:focus:ring-offset-slate-900 transition-colors"
                            >
                                <svg
                                    v-if="!copied"
                                    class="w-3 h-3 mr-1.5"
                                    fill="none"
                                    stroke="currentColor"
                                    viewBox="0 0 24 24"
                                >
                                    <path
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        stroke-width="2"
                                        d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"
                                    ></path>
                                </svg>
                                <svg
                                    v-else
                                    class="w-3 h-3 mr-1.5 text-green-600"
                                    fill="none"
                                    stroke="currentColor"
                                    viewBox="0 0 24 24"
                                >
                                    <path
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        stroke-width="2"
                                        d="M5 13l4 4L19 7"
                                    ></path>
                                </svg>
                                {{ copied ? "Copied!" : "Copy" }}
                            </button>
                        </div>
                    </div>
                </div>

                <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
                    <LocationInformation :ip-info="ipInfo" />

                    <NetworkInformation :ip-info="ipInfo" />

                    <!-- Properties -->
                    <PropertyInformation :ip-info="ipInfo" />
                    <!-- Map -->
                    <div
                        class="bg-white dark:bg-slate-900 rounded-lg border border-slate-200 dark:border-slate-800 p-6 shadow-sm"
                    >
                        <div class="flex items-center mb-4">
                            <div
                                class="w-8 h-8 bg-slate-100 dark:bg-slate-800 rounded-md flex items-center justify-center mr-3"
                            >
                                <svg
                                    class="w-4 h-4 text-slate-600 dark:text-slate-400"
                                    fill="none"
                                    stroke="currentColor"
                                    viewBox="0 0 24 24"
                                >
                                    <path
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        stroke-width="2"
                                        d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-1.447-.894L15 4m0 13V4m0 0L9 7"
                                    ></path>
                                </svg>
                            </div>
                            <h3
                                class="text-base font-semibold text-slate-900 dark:text-slate-100"
                            >
                                Geographic Location
                            </h3>
                        </div>

                        <MapView
                            v-if="ipInfo"
                            :lat="ipInfo.lat"
                            :lon="ipInfo.lon"
                            :city="ipInfo.city"
                            :country="ipInfo.country"
                            :query="ipInfo.query"
                        />
                    </div>
                    <HeaderInformation :header-info="headerInfo" />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import type { HeaderInfo, InfoResponse, IPInfo } from "./types/ip";
import MapView from "./components/MapView.vue";
import LocationInformation from "./components/LocationInformation.vue";
import NetworkInformation from "./components/NetworkInformation.vue";
import PropertyInformation from "./components/PropertyInformation.vue";
import HeaderInformation from "./components/HeaderInformation.vue";

const ipInfo = ref<IPInfo | null>(null);
const headerInfo = ref<HeaderInfo | null>(null);
const loading = ref(true);
const error = ref("");
const copied = ref(false);

const copyToClipboard = async (text: string) => {
    try {
        await navigator.clipboard.writeText(text);
        copied.value = true;
        setTimeout(() => {
            copied.value = false;
        }, 2000);
    } catch (err) {
        console.error("Failed to copy text: ", err);
    }
};

const lookupCurrentIP = async () => {
    loading.value = true;
    error.value = "";
    ipInfo.value = null;

    try {
        const response = await fetch("/api/my-ip");
        const data: InfoResponse = await response.json();

        if (!response.ok) {
            // @ts-ignore
            throw new Error(data.message || "Failed to lookup IP");
        }

        if (data.IPInfo.status === "success") {
            ipInfo.value = data.IPInfo;
            headerInfo.value = data.HeaderInfo;
        } else {
            // @ts-ignore
            throw new Error(data.message || "IP lookup failed");
        }
    } catch (err) {
        error.value =
            err instanceof Error ? err.message : "An unknown error occurred";
    } finally {
        loading.value = false;
    }
};

// Automatically lookup current IP when component mounts
onMounted(() => {
    lookupCurrentIP();
});
</script>

<style scoped>
/* Custom scrollbar for webkit browsers */
::-webkit-scrollbar {
    width: 6px;
}

::-webkit-scrollbar-track {
    background: transparent;
}

::-webkit-scrollbar-thumb {
    background: rgb(148 163 184);
    border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
    background: rgb(100 116 139);
}

.dark ::-webkit-scrollbar-thumb {
    background: rgb(71 85 105);
}

.dark ::-webkit-scrollbar-thumb:hover {
    background: rgb(100 116 139);
}
</style>
