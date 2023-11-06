import { defineStore, type StoreDefinition } from 'pinia';
import { v4 as uuidv4 } from 'uuid';
import { type Notification } from '~/composables/notification/interfaces/Notification.interface';
import { type NotificationConfig } from '~/composables/notification/interfaces/NotificationConfig.interface';

export interface NotificationsState {
    lastId: string;
    notifications: Notification[];
}

export const useNotificationsStore = defineStore('notifications', {
    state: () =>
        ({
            lastId: '0',
            notifications: [],
        }) as NotificationsState,
    persist: {
        paths: ['lastId'],
    },
    actions: {
        setLastId(lastId: bigint): void {
            this.lastId = lastId.toString();
        },
        removeNotification(id: string): void {
            this.notifications = this.notifications.filter((notification) => notification.id !== id);
        },
        dispatchNotification({ title, content, type, autoClose = true, duration = 6000 }: NotificationConfig) {
            const id = uuidv4();
            this.notifications.unshift({
                id,
                title,
                content,
                type,
            });

            if (autoClose) {
                setTimeout(() => {
                    this.removeNotification(id);
                }, duration);
            }
        },
    },
    getters: {
        getLastId(): bigint {
            return BigInt(this.lastId);
        },
        getNotifications(): Notification[] {
            return [...this.notifications].slice(0, 4);
        },
        getNotificationsCount(): number {
            return this.notifications.length;
        },
    },
});

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useNotificationsStore as unknown as StoreDefinition, import.meta.hot));
}
