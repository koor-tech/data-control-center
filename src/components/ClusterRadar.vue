<template>
    <div class="bg-gray-100 min-h-screen p-8" v-if="clusterStats">
        <h1 class="text-1xl font-semibold mb-4">Cluster Health Dashboard</h1>
        <!-- Services -->
        <div class="bg-white rounded p-4">
            <h2 class="text-lg font-semibold">Cluster id:{{ clusterStats.id }}</h2>
        </div>
        <ClusterHealthServices :data="clusterStats" />
    </div>
</template>

<script setup lang="ts">
import ClusterHealthServices from '~/components/health/ClusterHealthServices.vue';

const { $grpc } = useNuxtApp();
const { data: clusterStats } = useLazyAsyncData('clusterStats', async () => $grpc.getStatsClient().getClusterStats({}));
</script>
