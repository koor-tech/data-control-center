<script setup lang="ts">
import { ConnectError } from '@connectrpc/connect';
import ClusterRadar from '~/components/ClusterRadar.vue';
import StatusDials from '~/components/StatusDials.vue';
import ClusterHealthServices from '~/components/health/ClusterHealthServices.vue';
//import NodeSummaryList from '~/components/NodeSummaryList.vue';

useHead({
    title: 'Overview',
});
definePageMeta({
    title: 'Overview',
    requiresAuth: true,
});

const { $grpc } = useNuxtApp();

const { data: stats, pending, refresh, error } = useLazyAsyncData(`clusterstats`, () => $grpc.getStatsClient().getClusterStats({}));
watch(error, () => {
    if (error.value !== null) $grpc.handleError(error.value as ConnectError);
});

const { data: radar } = useLazyAsyncData(`clusterradar`, () => $grpc.getStatsClient().getClusterRadar({}));

const { data: resources } = useLazyAsyncData(`clusterresources`, () => $grpc.getStatsClient().getClusterResources({}));
</script>

<template>
    <div class="flex flex-col">
        <div class="">
            <ClusterHealthServices v-if="stats && stats.stats" :stats="stats.stats" />
        </div>
        <div class="row">
            <div class="col-4">
                <ClusterRadar v-if="radar && radar.radar" :radar="radar.radar" />
            </div>
            <div class="col-8">
                <StatusDials v-if="radar && radar.radar" :radar="radar.radar" />
            </div>
        </div>
        <div class="row">
            <div class="col-4">
                <!-- <NodeSummaryList />-->
            </div>
            <div class="col-8" v-if="resources">
                <ResourceInfoList :rows="resources.deployments" />
                <ResourceInfoList :rows="resources.resources" />
            </div>
        </div>
    </div>
</template>
