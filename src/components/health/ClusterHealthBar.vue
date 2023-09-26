<script lang="ts" setup>
import { ClusterHealthStats } from '~/composables/stats/types';
import { ClusterHealth } from '~~/gen/ts/api/resources/stats/v1/stats_pb';
import Container from '../partials/Container.vue';
import { clusterHealthToBGColor, clusterHealthToTextColor } from './helpers';

defineProps<{
    clusterStats: ClusterHealthStats;
}>();
</script>

<template>
    <Container>
        <header class="w-full">
            <div class="mx-auto grid grid-cols-2 items-center justify-center">
                <div class="inline-flex items-center justify-start gap-2">
                    <div
                        class="h-4 w-4 rounded-full p-1 animate-pulse"
                        :class="[clusterHealthToBGColor(clusterStats.health), clusterHealthToTextColor(clusterStats.health)]"
                    >
                        <div class="h-2 w-2 rounded-full bg-current" />
                    </div>
                    <p class="font-semibold text-gray-700 text-sm">
                        {{ ClusterHealth[clusterStats.health] }}
                    </p>
                </div>
                <p class="text-sm leading-7 text-gray-900 justify-self-end">
                    Cluster ID: <code>{{ clusterStats.id }}</code>
                </p>
            </div>
        </header>
    </Container>
</template>
