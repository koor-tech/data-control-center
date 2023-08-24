<script lang="ts" setup>
import { Chart as ChartJS, RadialLinearScale, PointElement, LineElement, Filler, Tooltip, Legend } from 'chart.js';
import { Radar } from 'vue-chartjs';
import * as chartConfig from './radarChartConfig';

ChartJS.register(RadialLinearScale, PointElement, LineElement, Filler, Tooltip, Legend);
</script>

<template>
    <Radar :data="chartConfig.data" :options="chartConfig.options" :style="chartConfig.style" />
</template>
