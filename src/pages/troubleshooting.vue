<script setup lang="ts">
import { ConnectError } from '@connectrpc/connect';
import { useThrottleFn } from '@vueuse/core';
import GenericDivider from '~/components/partials/GenericDivider.vue';
import { useClusterStore } from '~/store/cluster';

useHead({
    title: 'Troubleshooting',
});
definePageMeta({
    title: 'Troubleshooting',
    requiresAuth: true,
});

const { $grpc } = useNuxtApp();

const clusterStore = useClusterStore();

const { data: report, refresh } = useLazyAsyncData(
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

const canSubmit = ref(true);
const onSubmitThrottle = useThrottleFn(async (_) => {
    canSubmit.value = false;
    await refresh().finally(() => setTimeout(() => (canSubmit.value = true), 400));
}, 1000);
</script>

<template>
    <div class="p-2">
        <h2 class="text-xl q-mt-lg text-center">Troubleshooting</h2>
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
        <div v-if="report" class="flex">
            <GenericDivider />
            <textarea :value="report.report" />
        </div>
    </div>
</template>
