<script setup lang="ts">
import { ConnectError } from '@connectrpc/connect';
import ClusterHealthBar from '~/components/cluster/ClusterHealthBar.vue';
import ClusterHealthServices from '~/components/cluster/ClusterHealthServices.vue';
import { TransformStats } from '~/composables/stats/transform';
import { useStatsStore } from '~/store/stats';

useHead({
    title: 'Health Stats',
});
definePageMeta({
    title: 'Health Stats',
    requiresAuth: true,
});

const { $grpc } = useNuxtApp();

const statsStore = useStatsStore();

const { data: clusterStats } = useLazyAsyncData('clusterStats', async () => {
    try {
        const stats = await statsStore.getClusterStats();
        const dataStats = new TransformStats(stats);
        return dataStats.display();
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
});
</script>

<template>
    <div class="p-2">
        <div class="flex flex-col gap-2">
            <ClusterHealthBar v-if="clusterStats" :cluster-stats="clusterStats" />

            <ClusterHealthServices v-if="clusterStats?.stats" :stats="clusterStats?.stats" />
        </div>
    </div>
</template>
