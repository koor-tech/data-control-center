<script setup lang="ts">
import { TableIcon } from 'mdi-vue3';
import type { OsdService } from '~~/gen/ts/api/resources/ceph/v1/stats_pb';
import ServiceLiItem from '~/components/cluster/services/ServiceLiItem.vue';
import ServiceBox from '~/components/cluster/services/ServiceBox.vue';
import GenericTime from '~/components/partials/elements/GenericTime.vue';

defineProps<{
    stats: OsdService;
}>();
</script>

<template>
    <ServiceBox :icon="markRaw(TableIcon)">
        <template #default> OSDs </template>
        <template #content>
            <ServiceLiItem>
                <template #default> Total </template>
                <template #content>
                    {{ stats.osdCount }}
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> In </template>
                <template #content>
                    {{ stats.osdIn }}
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> Up </template>
                <template #content>
                    {{ stats.osdUp }}
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> In Updated </template>
                <template #content>
                    <GenericTime :value="stats.osdInUpdatedSince" />
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> Up Updated </template>
                <template #content>
                    <GenericTime :value="stats.osdUpUpdatedSince" />
                </template>
            </ServiceLiItem>
        </template>
    </ServiceBox>
</template>
