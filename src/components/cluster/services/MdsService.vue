<script setup lang="ts">
import { FileDocumentIcon } from 'mdi-vue3';
import type { MdsService } from '~~/gen/ts/api/resources/stats/v1/stats_pb';
import ServiceLiItem from '~/components/cluster/services/ServiceLiItem.vue';
import ServiceBox from '~/components/cluster/services/ServiceBox.vue';

defineProps<{
    stats: MdsService;
}>();
</script>

<template>
    <ServiceBox :icon="markRaw(FileDocumentIcon)">
        <template #default> MDS </template>
        <template #content>
            <ServiceLiItem>
                <template #default> Up </template>
                <template #content>
                    {{ stats.daemonsUp }}
                </template>
            </ServiceLiItem>
            <ServiceLiItem>
                <template #default> Hot Standby </template>
                <template #content>
                    {{ stats.hotStandbyCount }}
                </template>
            </ServiceLiItem>
        </template>
    </ServiceBox>
</template>
