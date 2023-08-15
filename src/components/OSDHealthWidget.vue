<template>
  <div class="text-h3 text-center">Pod Status</div>
  <div class="q-pa-md">
    <q-table title="Rook Ceph Pods" :rows="rows" :columns="columns" row-key="name" />
  </div>
</template>

<script setup lang="ts">
const columns = [
  { name: 'name', label: 'Name', field: 'name', align: 'left', sortable: true },
  { name: 'ready', label: 'Ready', field: 'ready', align: 'center' },
  { name: 'status', label: 'Status', field: 'status', align: 'center' },
  { name: 'restarts', label: 'Restarts', field: 'restarts' },
  { name: 'age', label: 'Age', field: 'age' },
]

const rows = [
  {
    name: 'rook-ceph-mon1-dw8wr',
    ready: '1/1',
    status: 'Running',
    restarts: 0,
    age: '17h',
  },
  {
    name: 'rook-ceph-mon1-kio3k',
    ready: '1/1',
    status: 'Running',
    restarts: 0,
    age: '17h',
  },
  {
    name: 'rook-ceph-osd-cdefg',
    ready: '1/1',
    status: 'Running',
    restarts: 0,
    age: '12h',
  },
  {
    name: 'rook-ceph-osd-bcdef',
    ready: '1/1',
    status: 'Running',
    restarts: 0,
    age: '12h',
  },
  {
    name: 'rook-ceph-osd-abcde',
    ready: '1/1',
    status: 'Running',
    restarts: 0,
    age: '8h',
  },
]
</script>

<style lang="scss" scoped></style>
