<script setup lang="ts">
import { MonitorIcon } from 'mdi-vue3';
import type { Usage } from '~~/gen/ts/api/resources/stats/v1/stats_pb';
import ServiceLiItem from '~/components/cluster/services/ServiceLiItem.vue';
import ServiceBox from '~/components/cluster/services/ServiceBox.vue';
import { formatBytesBigInt } from '~/utils/strings';

defineProps<{
    stats: Usage;
}>();
</script>

<template>
    <ServiceBox :icon="markRaw(MonitorIcon)">
        <template #default> Usage </template>
        <template #content>
            <ServiceLiItem>
                <template #default> Total </template>
                <template #content>
                    {{ formatBytesBigInt(stats.total) }}
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> Available </template>
                <template #content>
                    {{ formatBytesBigInt(stats.available) }}
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> Used </template>
                <template #content>
                    {{ formatBytesBigInt(stats.used) }}
                </template>
            </ServiceLiItem>
        </template>
    </ServiceBox>
</template>
