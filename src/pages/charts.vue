<template>
  <q-page>
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
  </q-page>
</template>

<script setup lang="ts">
import LineChart from '../components/LineChartWidget.vue';
import BarChart from '../components/BarChartWidget.vue';
</script>
