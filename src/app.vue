<script lang="ts" setup>
import { initFlowbite } from 'flowbite';
import { onMounted } from 'vue';
import { useConfigStore } from '~/store/config';
import { useNotificationsStore } from '~/store/notifications';

const configStore = useConfigStore();
const { loadConfig } = configStore;
const { appConfig } = storeToRefs(configStore);

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

    if (appConfig.value.updateAvailable !== undefined) {
        notifications.dispatchNotification({
            title: 'New version available!',
            content:
                'A new Data Control Center version is available on GitHub. Be sure to update to get new features and fixes.',
            type: 'info',
        });
    }
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
