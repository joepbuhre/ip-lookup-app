<template>
    <!-- Location Information -->
    <div
        v-if="ipInfo"
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
                        d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"
                    ></path>
                </svg>
            </div>
            <h3
                class="text-base font-semibold text-slate-900 dark:text-slate-100"
            >
                Location
            </h3>
        </div>

        <div class="space-y-3">
            <InfoItem
                label="Country"
                :value="`${ipInfo.country} ${getCountryFlag(
                    ipInfo.countryCode
                )}`"
            />
            <InfoItem label="Region" :value="ipInfo.regionName" />
            <InfoItem label="City" :value="ipInfo.city" />
            <InfoItem label="ZIP Code" :value="ipInfo.zip" />
            <InfoItem label="Timezone" :value="ipInfo.timezone" />
            <InfoItem
                label="Coordinates"
                :value="`${ipInfo.lat.toFixed(4)}, ${ipInfo.lon.toFixed(4)}`"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import { IPInfo } from "@/types/ip";
import { getCountryFlag } from "@/utils/utils";
import { defineProps } from "vue";
import InfoItem from "./InfoItem.vue";

const props = defineProps<{
    ipInfo: IPInfo | null;
}>();
</script>
