<script setup lang="ts">
import BarChart from '../components/BarChartWidget.vue';
import LineChart from '../components/LineChartWidget.vue';

useHead({
    title: 'Charts',
});
definePageMeta({
    title: 'Charts',
    requiresAuth: true,
});
</script>

<template>
    <div>
        <div class="text-h2 q-mt-lg text-center">Dashboard</div>
        <div class="flex flex-center">
            <div>
                <LineChart />
                <div class="text-center text-caption text-bold">Random Data</div>
            </div>
            <div>
                <BarChart />
                <div class="text-center text-caption text-bold">Disc Utilization</div>
            </div>
        </div>
    </div>
</template>
