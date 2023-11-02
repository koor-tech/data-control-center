import { Code, ConnectError, createPromiseClient, Interceptor, PromiseClient } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { Notification } from '~/composables/notification/interfaces/Notification.interface';
import { useAuthStore } from '~/store/auth';
import { useConfigStore } from '~/store/config';
import { useNotificationsStore } from '~/store/notifications';
import { AuthService } from '~~/gen/ts/api/services/auth/v1/auth_connect';
import { ClusterService } from '~~/gen/ts/api/services/cluster/v1/cluster_connect';
import { StatsService } from '~~/gen/ts/api/services/stats/v1/stats_connect';

export default defineNuxtPlugin(() => {
    return {
        provide: {
            grpc: new GRPCClients(),
        },
    };
});

const authInterceptor: Interceptor = (next) => async (req) => {
    const authStore = useAuthStore();
    if (authStore.accessToken) req.header.set('Authorization', 'Bearer ' + authStore.accessToken);
    return await next(req);
};

export class GRPCClients {
    // Handle GRPC errors
    async handleError(err: ConnectError): Promise<boolean> {
        const notifications = useNotificationsStore();

        const notification = {
            id: '',
            type: 'error',
            title: 'Internal Error',
            content: err.message ?? 'Unknown error',
        } as Notification;

        if (err.code !== undefined) {
            const route = useRoute();
            const redirect = route.query.redirect ?? route.fullPath;

            switch (err.code) {
                case Code.Internal:
                    break;

                case Code.Unavailable:
                    notification.title = '';
                    notification.content = 'Unavailable content';
                    break;

                case Code.Unauthenticated:
                    useAuthStore().clearAuthInfo();

                    notification.type = 'warning';
                    notification.title = 'Unauthenticated';
                    notification.content = 'Your request was . Please login again.';

                    // Only update the redirect query param if it isn't already set
                    navigateTo({
                        name: 'auth-login',
                        query: { redirect },
                        replace: true,
                        force: true,
                    });
                    break;

                case Code.PermissionDenied:
                    notification.title = 'Permission denied';
                    break;

                case Code.NotFound:
                    notification.title = 'Not Found';
                    break;

                default:
                    notification.title = 'Error occured';
                    notification.content = 'Message: ' + err.message + '(Code: ' + err.code.valueOf() + ')';
                    break;
            }
        }

        notifications.dispatchNotification({
            type: notification.type,
            title: notification.title,
            content: notification.content,
        });

        return true;
    }

    // GRPC Clients ===============================================================
    // Account / Auth - Unauthorized and authorized clients
    getUnAuthClient(): PromiseClient<typeof AuthService> {
        return createPromiseClient(
            AuthService,
            createConnectTransport({
                baseUrl: useConfigStore().appConfig.baseUrl,
            }),
        );
    }

    getAuthClient(): PromiseClient<typeof AuthService> {
        return createPromiseClient(
            AuthService,
            createConnectTransport({
                baseUrl: useConfigStore().appConfig.baseUrl,
                interceptors: [authInterceptor],
            }),
        );
    }

    getStatsClient(): PromiseClient<typeof StatsService> {
        return createPromiseClient(
            StatsService,
            createConnectTransport({
                baseUrl: useConfigStore().appConfig.baseUrl,
                interceptors: [authInterceptor],
            }),
        );
    }

    getClusterClient(): PromiseClient<typeof ClusterService> {
        return createPromiseClient(
            ClusterService,
            createConnectTransport({
                baseUrl: useConfigStore().appConfig.baseUrl,
                interceptors: [authInterceptor],
            }),
        );
    }
}
