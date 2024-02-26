<script setup lang="ts">
import type { ApexOptions } from 'apexcharts';
// @ts-expect-error types are currently broken, see https://github.com/apexcharts/vue3-apexcharts/issues/95
import VueApexCharts from 'vue3-apexcharts';
import GenericContainer from '~/components/partials/GenericContainer.vue';
import { ClusterRadar } from '~~/gen/ts/api/resources/ceph/v1/stats_pb';

const props = defineProps<{
    radar: ClusterRadar;
}>();

const options: ApexOptions = {
    chart: {
        type: 'radar',
        toolbar: {
            show: false,
        },
        fontFamily: 'Inter var',
    },
    colors: ['#00a77d'],
    dataLabels: {
        enabled: false,
    },
    plotOptions: {
        radar: {
            polygons: {
                strokeColors: '#e8e8e8',
                fill: {
                    colors: ['#f8f8f8', '#fff'],
                },
            },
        },
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
        tickAmount: 5,
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

const renderChart = ref(false);
onMounted(() => (renderChart.value = true));
</script>

<template>
    <GenericContainer>
        <template #title> Cluster Radar </template>
        <template #default>
            <div class="h-[12rem] max-h-[12rem] overflow-hidden">
                <!-- Chart -->
                <VueApexCharts
                    v-if="renderChart"
                    height="200px"
                    width="400px"
                    type="radar"
                    :options="options"
                    :series="series"
                />
            </div>
        </template>
    </GenericContainer>
</template>
