<template>
    <div
        class="flex justify-between items-center py-2 border-b border-slate-100 dark:border-slate-800 last:border-b-0"
    >
        <dt
            class="text-xs font-medium text-slate-600 dark:text-slate-400 uppercase tracking-wide"
        >
            {{ label }}
        </dt>
        <dd
            class="text-sm font-medium text-slate-900 dark:text-slate-100 text-right max-w-[60%] truncate"
        >
            {{ value || "N/A" }}
        </dd>
    </div>
</template>

<script setup lang="ts">
interface Props {
    label: string;
    value: string | number;
}

defineProps<Props>();
</script>
