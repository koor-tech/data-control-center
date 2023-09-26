<script setup lang="ts">
import NotificationProvider from '~/components/partials/notifications/NotificationProvider.vue';
import Sidebar from '~/components/partials/Sidebar.vue';
</script>

<template>
    <NotificationProvider>
        <Sidebar>
            <div class="h-full">
                <NuxtLoadingIndicator />
                <slot />
            </div>
        </Sidebar>
    </NotificationProvider>
</template>
