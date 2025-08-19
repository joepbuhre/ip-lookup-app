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
                        d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
                    ></path>
                </svg>
            </div>
            <h3
                class="text-base font-semibold text-slate-900 dark:text-slate-100"
            >
                Properties
            </h3>
        </div>

        <div class="space-y-2">
            <PropertyBadge :active="ipInfo.mobile" label="Mobile Connection" />
            <PropertyBadge :active="ipInfo.proxy" label="Proxy Server" />
            <PropertyBadge :active="ipInfo.hosting" label="Hosting Service" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { IPInfo } from "@/types/ip";
import { getCountryFlag } from "@/utils/utils";
import { defineProps } from "vue";
import PropertyBadge from "./PropertyBadge.vue";

const props = defineProps<{
    ipInfo: IPInfo | null;
}>();
</script>
