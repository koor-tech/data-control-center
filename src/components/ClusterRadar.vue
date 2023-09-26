<script setup lang="ts">
import { ApexOptions } from 'apexcharts';
import VueApexCharts from 'vue3-apexcharts';
import { ClusterRadar } from '~~/gen/ts/api/resources/stats/v1/stats_pb';
import Container from './partials/Container.vue';

const props = defineProps<{
    radar: ClusterRadar;
}>();

const options: ApexOptions = {
    chart: {
        type: 'radar',
        toolbar: {
            show: false,
        },
    },
    colors: ['#00a77d'],
    dataLabels: {
        enabled: false,
    },
    fill: {
        type: 'gradient',
        gradient: {
            shade: 'dark',
            type: 'horizontal',
            gradientToColors: ['#87D4F9'],
        },
    },
    xaxis: {
        categories: ['Cluster Health', 'Nodes Health', 'Capacity Available', 'Stability', 'Reliability'],
        min: 0,
        max: 100,
    },
    yaxis: {
        min: 0,
        max: 100,
        tickAmount: 4,
    },
};

const series = ref([
    {
        name: 'Cluster Radar',
        data: [
            props.radar.clusterHealth,
            props.radar.nodesHealth,
            props.radar.capacityAvailable,
            props.radar.stability,
            props.radar.reliability,
        ],
    },
]);
</script>

<template>
    <Container>
        <template v-slot:title> Cluster Radar </template>
        <template>
            <!-- Chart -->
            <VueApexCharts
                class="flex justify-center py-1"
                height="250px"
                width="300px"
                type="radar"
                :options="options"
                :series="series"
            />
        </template>
    </Container>
</template>
