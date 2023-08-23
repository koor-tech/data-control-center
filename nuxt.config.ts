import fs from 'fs';
import { defineNuxtConfig } from 'nuxt/config';
import path from 'path';
import mkcert from 'vite-plugin-mkcert';

type PackageJson = {
    version: string;
};

const packageJson = fs.readFileSync('./package.json');
const version: string = (JSON.parse(packageJson.toString()) as PackageJson).version || '0.0.0';

// https://nuxt.com/docs/api/configuration/nuxt-config
const config = defineNuxtConfig({
    srcDir: 'src/',
    telemetry: false,
    ssr: false,
    modules: [
        '@nuxt/devtools',
        '@pinia/nuxt',
        '@pinia-plugin-persistedstate/nuxt',
        'nuxt-typed-router',
        '@nuxtjs/robots',
        '@nuxtjs/tailwindcss',
        '@vee-validate/nuxt',
    ],
    devtools: {
        enabled: true,
        vscode: {},
        timeline: {
            enabled: true,
        },
    },
    pinia: {
        autoImports: [
            // automatically imports `defineStore`
            'defineStore', // import { defineStore } from 'pinia'
            ['defineStore', 'definePiniaStore'], // import { defineStore as definePiniaStore } from 'pinia'
            'acceptHMRUpdate', // import { acceptHMRUpdate } from 'pinia'
            'storeToRefs',
        ],
    },
    robots: {
        rules: {
            UserAgent: '*',
            Disallow: '/',
            Allow: ['/$', '/index.html'],
        },
    },
    piniaPersistedstate: {
        storage: 'localStorage',
        debug: true,
    },
    veeValidate: {
        // disable or enable auto imports
        autoImports: true,
        // Use different names for components
        componentNames: {
            Form: 'VeeForm',
            Field: 'VeeField',
            FieldArray: 'VeeFieldArray',
            ErrorMessage: 'VeeErrorMessage',
        },
    },
    vite: {
        define: {
            __APP_VERSION__: '"' + version + '"',
        },
        build: {
            commonjsOptions: {
                transformMixedEsModules: true,
            },
            manifest: true,
        },
        server: {
            hmr: {
                protocol: 'wss',
            },
            https: true,
            proxy: {
                '/api': {
                    target: 'http://localhost:8282',
                },
            },
        },
        plugins: [mkcert()],
    },
    css: [
        // Inter font (all weights)
        '@fontsource/inter/100.css',
        '@fontsource/inter/200.css',
        '@fontsource/inter/300.css',
        '@fontsource/inter/400.css',
        '@fontsource/inter/500.css',
        '@fontsource/inter/600.css',
        '@fontsource/inter/700.css',
        '@fontsource/inter/800.css',
        '@fontsource/inter/900.css',
    ],
    postcss: {
        plugins: {
            tailwindcss: {
                configPath: '~~/tailwind.config',
            },
            autoprefixer: {},
        },
    },
    typescript: {
        strict: true,
        tsConfig: {
            compilerOptions: {
                removeComments: true,
            },
        },
    },
    devServer: {},
    app: {
        head: {
            charset: 'utf-8',
            viewport: 'width=device-width, initial-scale=1',
            link: [{ rel: 'icon', type: 'image/png', href: '/images/logo.png' }],
            htmlAttrs: {
                class: 'h-full',
            },
            bodyAttrs: {
                class: 'h-full',
            },
        },
        pageTransition: { name: 'page', mode: 'out-in' },
    },
    experimental: {
        polyfillVueUseHead: false,
    },
});

if (process.env.NODE_ENV !== 'production') {
    config.devServer!.https = {
        // Use vite-mkcert-plugin's cert + key for localhost
        key: path.resolve(process.env.HOME!, '.vite-plugin-mkcert', 'dev.pem'),
        cert: path.resolve(process.env.HOME!, '.vite-plugin-mkcert', 'cert.pem'),
    };
}

export default config;
