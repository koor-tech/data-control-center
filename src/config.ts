type AppConfig = {
    baseUrl: string;
};

const config: AppConfig = {
    baseUrl: '/api',
};

export default config;

export async function loadConfig(): Promise<void> {
    /*
    return new Promise(async (res, rej) => {
        try {
            const resp = await fetch('/api/config', {
                method: 'POST',
            });
            if (!resp.ok) {
                const text = await resp.text();
                throw createError({
                    statusCode: 500,
                    statusMessage: 'Failed to get Koor Data Control Center config from backend',
                    message: text,
                    fatal: true,
                    unhandled: false,
                });
            }
            const data = (await resp.json()) as AppConfig;
            config.baseUrl = data.baseUrl;

            return res();
        } catch (e) {
            return rej(e);
        }
    });
    */
}
