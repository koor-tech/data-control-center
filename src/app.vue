<script lang="ts" setup>
import { useConfigStore } from '~/store/config';

const configStore = useConfigStore();
const { loadConfig } = configStore;

useHead({
    htmlAttrs: {
        class: 'h-full bg-base-900',
        lang: 'en',
    },
    bodyAttrs: {
        class: 'h-full overflow-hidden',
    },
    titleTemplate: (title?: string) => {
        return title ? `${title} - Koor Data Control Center` : 'Koor Data Control Center';
    },
});

await loadConfig();
</script>

<template>
    <NuxtLayout>
        <NuxtPage :transition="{
            name: 'page',
            mode: 'out-in',
        }" />
    </NuxtLayout>
</template>
