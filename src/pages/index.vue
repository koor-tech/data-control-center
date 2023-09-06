<template>
    <div class="flex flex-col">
        <div class="">
            <ClusterHealthServices v-if="stats" :stats="stats" />
        </div>
        <div class="row">
            <div class="col-4">
                <ClusterRadar />
            </div>
            <div class="col-8">
                <StatusDials />
            </div>
        </div>
        <div class="row">
            <div class="col-4">
                <!-- <NodeSummaryList />-->
            </div>
            <div class="col-8">
                <OSDHealthWidget />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ConnectError } from '@connectrpc/connect';
import ClusterRadar from '~/components/ClusterRadar.vue';
import OSDHealthWidget from '~/components/OSDHealthWidget.vue';
import StatusDials from '~/components/StatusDials.vue';
import ClusterHealthServices from '~/components/health/ClusterHealthServices.vue';
//import NodeSummaryList from '~/components/NodeSummaryList.vue';

const { $grpc } = useNuxtApp();

const { data: stats, pending, refresh, error } = useLazyAsyncData(`clusterstats`, () => $grpc.getStatsClient().getClusterStats({}));
watch(error, () => {
    if (error.value !== null) $grpc.handleError(error.value as ConnectError);
});
</script>
