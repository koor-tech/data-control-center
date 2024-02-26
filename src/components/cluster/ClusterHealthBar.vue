<script lang="ts" setup>
import { clusterHealthToBGColor, clusterHealthToTextColor } from '~/components/cluster/helpers';
import GenericContainer from '~/components/partials/GenericContainer.vue';
import { ClusterHealth } from '~~/gen/ts/api/resources/ceph/v1/stats_pb';

defineProps<{
    clusterId: string;
    health: ClusterHealth;
}>();
</script>

<template>
    <GenericContainer>
        <header class="w-full">
            <div class="mx-auto grid grid-cols-2 items-center justify-center">
                <div class="inline-flex items-center justify-start gap-2">
                    <div
                        class="h-4 w-4 rounded-full p-1 animate-pulse"
                        :class="[clusterHealthToBGColor(health), clusterHealthToTextColor(health)]"
                    >
                        <div class="h-2 w-2 rounded-full bg-current" />
                    </div>
                    <p class="font-semibold text-gray-700 text-sm">
                        {{ ClusterHealth[health] }}
                    </p>
                </div>
                <p class="text-sm leading-7 text-gray-900 justify-self-end">
                    Cluster ID: <code>{{ clusterId }}</code>
                </p>
            </div>
        </header>
    </GenericContainer>
</template>
