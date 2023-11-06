declare module '#app' {
    interface PageMeta {
        title?: string;
        requiresAuth?: boolean;
        permission?: string;
        authOnlyToken?: boolean;
    }
}

// It is always important to ensure you import/export something when augmenting a type
export {};
