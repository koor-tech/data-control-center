import { type RoutesNamedLocations } from '@typed-router';
// eslint-disable-next-line import/order
import type { DefineComponent } from 'vue';

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
