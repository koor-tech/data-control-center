import { NuxtError } from 'nuxt/app';
import { StoreDefinition, defineStore } from 'pinia';

type ProviderConfig = {
    name: string;
    label: string;
};

type LoginConfig = {
    providers: ProviderConfig[];
};

type AppConfig = {
    baseUrl: string;
    login: LoginConfig;
};

export interface ConfigState {
    fetched: boolean;
    appConfig: AppConfig;
}

export const useConfigStore = defineStore('config', {
    state: () =>
        ({
            fetched: false,
            appConfig: {
                baseUrl: '/api',
                login: {
                    providers: [],
                },
            } as AppConfig,
        }) as ConfigState,
    persist: false,
    actions: {
        async loadConfig(): Promise<void> {
            if (this.fetched) {
                return;
            }

            try {
                // 6 seconds should be enough
                const abort = new AbortController();
                const tId = setTimeout(() => abort.abort(), 8000);

                const resp = await fetch('/api/config', {
                    method: 'POST',
                    signal: abort.signal,
                });
                clearTimeout(tId);

                if (!resp.ok) {
                    const text = await resp.text();
                    throw createError({
                        statusCode: 500,
                        statusMessage: 'Failed to get Koor data-control-center config from backend',
                        message: text,
                        fatal: true,
                        unhandled: false,
                    });
                }
                const data = (await resp.json()) as AppConfig;
                data.baseUrl = '/api';
                this.appConfig = data;

                this.fetched = true;
            } catch (e) {
                showError(e as NuxtError);
                throw e;
            }
        },
    },
});

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useConfigStore as unknown as StoreDefinition, import.meta.hot));
}
