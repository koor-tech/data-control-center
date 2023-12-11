<script lang="ts" setup>
import type { ConnectError } from '@connectrpc/connect';
import { useThrottleFn } from '@vueuse/core';
import DataNoDataBlock from '~/components/partials/data/DataNoDataBlock.vue';
import DataPendingBlock from '~/components/partials/data/DataPendingBlock.vue';
import DataErrorBlock from '~/components/partials/data/DataErrorBlock.vue';
import {
    GetNetworkTestStatusResponse,
    NetworkTestOutputFormat,
    StartNetworkTestResponse,
} from '~~/gen/ts/api/services/cluster/v1/cluster_pb';

const { $grpc } = useNuxtApp();

const {
    data: testStatus,
    error,
    pending,
    refresh,
} = useLazyAsyncData('cluster-network_test_status', async () => getNetworkTestStatus());

async function getNetworkTestStatus(): Promise<GetNetworkTestStatusResponse | undefined> {
    try {
        return await $grpc.getClusterClient().getNetworkTestStatus({});
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
}

async function startNetworkTest(): Promise<StartNetworkTestResponse | undefined> {
    try {
        return await $grpc.getClusterClient().startNetworkTest({
            hostNetwork: false,
            outputFormat: NetworkTestOutputFormat.CSV,
        });
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
}

const canSubmit = ref(true);
const onSubmitThrottle = useThrottleFn(async () => {
    canSubmit.value = false;
    await startNetworkTest().finally(() => setTimeout(() => (canSubmit.value = true), 400));
}, 1000);
</script>

<template>
    <div class="p-2">
        <h2 class="text-xl q-mt-lg text-center">Ancientt Network Test</h2>

        <DataErrorBlock v-if="error" :retry="refresh" :message="error.message" />
        <DataPendingBlock v-else-if="pending" message="Loading Troubleshoot report" />
        <DataNoDataBlock v-else-if="testStatus === null || testStatus === undefined" />

        <template v-else>
            <p class="font-semibold text-gray-700 text-sm">
                <template v-if="testStatus">
                    {{ testStatus.running }}
                </template>
                <template v-else> N/A </template>
            </p>

            <button
                type="submit"
                class="flex justify-center w-full px-3 py-2 text-sm font-semibold transition-colors rounded-md text-neutral focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2"
                :disabled="!canSubmit"
                :class="[
                    !canSubmit
                        ? 'disabled bg-base-500 hover:bg-base-400 focus-visible:outline-base-500'
                        : 'bg-primary-500 hover:bg-primary-400 focus-visible:outline-primary-500',
                ]"
                @click="onSubmitThrottle"
            >
                Start Network Test
            </button>
        </template>
    </div>
</template>
