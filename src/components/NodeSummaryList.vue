<template>
  <div class="text-h3 text-center">Storage Nodes</div>
  <ul>
    <div v-for="(item, index) in items" :key="item.name">
      <NodeSummaryItem :name="item.name" :description="item.description" :capacity="item.capacity" :health="item.health"
        :last-ping="item.lastPing" />
      <q-separator v-if="index + 1 < items.length" spaced inset />
    </div>
  </ul>
</template>

<script setup>
import NodeSummaryItem from './NodeSummaryItem.vue'
const items = [
  {
    name: 'rack1-slots-1-3',
    description: '3 drives, 9 OSDs, 10 GiB Network',
    capacity: '180 GB',
    health: 0,
    lastPing: '5 min',
  },
  {
    name: 'rack1-slots-4-7',
    description: '4 drives, 12 OSDs, 10 GiB Network',
    capacity: '1 TB',
    health: 1,
    lastPing: '5 min',
  },
  {
    name: 'rack1-slots-8-11',
    description: '4 drives, 12 OSDs, 10 GiB Network',
    capacity: '1 TB',
    health: 2,
    lastPing: '7 min',
  },
  {
    name: 'rack2-slots-1-6',
    description: '6 drives, 18 OSDs, 10 GiB Network',
    capacity: '3.6 TB',
    health: 0,
    lastPing: '10 min',
  },
]
</script>

<style lang="scss" scoped></style>
