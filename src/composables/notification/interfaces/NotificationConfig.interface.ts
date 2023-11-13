import type { NotificationType } from '~/composables/notification/interfaces/Notification.interface';

export interface NotificationConfig {
    title: string;
    content: string;
    duration?: number;
    autoClose?: boolean;
    type?: NotificationType;
}
