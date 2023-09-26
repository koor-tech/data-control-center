<script setup lang="ts">
import { ClusterRadar } from '~~/gen/ts/api/resources/stats/v1/stats_pb';
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
    <div class="bg-white rounded-lg shadow dark:bg-gray-800 px-2 py-2">
        <div class="flex justify-center mb-3 items-center">
            <h3 class="text-xl font-bold leading-none text-gray-900 dark:text-white pr-1">Key Indicators</h3>
        </div>
        <div class="h-full grid grid-cols-5 gap-4 justify-center justify-items-center content-center">
            <DialIndicator
                v-for="dial in dials"
                :key="dial.caption"
                :value="dial.value"
                :min="dial.min"
                :max="dial.max"
                :caption="dial.caption"
            />
        </div>
    </div>
</template>
