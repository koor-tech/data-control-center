<script lang="ts" setup>
import { HelpIcon } from 'mdi-vue3';
import type { DefineComponent } from 'vue';

withDefaults(
    defineProps<{
        icon?: DefineComponent;
        color?: string;
    }>(),
    {
        icon: HelpIcon,
        color: 'bg-accent-500',
    },
);
</script>

<template>
    <div class="relative overflow-hidden rounded-lg bg-white px-4 py-4 shadow">
        <dt>
            <div class="flex justify-center rounded-md p-3 text-white" :class="color">
                <component :is="icon" class="h-6 w-6" aria-hidden="true" />
                <p class="ml-4 col-span-2 text-sm font-medium">
                    <slot />
                </p>
            </div>
        </dt>
        <dd class="flex">
            <ul role="list" class="w-full divide-y divide-gray-100">
                <slot name="content" />
            </ul>
        </dd>
    </div>
</template>
