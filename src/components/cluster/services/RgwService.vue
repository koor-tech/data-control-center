<script setup lang="ts">
import { ArchiveArrowDownIcon } from 'mdi-vue3';
import type { RgwService } from '~~/gen/ts/api/resources/ceph/v1/stats_pb';
import ServiceLiItem from '~/components/cluster/services/ServiceLiItem.vue';
import ServiceBox from '~/components/cluster/services/ServiceBox.vue';

defineProps<{
    stats: RgwService;
}>();
</script>

<template>
    <ServiceBox :icon="markRaw(ArchiveArrowDownIcon)">
        <template #default> RGW </template>
        <template #content>
            <ServiceLiItem>
                <template #default> Active </template>
                <template #content>
                    {{ stats.activeDaemon }}
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> Hosts </template>
                <template #content>
                    {{ stats.hostCount }}
                </template>
            </ServiceLiItem>
        </template>
    </ServiceBox>
</template>
