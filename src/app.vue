<script lang="ts" setup>
import { initFlowbite } from 'flowbite';
import { onMounted } from 'vue';
import { useConfigStore } from '~/store/config';

const configStore = useConfigStore();
const { loadConfig } = configStore;

useHead({
    htmlAttrs: {
        class: 'h-full bg-white',
        lang: 'en',
    },
    bodyAttrs: {
        class: 'h-full overflow-hidden',
    },
    titleTemplate: (title?: string) => {
        return title ? `${title} - Koor Data Control Center` : 'Koor Data Control Center';
    },
});

// Initialize components based on data attribute selectors
onMounted(() => {
    initFlowbite();
});

await loadConfig();
</script>

<template>
    <NuxtLayout>
        <NuxtPage
            :transition="{
                name: 'page',
                mode: 'out-in',
            }"
        />
    </NuxtLayout>
</template>
