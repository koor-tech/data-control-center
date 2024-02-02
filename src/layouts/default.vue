<script setup lang="ts">
import LoadingBar from '~/components/partials/LoadingBar.vue';
import NotificationProvider from '~/components/partials/notifications/NotificationProvider.vue';
import SidebarContainer from '~/components/partials/SidebarContainer.vue';
</script>

<template>
    <NotificationProvider>
        <SidebarContainer>
            <div class="h-full">
                <LoadingBar />
                <slot />
            </div>
        </SidebarContainer>
    </NotificationProvider>
</template>
