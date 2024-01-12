<script lang="ts" setup>
import { saveAs } from 'file-saver';
import { RefreshIcon } from 'mdi-vue3';
import type { ConnectError } from '@connectrpc/connect';
import { useThrottleFn } from '@vueuse/core';
import DataNoDataBlock from '~/components/partials/data/DataNoDataBlock.vue';
import DataPendingBlock from '~/components/partials/data/DataPendingBlock.vue';
import DataErrorBlock from '~/components/partials/data/DataErrorBlock.vue';
import {
    CancelNetworkTestResponse,
    GetNetworkTestStatusResponse,
    NetworkTestOutputFormat,
    StartNetworkTestResponse,
} from '~~/gen/ts/api/services/troubleshooting/v1/troubleshooting_pb';
import GenericDivider from '~/components/partials/GenericDivider.vue';

const { $grpc } = useNuxtApp();

const {
    data: testStatus,
    error,
    pending,
    refresh,
} = useLazyAsyncData('cluster-network_test_status', async () => getNetworkTestStatus());

async function getNetworkTestStatus(): Promise<GetNetworkTestStatusResponse | undefined> {
    try {
        return await $grpc.getTroubleshootingClient().getNetworkTestStatus({});
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
}

async function startNetworkTest(): Promise<StartNetworkTestResponse | undefined> {
    try {
        const resp = await $grpc.getTroubleshootingClient().startNetworkTest({
            hostNetwork: false,
            outputFormat: NetworkTestOutputFormat.CSV,
        });

        if (testStatus.value !== undefined && testStatus.value !== null) {
            testStatus.value.running = true;
            testStatus.value.complete = false;
            testStatus.value.logs = '';
        }

        refresh();

        return resp;
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
}

async function cancelNetworkTest(): Promise<CancelNetworkTestResponse | undefined> {
    try {
        const resp = await $grpc.getTroubleshootingClient().cancelNetworkTest({});

        if (testStatus.value !== undefined && testStatus.value !== null) {
            testStatus.value.running = false;
        }

        return resp;
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
}

const canSubmit = ref(true);
const onSubmitThrottle = useThrottleFn(async () => {
    canSubmit.value = false;
    if (testStatus.value?.running) {
        cancelNetworkTest().finally(() => setTimeout(() => (canSubmit.value = true), 400));
    } else {
        startNetworkTest().finally(() => setTimeout(() => (canSubmit.value = true), 400));
    }
}, 1000);

async function getNetworkTestResults(): Promise<void> {
    try {
        const resp = await $grpc.getTroubleshootingClient().getNetworkTestResults({});

        const blob = new Blob([resp.fileContents], { type: resp.fileType });
        saveAs(blob, resp.fileName);
    } catch (e) {
        $grpc.handleError(e as ConnectError);
    }
}
</script>

<template>
    <div class="p-2">
        <h2 class="text-xl q-mt-lg text-center">Ancientt Network Test</h2>

        <div class="flex items-center gap-1">
            <button
                type="button"
                class="flex justify-center w-full px-3 py-2 text-sm font-semibold transition-colors rounded-md text-neutral focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2"
                :disabled="!canSubmit"
                :class="[
                    !canSubmit
                        ? 'disabled bg-base-500 hover:bg-base-400 focus-visible:outline-base-500'
                        : 'bg-primary-500 hover:bg-primary-400 focus-visible:outline-primary-500',
                ]"
                @click="onSubmitThrottle"
            >
                <template v-if="!testStatus">Loading Status ...</template>
                <template v-else-if="!testStatus.running"> Start Network Test </template>
                <template v-else> Stop Network Test </template>
            </button>
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
            <DataNoDataBlock v-else-if="testStatus === null || testStatus === undefined" />

            <template v-else>
                <p class="font-semibold text-gray-700 text-sm">
                    <template v-if="testStatus">
                        <span v-if="testStatus.running"> Network Test is currently running. </span>
                        <span v-else> No network Test are currently running. </span>
                    </template>
                    <template v-else> N/A </template>
                </p>

                <template v-if="testStatus.logs.length > 0">
                    <GenericDivider label="Network Test Logs" />

                    <textarea
                        :value="testStatus.logs.split('\n').reverse().join('\n')"
                        disabled
                        rows="10"
                        name="logs"
                        class="mt-2 block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                    />
                </template>

                <button
                    v-if="testStatus.complete"
                    type="button"
                    class="mt-4 flex-1 flex justify-center w-full px-3 py-2 text-sm font-semibold transition-colors rounded-md text-neutral focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2"
                    :disabled="!canSubmit"
                    :class="[
                        !canSubmit
                            ? 'disabled bg-base-500 hover:bg-base-400 focus-visible:outline-base-500'
                            : 'bg-primary-500 hover:bg-primary-400 focus-visible:outline-primary-500',
                    ]"
                    @click="getNetworkTestResults()"
                >
                    Download Test Results
                </button>
            </template>
        </div>
    </div>
</template>
