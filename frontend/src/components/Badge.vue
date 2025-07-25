<template>
    <span
        class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium"
        :class="
            active ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'
        "
    >
        <svg
            v-if="active"
            class="w-4 h-4 mr-1"
            fill="currentColor"
            viewBox="0 0 20 20"
        >
            <path
                fill-rule="evenodd"
                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                clip-rule="evenodd"
            />
        </svg>
        <svg
            v-else
            class="w-4 h-4 mr-1"
            fill="currentColor"
            viewBox="0 0 20 20"
        >
            <path
                fill-rule="evenodd"
                d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                clip-rule="evenodd"
            />
        </svg>
        {{ label }}
    </span>
</template>

<script setup lang="ts">
interface Props {
    label: string;
    active: boolean;
}

defineProps<Props>();
</script>
