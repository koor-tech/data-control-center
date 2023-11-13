<script setup lang="ts">
import { ConnectError } from '@connectrpc/connect';
import ClusterHealthBar from '~/components/cluster/ClusterHealthBar.vue';
import ClusterServices from '~/components/cluster/ClusterServices.vue';
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

const { data: stats } = useLazyAsyncData('clusterStats', async () => {
    try {
        return await statsStore.getClusterStats();
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
});
</script>

<template>
    <div class="p-2">
        <div v-if="stats" class="flex flex-col gap-2">
            <ClusterHealthBar :health="stats.status" :cluster-id="stats.id" />

            <ClusterServices :stats="stats" />
        </div>
    </div>
</template>
