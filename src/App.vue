<template>
  <router-view />
</template>

<script setup lang="ts">
import stats from '../src/services/stats';

async function main() {
    const res = stats;
    console.log(res.prototype.getStats());
}
void main();
</script>
