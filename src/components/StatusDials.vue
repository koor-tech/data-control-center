<template>
  <div class="text-h3 text-center">Key Indicators</div>
  <div class="grid grid-cols-5 gap-4">
    <DialIndicator v-for="dial in dials" :key="dial.caption" :value="dial.value" :min="dial.min" :max="dial.max"
      :caption="dial.caption" />
  </div>
</template>

<script setup lang="ts">
import DialIndicator from './DialIndicator.vue';

const dials = [
  { value: 37, min: 0, max: 100, caption: 'Cluster Health', },
  { value: 9, min: 0, max: 1, caption: 'Nodes Health', },
  { value: 50, min: -50, max: 50, caption: 'Capacity Used', },
  { value: 5, min: 0, max: 8, caption: 'Stability', },
  { value: 3, min: 0, max: 21, caption: 'Reliability', },
];
</script>
