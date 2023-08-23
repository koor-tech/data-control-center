<template>
  <q-item>
    <q-item-section>
      <li>{{ name }}</li>
      <li caption>{{ description }}</li>
    </q-item-section>

    <q-item-section side top>
      <li caption>{{ lastPing }} ago</li>
      <q-icon name="star" :color="decodeHealth(health)" />
      <q-badge color="teal" :label="capacity" />
    </q-item-section>
  </q-item>
</template>

<script setup lang="ts">
defineProps({
  name: String,
  description: String,
  lastPing: String,
  capacity: String,
  health: Number,
})

const decodeHealth = (code?: number) => {
  switch (code) {
    case 1:
      return 'yellow'
    case 2:
      return 'red'
    default:
      return 'green'
  }
}
</script>

<style lang="scss" scoped></style>
