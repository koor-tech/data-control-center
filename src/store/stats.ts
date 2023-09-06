import { defineStore } from 'pinia';

export interface StatsState {}

export const useStatsStore = defineStore('stats', {
    state: () => ({}) as StatsState,
    persist: false,
    actions: {
        // TODO
    },
});
