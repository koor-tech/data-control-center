<script lang="ts" setup>
import { CloseCircleIcon } from 'mdi-vue3';
import { DefineComponent, markRaw } from 'vue';

withDefaults(
    defineProps<{
        title: string;
        message?: string | undefined;
        icon?: DefineComponent;
        type?: 'error' | 'warn';
    }>(),
    {
        message: undefined,
        icon: markRaw(CloseCircleIcon),
        type: 'error',
    },
);
</script>

<template>
    <div :class="`p-4 mt-6 rounded-md bg-error-100`">
        <div class="flex">
            <div class="flex-shrink-0">
                <component :is="icon" :class="`w-5 h-5 text-error-400`" aria-hidden="true" />
            </div>
            <div class="ml-3">
                <h3 :class="`text-sm font-medium text-error-600`">
                    {{ title }}
                </h3>
                <div v-if="message" :class="`mt-2 text-sm text-error-600`">
                    {{ message }}
                </div>
            </div>
        </div>
    </div>
</template>
