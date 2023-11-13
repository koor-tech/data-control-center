<script setup lang="ts">
import { ClipboardIcon } from 'mdi-vue3';
import type { MgrService } from '~~/gen/ts/api/resources/stats/v1/stats_pb';
import ServiceLiItem from '~/components/cluster/services/ServiceLiItem.vue';
import ServiceBox from '~/components/cluster/services/ServiceBox.vue';
import Time from '~/components/partials/elements/Time.vue';
import GenericBadge from '~/components/partials/GenericBadge.vue';

defineProps<{
    stats: MgrService;
}>();
</script>

<template>
    <ServiceBox :icon="markRaw(ClipboardIcon)">
        <template #default> MGR </template>
        <template #content>
            <ServiceLiItem>
                <template #default> Standby MGRs </template>
                <template #content>
                    <GenericBadge v-for="daemon in stats.standbys" :key="daemon">
                        {{ daemon }}
                    </GenericBadge>
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> Updated Since </template>
                <template #content>
                    <Time :value="stats.updatedSince" />
                </template>
            </ServiceLiItem>
        </template>
    </ServiceBox>
</template>
