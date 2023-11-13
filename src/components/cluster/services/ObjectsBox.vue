<script setup lang="ts">
import { CubeIcon } from 'mdi-vue3';
import type { Objects } from '~~/gen/ts/api/resources/stats/v1/stats_pb';
import ServiceLiItem from '~/components/cluster/services/ServiceLiItem.vue';
import ServiceBox from '~/components/cluster/services/ServiceBox.vue';
import { formatBytesBigInt } from '~/utils/strings';

defineProps<{
    stats: Objects;
}>();
</script>

<template>
    <ServiceBox :icon="markRaw(CubeIcon)">
        <template #default> Objects </template>
        <template #content>
            <ServiceLiItem>
                <template #default> Total </template>
                <template #content>
                    {{ stats.objectCount }}
                </template>
            </ServiceLiItem>

            <ServiceLiItem>
                <template #default> Size </template>
                <template #content>
                    {{ formatBytesBigInt(stats.size) }}
                </template>
            </ServiceLiItem>
        </template>
    </ServiceBox>
</template>
