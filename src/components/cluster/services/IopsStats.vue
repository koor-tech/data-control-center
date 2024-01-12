<script setup lang="ts">
import { SwapVerticalIcon } from 'mdi-vue3';
import type { IOPS } from '~~/gen/ts/api/resources/ceph/v1/stats_pb';
import ServiceLiItem from '~/components/cluster/services/ServiceLiItem.vue';
import ServiceBox from '~/components/cluster/services/ServiceBox.vue';
import { formatBytesBigInt } from '~/utils/strings';

defineProps<{
    stats: IOPS;
}>();
</script>

<template>
    <ServiceBox :icon="markRaw(SwapVerticalIcon)">
        <template #default> IOPS </template>
        <template #content>
            <ServiceLiItem>
                <template #default> Read </template>
                <template #content>
                    {{ formatBytesBigInt(stats.clientRead) }}
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> Write </template>
                <template #content>
                    {{ formatBytesBigInt(stats.clientWrite) }}
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> Read Ops </template>
                <template #content> {{ stats.clientReadOps }} p/s </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> Write Ops </template>
                <template #content> {{ stats.clientWrite }} p/s </template>
            </ServiceLiItem>
        </template>
    </ServiceBox>
</template>
