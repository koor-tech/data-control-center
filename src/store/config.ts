import { NuxtError } from 'nuxt/app';
import { StoreDefinition, defineStore } from 'pinia';

export interface ConfigState {
    fetched: boolean;
    appConfig: AppConfig;
}

type AppConfig = {
    baseUrl: string;
    login: LoginConfig;
};

type LoginConfig = {
    providers: ProviderConfig[];
};

type ProviderConfig = {
    name: string;
    label: string;
};

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
            return new Promise(async (res, rej) => {
                if (this.fetched) {
                    return res();
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

                    return res();
                } catch (e) {
                    showError(e as NuxtError);
                    return rej(e);
                }
            });
        },
    },
});

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useConfigStore as unknown as StoreDefinition, import.meta.hot));
}
