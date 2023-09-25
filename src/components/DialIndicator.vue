<script setup lang="ts">
import { ApexOptions } from 'apexcharts';
import VueApexCharts from 'vue3-apexcharts';

const props = defineProps<{
    min: number;
    max: number;
    value: number;
    caption: string;
    label?: string;
    labels?: string[];
}>();

const options: ApexOptions = {
    chart: {
        type: 'radialBar',
        toolbar: {
            show: false,
        },
    },
    colors: ['#00a77d'],
    plotOptions: {
        radialBar: {
            startAngle: -135,
            endAngle: 135,
            track: {
                background: '#41444e',
                startAngle: -135,
                endAngle: 135,
            },
            dataLabels: {
                name: {
                    show: false,
                },
                value: {
                    fontSize: '24px',
                    show: true,
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
            stops: [props.min, props.max],
        },
    },
    stroke: {
        lineCap: 'butt',
    },
    yaxis: {
        min: props.min,
        max: props.max,
    },
    grid: {
        padding: {
            left: 0,
            right: 0,
        },
    },
};

const series = [props.value];
</script>

<template>
    <div class="bg-white rounded-lg shadow dark:bg-gray-800 px-2 py-2">
        <div class="flex justify-center mb-3 items-center">
            <h5 class="text-xl font-bold leading-none text-gray-900 dark:text-white pr-1">{{ caption }}</h5>
        </div>

        <!-- Chart -->
        <div>
            <VueApexCharts
                class="flex justify-center py-1"
                height="150px"
                width="150px"
                type="radialBar"
                :options="options"
                :series="series"
            />
        </div>
    </div>
</template>
