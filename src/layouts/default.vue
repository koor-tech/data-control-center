<script setup lang="ts">
import LoadingBar from '~/components/partials/LoadingBar.vue';
import NotificationProvider from '~/components/partials/notifications/NotificationProvider.vue';
import Sidebar from '~/components/partials/Sidebar.vue';
</script>

<template>
    <NotificationProvider>
        <Sidebar>
            <div class="h-full">
                <LoadingBar />
                <slot />
            </div>
        </Sidebar>
    </NotificationProvider>
</template>
