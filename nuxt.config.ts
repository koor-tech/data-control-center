import fs from 'fs';
import { defineNuxtConfig } from 'nuxt/config';

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
        'nuxt-typed-router',
        '@nuxtjs/eslint-module',
        '@nuxtjs/robots',
        '@nuxtjs/tailwindcss',
        '@pinia/nuxt',
        '@pinia-plugin-persistedstate/nuxt',
        '@vee-validate/nuxt',
    ],
    devtools: {
        enabled: true,
        vscode: {},
        timeline: {
            enabled: true,
        },
    },
    eslint: {
        lintOnStart: false,
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
        debug: false,
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
                protocol: 'ws',
            },
            proxy: {
                '/api': {
                    target: 'http://localhost:8282',
                },
            },
        },
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
            'postcss-import': {},
            'tailwindcss/nesting': 'postcss-nesting',
            tailwindcss: {},
            autoprefixer: {},
            ...(process.env.NODE_ENV === 'production' ? { cssnano: {} } : {}),
        },
    },
    typescript: {
        typeCheck: false,
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
            link: [{ rel: 'icon', type: 'image/png', href: '/icons/favicon-128x128.png' }],
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
        emitRouteChunkError: 'automatic',
    },
});

export default config;
