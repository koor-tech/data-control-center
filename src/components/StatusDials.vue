<script setup lang="ts">
import { ClusterRadar } from '~~/gen/ts/api/resources/stats/stats_pb';
import DialIndicator from './DialIndicator.vue';

const props = defineProps<{
    radar: ClusterRadar;
}>();

const dials = [
    { value: props.radar.clusterHealth, min: 0, max: 100, caption: 'Cluster Health' },
    { value: props.radar.nodesHealth, min: 0, max: 100, caption: 'Nodes Health' },
    { value: props.radar.capacityAvailable, min: 0, max: 100, caption: 'Capacity Available' },
    { value: props.radar.stability, min: 0, max: 100, caption: 'Stability' },
    { value: props.radar.reliability, min: 0, max: 100, caption: 'Reliability' },
];
</script>

<template>
    <div class="text-h3 text-center">Key Indicators</div>
    <div class="grid grid-cols-5 gap-4">
        <DialIndicator
            v-for="dial in dials"
            :key="dial.caption"
            :value="dial.value"
            :min="dial.min"
            :max="dial.max"
            :caption="dial.caption"
        />
    </div>
</template>
