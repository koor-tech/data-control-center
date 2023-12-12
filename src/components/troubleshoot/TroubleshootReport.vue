<script lang="ts" setup>
import { ConnectError } from '@connectrpc/connect';
import { useClipboard, useThrottleFn } from '@vueuse/core';
import GenericDivider from '~/components/partials/GenericDivider.vue';
import DataErrorBlock from '~/components/partials/data/DataErrorBlock.vue';
import DataNoDataBlock from '~/components/partials/data/DataNoDataBlock.vue';
import DataPendingBlock from '~/components/partials/data/DataPendingBlock.vue';
import { useClusterStore } from '~/store/cluster';
import { useNotificationsStore } from '~/store/notifications';

const { $grpc } = useNuxtApp();

const clusterStore = useClusterStore();

const {
    data: report,
    error,
    pending,
    refresh,
} = useLazyAsyncData(
    'cluster-troubleshoot_report',
    async () => {
        try {
            return await clusterStore.getTroubleshootReport();
        } catch (e) {
            $grpc.handleError(e as ConnectError);
        }
    },
    { immediate: false },
);

const notifications = useNotificationsStore();
const { copy } = useClipboard();

function copyToClipboard(): void {
    if (!report.value || !report.value.report) {
        return;
    }

    copy(report.value.report);

    notifications.dispatchNotification({
        title: 'Copied to clipboard',
        content: 'Troubleshoot report has been copied to your clipboard.',
        type: 'success',
    });
}

const canSubmit = ref(true);
const onSubmitThrottle = useThrottleFn(async (_) => {
    canSubmit.value = false;
    await refresh().finally(() => setTimeout(() => (canSubmit.value = true), 400));
}, 1000);
</script>

<template>
    <div class="p-2">
        <h2 class="text-xl q-mt-lg text-center">Troubleshooting Report</h2>
        <div class="flex">
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
                Generate Troubleshoot Report
            </button>
        </div>
        <div v-if="report" class="mt-4 flex flex-col">
            <GenericDivider label="Report" />

            <DataErrorBlock v-if="error" :retry="refresh" :message="error.message" />
            <DataPendingBlock v-else-if="pending" message="Loading Troubleshoot report" />
            <DataNoDataBlock v-else-if="report === null || report.report === undefined" />

            <template v-else>
                <textarea
                    :value="report.report"
                    disabled
                    rows="10"
                    name="comment"
                    class="mt-2 block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
                <button
                    type="button"
                    class="mt-2 flex justify-center bg-primary-500 hover:bg-primary-400 focus-visible:outline-primary-500 w-full px-3 py-2 text-sm font-semibold transition-colors rounded-md text-neutral focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2"
                    @click="copyToClipboard()"
                >
                    Copy to Clipboard
                </button>
            </template>
        </div>
    </div>
</template>
