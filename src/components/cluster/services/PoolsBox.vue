<script setup lang="ts">
import { PoolIcon } from 'mdi-vue3';
import type { Pools } from '~~/gen/ts/api/resources/ceph/v1/stats_pb';
import ServiceLiItem from '~/components/cluster/services/ServiceLiItem.vue';
import ServiceBox from '~/components/cluster/services/ServiceBox.vue';

defineProps<{
    stats: Pools;
}>();
</script>

<template>
    <ServiceBox :icon="markRaw(PoolIcon)">
        <template #default> Pools </template>
        <template #content>
            <ServiceLiItem>
                <template #default> Count </template>
                <template #content>
                    {{ stats.pools }}
                </template>
            </ServiceLiItem>
        </template>
    </ServiceBox>
</template>
