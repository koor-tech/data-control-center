<template>
  <div class="px-4">
    <div class="text-h2 q-mt-lg text-center">Control Panel</div>
    <div class="flex">
      <div class="control-panel">
        <div class="q-pa-md">
          <ClusterSettings label="Preferences" />
        </div>
      </div>
      <div class="control-panel">
        <div class="q-pa-md">
          <ClusterSettings label="Dangerous" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ClusterSettings from '../components/ClusterSettings.vue';
</script>

<style lang="scss">
.control-panel {
  width: 500px;
}
</style>
