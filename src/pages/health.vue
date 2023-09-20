<script setup lang="ts">
import { statuses } from '~/composables/stats/types';
import { useStatsStore } from '~/store/stats';

useHead({
    title: 'Health Stats',
});
definePageMeta({
    title: 'Health Stats',
    requiresAuth: true,
});

const statsStore = useStatsStore();

const { data: clusterStats, error } = useLazyAsyncData('clusterStats', async () => await statsStore.getClusterStats());
</script>

<template>
    <div class="w-full h-full flex">
        <div class="w-full mx-4 my-4 rounded-lg border border-gray-300 py-4">
            <header class="inset-x-0 top-0 z-50 flex h-16 border-b border-gray-900/10">
                <div class="mx-auto flex w-full max-w-7xl items-center justify-between">
                    <div class="flex items-center gap-x-6">Cluster Health stats</div>
                </div>
            </header>
            <!-- Your content goes here -->
            <div class="w-full relative isolate">
                <!-- Secondary navigation -->
                <header class="w-full border-b border-b-gray-300">
                    <div class="mx-auto flex max-w-7xl flex-wrap items-center gap-6 sm:flex-nowrap">
                        <h1 v-if="clusterStats" class="text-base font-semibold leading-7 text-gray-900">
                            Cluster id: {{ clusterStats.id }}
                        </h1>
                        <div v-if="clusterStats" :class="[statuses[clusterStats.health], 'flex-none rounded-full p-1']">
                            <div class="h-2 w-2 rounded-full bg-current" />
                        </div>
                        <p v-if="clusterStats" class="text-gray-700 text-sm">
                            {{ clusterStats.health }}
                        </p>
                    </div>
                </header>

                <!-- Stats -->
                <dl class="mt-5 grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-5 px-2">
                    <template v-if="clusterStats && clusterStats.stats">
                        <template v-for="item in clusterStats.stats">
                            <HealthServices v-if="item.description.length" :statsContainer="item" />
                        </template>
                    </template>
                </dl>
            </div>
        </div>
    </div>
</template>
