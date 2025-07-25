<template>
    <div
        class="flex items-center justify-between py-2 border-b border-slate-100 dark:border-slate-800 last:border-b-0"
    >
        <span
            class="text-xs font-medium text-slate-600 dark:text-slate-400 uppercase tracking-wide"
            >{{ label }}</span
        >
        <div class="flex items-center">
            <div
                class="w-2 h-2 rounded-full mr-2"
                :class="
                    active ? 'bg-green-500' : 'bg-slate-300 dark:bg-slate-600'
                "
            ></div>
            <span
                class="text-xs font-medium"
                :class="
                    active
                        ? 'text-green-700 dark:text-green-400'
                        : 'text-slate-500 dark:text-slate-400'
                "
            >
                {{ active ? "Yes" : "No" }}
            </span>
        </div>
    </div>
</template>

<script setup lang="ts">
interface Props {
    label: string;
    active: boolean;
}

defineProps<Props>();
</script>
