<script lang="ts" setup>
const props = withDefaults(
    defineProps<{
        color: 'gray' | 'error' | 'warn' | 'info' | 'success' | 'primary' | 'secondary';
    }>(),
    {
        color: 'gray',
    },
);

const colorClasses = computed(() => {
    switch (props.color) {
        case 'error':
            return 'bg-error-50 ring-error-500/10';
        case 'warn':
            return 'bg-warn-50 ring-warn-500/10';
        case 'info':
            return 'bg-info-50 ring-info-500/10';
        case 'success':
            return 'bg-success-50 ring-success-500/10';
        case 'primary':
            return 'bg-primary-50 ring-primary-500/10';
        case 'secondary':
            return 'bg-secondary-50 ring-secondary-500/10';
        case 'gray':
        default:
            return 'bg-gray-50 ring-gray-500/10';
    }
});
</script>
<template>
    <span
        class="inline-flex items-center rounded-md px-2 py-1 text-xs font-medium ring-1 ring-inset text-gray-600"
        :class="colorClasses"
    >
        <slot />
    </span>
</template>
