<template>
    <div
        class="bg-white/5 backdrop-blur-sm rounded-xl p-4 border border-white/10 hover:bg-white/10 transition-all duration-300 hover:scale-105"
    >
        <dt class="text-sm font-medium text-gray-400 mb-2">{{ label }}</dt>
        <dd class="text-lg font-semibold text-white break-words">
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
