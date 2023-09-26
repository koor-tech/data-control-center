<script setup lang="ts">
import BarChart from '~/components/BarChartWidget.vue';
import LineChart from '~/components/LineChartWidget.vue';
import Container from '~/components/partials/Container.vue';

useHead({
    title: 'Charts',
});
definePageMeta({
    title: 'Charts',
    requiresAuth: true,
});
</script>

<template>
    <div class="p-2">
        <Container>
            <template v-slot:title> Random Data </template>
            <template v-slot:default>
                <LineChart />
            </template>
        </Container>
        <Container>
            <template v-slot:title> Disc Utilization </template>
            <template v-slot:default>
                <BarChart />
            </template>
        </Container>
    </div>
</template>
