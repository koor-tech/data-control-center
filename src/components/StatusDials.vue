<template>
  <div class="text-h3 text-center">Key Indicators</div>
  <div class="q-pa-md flex flex-center">
    <DialIndicator :value="37" :min="0" :max="100" caption="Cap Used %" />
    <DialIndicator :value="-12" :min="-50" :max="50" caption="Balance" />
    <DialIndicator :value="9" :min="0" :max="12" caption="Healthy Nodes" />
    <DialIndicator :value="5" :min="0" :max="8" caption="Pizzas Baking" />
    <DialIndicator :value="3" :min="0" :max="21" caption="Crying Children" />
    <DialIndicator :value="127" :min="0" :max="1000" caption="Birthdays" />
    <DialIndicator :value="87" :min="0" :max="100" caption="Likes %" />
    <DialIndicator :value="20" :min="-50" :max="50" caption="Utilization" />
    <DialIndicator :value="4" :min="0" :max="30" caption="Fridays per Month" />
  </div>
</template>

<script setup>
import DialIndicator from './DialIndicator.vue'
</script>

<style lang="scss" scoped></style>
