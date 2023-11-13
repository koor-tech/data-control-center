<script setup lang="ts">
const slots = useSlots();
</script>

<template>
    <li class="w-full flex justify-between gap-x-2 py-1">
        <div v-if="!!slots['default']" class="flex min-w-0 gap-x-2">
            <div class="min-w-0 flex-auto">
                <p class="mt-1 truncate text-xs leading-5 text-gray-500">
                    <slot />
                </p>
            </div>
        </div>
        <div class="hidden sm:flex sm:flex-col sm:items-end">
            <p class="text-sm leading-6 text-gray-900">
                <slot name="content" />
            </p>
        </div>
    </li>
</template>
