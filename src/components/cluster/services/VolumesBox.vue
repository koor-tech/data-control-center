<script setup lang="ts">
import { HarddiskIcon } from 'mdi-vue3';
import type { Data } from '~~/gen/ts/api/resources/ceph/v1/stats_pb';
import ServiceLiItem from '~/components/cluster/services/ServiceLiItem.vue';
import ServiceBox from '~/components/cluster/services/ServiceBox.vue';

defineProps<{
    stats: Data;
}>();
</script>

<template>
    <ServiceBox :icon="markRaw(HarddiskIcon)">
        <template #default> Volumes </template>
        <template #content>
            <ServiceLiItem>
                <template #default> Count </template>
                <template #content>
                    {{ stats.volumes }}
                </template>
            </ServiceLiItem>
        </template>
    </ServiceBox>
</template>
