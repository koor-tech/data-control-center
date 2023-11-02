import { ConnectError } from '@connectrpc/connect';
import { useNotificationsStore } from '~/store/notifications';

export interface AuthState {
    accessToken: null | string;
    accessTokenExpiration: null | Date;
    accountID: bigint;
    loggingIn: boolean;
    loginError: null | string;
    permissions: string[];
    username: null | string;
}

export const useAuthStore = defineStore('auth', {
    state: () =>
        ({
            accessToken: null,
            accessTokenExpiration: null,
            loggingIn: false,
            loginError: null,
            permissions: [] as string[],
            username: null,
        }) as AuthState,
    persist: {
        paths: ['accessToken', 'accessTokenExpiration'],
    },
    actions: {
        loginStart(): void {
            this.loggingIn = true;
        },
        loginStop(errorMessage: null | string): void {
            this.loggingIn = false;
            this.loginError = errorMessage;
        },
        setAccessToken(accessToken: null | string, expiration: null | bigint | Date): void {
            this.accessToken = accessToken;
            if (typeof expiration === 'bigint') expiration = new Date(expiration.toString());
            this.accessTokenExpiration = expiration;
        },
        setPermissions(permissions: string[]): void {
            this.permissions = permissions.sort();
        },
        clearAuthInfo(): void {
            this.accessToken = null;
            this.accessTokenExpiration = null;
            this.accountID = 0n;
            this.permissions = [];
            this.username = null;
        },

        async doLogin(username: string, password: string): Promise<void> {
            // Start login
            this.loginStart();
            this.setPermissions([]);

            const { $grpc } = useNuxtApp();

            try {
                const call = $grpc.getUnAuthClient().login({
                    username,
                    password,
                });
                const response = await call;

                this.loginStop(null);
                this.setAccessToken(response.token, toDate(response.expires?.timestamp));

                await navigateTo({ name: 'index' });
            } catch (e) {
                this.loginStop((e as Error).message);
                this.setAccessToken(null, null);
                if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                throw e;
            }
        },
        async doLogout(): Promise<void> {
            const { $grpc } = useNuxtApp();
            try {
                await $grpc.getAuthClient().logout({});
                this.clearAuthInfo();
            } catch (e) {
                if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);

                this.clearAuthInfo();

                useNotificationsStore().dispatchNotification({
                    title: 'Error while logging you out',
                    content: 'Message: ' + (e as ConnectError).message,
                    type: 'error',
                });

                throw e;
            }
        },
    },

    getters: {
        getAccessTokenExpiration(state): null | Date {
            if (typeof state.accessTokenExpiration === 'string')
                state.accessTokenExpiration = new Date(Date.parse(state.accessTokenExpiration));

            return state.accessTokenExpiration;
        },
    },
});
