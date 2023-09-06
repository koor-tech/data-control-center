<template>
    <div class="bg-white rounded-lg shadow dark:bg-gray-800 px-2 py-2">
        <div class="flex justify-center mb-3 items-center">
            <h5 class="text-xl font-bold leading-none text-gray-900 dark:text-white pr-1">Cluster Radar</h5>
        </div>

        <!-- Chart -->
        <div>
            <VueApexCharts class="flex justify-center py-1" height="250px" width="250px" type="radar" :options="options"
                :series="series" />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ApexOptions } from 'apexcharts';
import VueApexCharts from 'vue3-apexcharts';

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
        categories: ['Cluster Health', 'Nodes Health', 'Capacity Used', 'Stability', 'Reliability'],
        min: 0,
        max: 100,
    },
    yaxis: {
        min: 0,
        max: 100,
        tickAmount: 4,
    },
};

const series = [
    {
        name: 'Cluster 1',
        data: [100, 50, 75, 10, 35],
    },
];
</script>
