<script setup lang="ts">
import { AlertBoxIcon } from 'mdi-vue3';
import { type Crash } from '~~/gen/ts/api/resources/ceph/v1/stats_pb';
import ServiceLiItem from '~/components/cluster/services/ServiceLiItem.vue';
import ServiceBox from '~/components/cluster/services/ServiceBox.vue';

defineProps<{
    alerts: Crash[];
}>();
</script>

<template>
    <ServiceBox :icon="markRaw(AlertBoxIcon)" :color="alerts.length > 0 ? 'bg-warn-500' : 'bg-success-500'">
        <template #default> Alerts </template>
        <template #content>
            <template v-if="alerts.length > 0"> </template>
            <ServiceLiItem v-for="(alert, idx) in alerts" :key="idx">
                <template #content> {{ alert.description }} </template>
                <!-- TODO show bullet points infront of alerts -->
            </ServiceLiItem>
        </template>
    </ServiceBox>
</template>
