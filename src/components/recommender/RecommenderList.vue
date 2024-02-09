<script lang="ts" setup>
import type { ConnectError } from '@connectrpc/connect';
import { RefreshIcon } from 'mdi-vue3';
import type { ListClusterRecommendationsResponse } from '~~/gen/ts/api/services/stats/v1/stats_pb';
import DataNoDataBlock from '~/components/partials/data/DataNoDataBlock.vue';
import DataPendingBlock from '~/components/partials/data/DataPendingBlock.vue';
import DataErrorBlock from '~/components/partials/data/DataErrorBlock.vue';

const { $grpc } = useNuxtApp();

const canSubmit = ref(true);

const { data, error, pending, refresh } = useLazyAsyncData('cluster-network_test_status', async () =>
    listClusterRecommendations().finally(() => setTimeout(() => (canSubmit.value = true))),
);

async function listClusterRecommendations(): Promise<ListClusterRecommendationsResponse | undefined> {
    try {
        return await $grpc.getStatsClient().listClusterRecommendations({});
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
}
</script>

<template>
    <div class="p-2">
        <h2 class="text-xl q-mt-lg text-center">Cluster Recommendations</h2>

        <div class="flex items-center gap-1">
            <button
                type="button"
                class="flex-1 flex justify-center w-full px-3 py-2 text-sm font-semibold transition-colors rounded-md text-neutral focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2"
                :disabled="!canSubmit"
                :class="[
                    !canSubmit
                        ? 'disabled bg-base-500 hover:bg-base-400 focus-visible:outline-base-500'
                        : 'bg-primary-500 hover:bg-primary-400 focus-visible:outline-primary-500',
                ]"
                @click="refresh()"
            >
                <RefreshIcon class="h-5 w-5" />
            </button>
        </div>

        <div>
            <DataErrorBlock v-if="error" :retry="refresh" :message="error.message" />
            <DataPendingBlock v-else-if="pending" message="Loading Network Test status" />
            <DataNoDataBlock v-else-if="data === null || data === undefined" />

            <template v-else>
                <div class="mt-2 gap-2 grid grid-col-2">
                    <RecommenderListEntry
                        v-for="recommendation in data.recommendations.sort((a, b) => a.type - b.type)"
                        :key="recommendation.title"
                        :recommendation="recommendation"
                    />
                </div>
            </template>
        </div>
    </div>
</template>
