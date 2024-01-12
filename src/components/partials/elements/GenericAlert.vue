<script lang="ts" setup>
import { CloseCircleIcon } from 'mdi-vue3';
import { type DefineComponent, markRaw } from 'vue';

withDefaults(
    defineProps<{
        title: string;
        message?: string | undefined;
        icon?: DefineComponent;
        boxClass?: string;
        iconClass?: string;
        textClass?: string;
    }>(),
    {
        message: undefined,
        icon: markRaw(CloseCircleIcon),
        boxClass: 'bg-error-100',
        iconClass: 'text-error-400',
        textClass: 'text-error-600',
    },
);
</script>

<template>
    <div class="p-4 mt-6 rounded-md" :class="boxClass">
        <div class="flex">
            <div class="flex-shrink-0">
                <component :is="icon" class="w-5 h-5" :class="iconClass" aria-hidden="true" />
            </div>
            <div class="ml-3">
                <h3 class="text-sm font-medium text-error-600" :class="textClass">
                    {{ title }}
                </h3>
                <div v-if="message" class="mt-2 text-sm" :class="textClass">
                    {{ message }}
                </div>
            </div>
        </div>
    </div>
</template>
