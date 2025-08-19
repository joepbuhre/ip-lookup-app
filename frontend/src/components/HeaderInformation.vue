<template>
    <!-- Location Information -->
    <div
        v-if="headerInfo"
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
                        d="M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2"
                    ></path>
                </svg>
            </div>
            <h3
                class="text-base font-semibold text-slate-900 dark:text-slate-100"
            >
                Network
            </h3>
        </div>

        <div class="space-y-3">
            <InfoItem
                :label="inf.toString()"
                v-for="(val, inf) in headerInfo"
                :value="val"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import type { HeaderInfo } from "@/types/ip";
import { getCountryFlag } from "@/utils/utils";
import { defineProps } from "vue";
import InfoItem from "./InfoItem.vue";

const props = defineProps<{
    headerInfo: HeaderInfo | null;
}>();
</script>
