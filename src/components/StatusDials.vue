<template>
  <div class="text-h3 text-center">Key Indicators</div>
  <div class="q-pa-md flex flex-center">
    <DialIndicator v-for="dial in dials" :key="dial.caption" :value="dial.value" :min="dial.min" :max="dial.max"
      :caption="dial.caption" />
  </div>
</template>

<script setup lang="ts">
import DialIndicator from './DialIndicator.vue';

const dials = [
  { value: 37, min: 0, max: 100, caption: 'Cap Used %', },
  { value: -12, min: -50, max: 50, caption: 'Balance', },
  { value: 9, min: 0, max: 1, caption: 'Healthy Nodes', },
  { value: 5, min: 0, max: 8, caption: 'Pizzas Baking', },
  { value: 3, min: 0, max: 21, caption: 'Crying Children', },
  { value: 127, min: 0, max: 1000, caption: 'Birthdays', },
  { value: 87, min: 0, max: 100, caption: 'Likes %', },
  { value: 20, min: -50, max: 50, caption: 'Utilization', },
  { value: 4, min: 0, max: 30, caption: 'Fridays per Month', },
];
</script>

<style lang="scss" scoped></style>
