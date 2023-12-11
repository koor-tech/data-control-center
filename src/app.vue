<script lang="ts" setup>
import { initFlowbite } from 'flowbite';
import { onMounted } from 'vue';
import { useNotificationsStore } from '~/store/notifications';

const appConfig = useAppConfig();

const notifications = useNotificationsStore();

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

    if (appConfig.updateAvailable !== null && appConfig.updateAvailable !== undefined && appConfig.updateAvailable !== '') {
        notifications.dispatchNotification({
            title: 'New version available!',
            content: `A new Data Control Center version (${appConfig.updateAvailable}) is available on GitHub. Be sure to update to get new features and fixes.`,
            type: 'info',
        });
    }
});
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
