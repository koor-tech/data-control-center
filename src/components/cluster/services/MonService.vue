<script setup lang="ts">
import { MonitorIcon } from 'mdi-vue3';
import type { MonService } from '~~/gen/ts/api/resources/ceph/v1/stats_pb';
import GenericBadge from '~/components/partials/GenericBadge.vue';
import ServiceLiItem from '~/components/cluster/services/ServiceLiItem.vue';
import ServiceBox from '~/components/cluster/services/ServiceBox.vue';
import GenericTime from '~/components/partials/elements/GenericTime.vue';

defineProps<{
    stats: MonService;
}>();
</script>

<template>
    <ServiceBox :icon="markRaw(MonitorIcon)">
        <template #default> MONs </template>
        <template #content>
            <ServiceLiItem>
                <template #default> Daemons </template>
                <template #content>
                    <GenericBadge v-for="daemon in stats.quorum" :key="daemon">
                        {{ daemon }}
                    </GenericBadge>
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> Quorum </template>
                <template #content>
                    <GenericBadge :color="stats.quorum.length >= stats.daemonCount ? 'success' : 'error'">
                        {{ stats.quorum.length >= stats.daemonCount ? 'OK' : 'DANGER' }}
                    </GenericBadge>
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> Created Since </template>
                <template #content>
                    <GenericTime :value="stats.createdSince" />
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> Updated Since </template>
                <template #content>
                    <GenericTime :value="stats.updatedSince" />
                </template>
            </ServiceLiItem>
        </template>
    </ServiceBox>
</template>
