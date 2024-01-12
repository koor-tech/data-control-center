<script setup lang="ts">
import { GroupIcon } from 'mdi-vue3';
import type { PGs } from '~~/gen/ts/api/resources/ceph/v1/stats_pb';
import ServiceLiItem from '~/components/cluster/services/ServiceLiItem.vue';
import ServiceBox from '~/components/cluster/services/ServiceBox.vue';

defineProps<{
    stats: PGs;
}>();
</script>

<template>
    <ServiceBox :icon="markRaw(GroupIcon)">
        <template #default> Placement Groups </template>
        <template #content>
            <ServiceLiItem>
                <template #default> Active+Clean </template>
                <template #content>
                    {{ stats.activeClean }}
                </template>
            </ServiceLiItem>
        </template>
    </ServiceBox>
</template>
