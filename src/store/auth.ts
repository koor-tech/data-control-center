import { defineStore } from 'pinia';

export interface AuthState {
    accessToken: null | string;
    accessTokenExpiration: null | Date;
    accountID: bigint;
    loggingIn: boolean;
    loginError: null | string;
    permissions: string[];
}

export const useAuthStore = defineStore('auth', {
    state: () =>
        ({
            accessToken: null,
            accessTokenExpiration: null,
            loggingIn: false,
            loginError: null,
            permissions: [] as string[],
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
        signOut(): void {
            this.accessToken = null;
            this.accessTokenExpiration = null;
            this.accountID = 0n;
            this.permissions = [];
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
