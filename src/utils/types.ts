import type { DefineComponent } from 'vue';
import { type RoutesNamedLocations } from '~~/.nuxt/typed-router/__routes';

export type CardElement = {
    title: string;
    description?: string;
    href?: RoutesNamedLocations;
    permission?: string;
    icon?: DefineComponent;
    iconForeground?: string;
    iconBackground?: string;
};

export type CardElements = CardElement[];
