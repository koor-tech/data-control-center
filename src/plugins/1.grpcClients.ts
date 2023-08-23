import { createPromiseClient, PromiseClient } from '@connectrpc/connect'
import { createConnectTransport,  } from '@connectrpc/connect-web'
// import type { PromiseClient } from '@connectrpc/connect'
// import { createPromiseClient, createConnectTransport } from '@improbable-eng/grpc-web';

import { AuthService } from '~~/gen/ts/api/services/auth/auth_connect'
import config from '~/config';
import {StatsService} from "~~/gen/ts/api/services/stats/stats_connect";
//import { useAuthStore } from '~/store/auth';

export default defineNuxtPlugin(() => {
    return {
        provide: {
            grpc: new GRPCClients(),
        },
    };
});

export class GRPCClients {
    //private authInterceptor: AuthInterceptor;

    constructor() {
        //this.authInterceptor = new AuthInterceptor();

        //const interceptors: RpcInterceptor[] = [this.authInterceptor];
    }

    // GRPC Clients ===============================================================
    // Account / Auth - Unauthorized and authorized clients
    getUnAuthClient(): PromiseClient<typeof AuthService> {
        return createPromiseClient(
            AuthService,
            createConnectTransport({
                baseUrl: config.baseUrl,
            })
        );
    }

    getStatsClient(): PromiseClient<typeof StatsService> {
        return createPromiseClient(
            StatsService,
            createConnectTransport({
                baseUrl: config.baseUrl,
            })
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
