declare module '#app' {
    interface PageMeta {
        title?: string;
        requiresAuth?: boolean;
        permission?: string;
        authOnlyToken?: boolean;
    }
}

declare module 'vue-router' {
    interface PageMeta {
        title?: string;
        requiresAuth?: boolean;
        permission?: string;
        authOnlyToken?: boolean;
    }
}

type ProviderConfig = {
    name: string;
    label: string;
};

type LoginConfig = {
    providers: ProviderConfig[];
};

export type AppConfig = {
    version: string;
    login: LoginConfig;
    readOnly: boolean;
    updateAvailable: null | string;
};

// It is always important to ensure you import/export something when augmenting a type
export {};
