<script lang="ts" setup>
import type { ConnectError } from '@connectrpc/connect';
import { useThrottleFn } from '@vueuse/core';
import { NetworkTestOutputFormat, StartNetworkTestResponse } from '~~/gen/ts/api/services/cluster/v1/cluster_pb';

const { $grpc } = useNuxtApp();

// TODO

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
    <div>
        <button type="button" @click="onSubmitThrottle()">Start</button>
        <!-- TODO -->
    </div>
</template>
