<script setup lang="ts">
import type { ApexOptions } from 'apexcharts';
// @ts-expect-error types are currently broken, see https://github.com/apexcharts/vue3-apexcharts/issues/95
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
        fontFamily: 'Inter var',
    },
    title: {
        text: props.caption,
        align: 'center',
        style: {
            fontWeight: 500,
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
    <VueApexCharts
        class="flex justify-center py-1"
        height="150px"
        width="100%"
        type="radialBar"
        :options="options"
        :series="series"
    />
</template>
