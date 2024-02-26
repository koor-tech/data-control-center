<script lang="ts" setup>
import NotificationItem from '~/components/partials/notifications/NotificationItem.vue';
import { useNotificationsStore } from '~/store/notifications';

const notificationsStore = useNotificationsStore();
const { getNotifications } = storeToRefs(notificationsStore);
</script>

<template>
    <div>
        <slot />

        <!-- Global notification live region, render this permanently at the end of the document -->
        <div aria-live="assertive" class="pointer-events-none fixed inset-0 z-50 flex items-end px-4 py-16">
            <div class="flex w-full flex-col items-center space-y-4">
                <NotificationItem
                    v-for="(notification, idx) in getNotifications"
                    :key="notification.id?.toString()"
                    :notification="notification"
                    :class="idx > 0 ? 'mt-4' : ''"
                />
            </div>
        </div>
    </div>
</template>
