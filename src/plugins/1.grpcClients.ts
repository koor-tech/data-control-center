import { createPromiseClient, Interceptor, PromiseClient } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { useAuthStore } from '~/store/auth';
import { useConfigStore } from '~/store/config';
import { useNotificationsStore } from '~/store/notifications';
import { AuthService } from '~~/gen/ts/api/services/auth/auth_connect';
import { StatsService } from '~~/gen/ts/api/services/stats/stats_connect';

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
    //private authInterceptor: AuthInterceptor;

    constructor() {
        //this.authInterceptor = new AuthInterceptor();
        //const interceptors: RpcInterceptor[] = [this.authInterceptor];
    }

    // Handle GRPC errors
    async handleError(err: Error): Promise<boolean> {
        const notifications = useNotificationsStore();
        notifications.dispatchNotification({
            type: 'error',
            title: 'Error during API request',
            content: err.message,
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
}

/*
export class AuthInterceptor implements RpcInterceptor {
    interceptUnary(next: NextUnaryFn, method: MethodInfo, input: object, options: RpcOptions): UnaryCall {
        if (!options.meta) {
            options.meta = {};
        }

        const { accessToken } = useAuthStore();
        if (accessToken !== null) {
            options.meta['Authorization'] = 'Bearer ' + accessToken;
        }

        return next(method, input, options);
    }

    interceptServerStreaming?(
        next: NextServerStreamingFn,
        method: MethodInfo,
        input: object,
        options: RpcOptions,
    ): ServerStreamingCall {
        if (!options.meta) {
            options.meta = {};
        }

        const { accessToken } = useAuthStore();
        if (accessToken !== null) {
            options.meta['Authorization'] = 'Bearer ' + accessToken;
        }

        return next(method, input, options);
    }
}
*/
