/** @type {import('tailwindcss').Config} */

const defaultTheme = require('tailwindcss/defaultTheme');

module.exports = {
    mode: 'jit',
    content: [
        './src/components/**/*.{js,vue,ts}',
        './src/layouts/**/*.vue',
        './src/pages/**/*.vue',
        './src/plugins/**/*.{js,ts}',
        './nuxt.config.{js,ts}',
        './node_modules/flowbite/**/*.{js,ts}',
    ],
    plugins: [require('@tailwindcss/forms'), require('@tailwindcss/typography'), require('flowbite/plugin')],
    theme: {
        extend: {
            fontFamily: {
                sans: ['Inter var', ...defaultTheme.fontFamily.sans],
            },
            boxShadow: {
                float: '0px 0px 22px 0px rgba(0,0,0,0.5)',
                glow: '0px 0px 22px 0px rgba(255,255,255,0.2)',
            },
            colors: {
                neutral: '#ffffff',
                base: {
                    100: '#e2e3e6',
                    200: '#c4c6cd',
                    300: '#a7aab4',
                    400: '#898d9b',
                    500: '#6c7182',
                    600: '#565a68',
                    700: '#41444e',
                    800: '#2b2d34',
                    850: '#202227',
                    900: '#16171a',
                },
                accent: {
                    50: '#f4f7f7',
                    100: '#e2ebeb',
                    200: '#c8d9d9',
                    300: '#a1bebf',
                    400: '#739b9d',
                    500: '#588082',
                    600: '#4c6b6e',
                    700: '#42595c',
                    800: '#3b4c4f',
                    900: '#354244',
                    950: '#1d2628',
                },
                primary: {
                    50: '#eafff7',
                    100: '#ccffe9',
                    200: '#9efcd8',
                    300: '#60f5c6',
                    400: '#19dda5',
                    500: '#00cd98',
                    600: '#00a77d',
                    700: '#008668',
                    800: '#006a53',
                    900: '#005746',
                    950: '#003129',
                },
                secondary: {
                    50: '#f4f7f7',
                    100: '#e2ebeb',
                    200: '#c8d9d9',
                    300: '#a1bebf',
                    400: '#739b9d',
                    500: '#588082',
                    600: '#4c6b6e',
                    700: '#42595c',
                    800: '#3b4c4f',
                    900: '#354244',
                    950: '#1d2628',
                },
                error: {
                    100: '#f7d4d7',
                    200: '#efa8af',
                    300: '#e77d88',
                    400: '#df5160',
                    500: '#d72638',
                    600: '#ac1e2d',
                    700: '#811722',
                    800: '#560f16',
                    900: '#2b080b',
                },
                warn: {
                    100: '#fdefd5',
                    200: '#fbdfab',
                    300: '#face81',
                    400: '#f8be57',
                    500: '#f6ae2d',
                    600: '#c58b24',
                    700: '#94681b',
                    800: '#624612',
                    900: '#312309',
                },
                info: {
                    100: '#ddf8f9',
                    200: '#bbf1f3',
                    300: '#99ebec',
                    400: '#77e4e6',
                    500: '#55dde0',
                    600: '#44b1b3',
                    700: '#338586',
                    800: '#22585a',
                    900: '#112c2d',
                },
                success: {
                    100: '#d3f5e1',
                    200: '#a7ebc4',
                    300: '#7be2a6',
                    400: '#4fd889',
                    500: '#23ce6b',
                    600: '#1ca556',
                    700: '#157c40',
                    800: '#0e522b',
                    900: '#072915',
                },
            },
            maxWidth: {
                '8xl': '88rem',
            },
            animation: {
                wiggle: 'wiggle 1s infinite',
            },
            keyframes: {
                wiggle: {
                    '0%': {
                        transform: 'rotate(0deg)',
                    },
                    '80%': {
                        transform: 'rotate(0deg)',
                    },
                    '85%': {
                        transform: 'rotate(5deg)',
                    },
                    '95%': {
                        transform: 'rotate(-5deg)',
                    },
                    '100%': {
                        transform: 'rotate(0deg)',
                    },
                },
            },
            zIndex: {
                100: 100,
                90: 90,
            },
        },
    },
};
