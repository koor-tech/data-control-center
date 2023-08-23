<script lang="ts" setup>
import { NuxtError } from 'nuxt/app';
import { loadConfig } from '~/config';

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

try {
    await loadConfig();
} catch (e) {
    showError(e as NuxtError);
}
</script>

<template>
    <NuxtLayout>
        <NuxtPage :transition="{
            name: 'page',
            mode: 'out-in',
        }" />
    </NuxtLayout>
</template>
