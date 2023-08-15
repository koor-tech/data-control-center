<template>
  <q-list bordered padding>
    <q-item-label header>{{ label || 'Setting' }}</q-item-label>

    <q-item>
      <q-item-section side>
        <q-icon color="teal" name="volume_down" />
      </q-item-section>
      <q-item-section>
        <q-slider v-model="settings.volume" :min="0" :max="10" label color="teal" />
      </q-item-section>
      <q-item-section side>
        <q-icon color="teal" name="volume_up" />
      </q-item-section>
    </q-item>

    <q-item>
      <q-item-section side>
        <q-icon color="deep-orange" name="brightness_medium" />
      </q-item-section>
      <q-item-section>
        <q-slider v-model="settings.brightness" :min="0" :max="10" label color="deep-orange" />
      </q-item-section>
    </q-item>

    <q-item>
      <q-item-section side>
        <q-icon color="primary" name="mic" />
      </q-item-section>
      <q-item-section>
        <q-slider v-model="settings.mic" :min="0" :max="50" label />
      </q-item-section>
    </q-item>
  </q-list>
</template>

<script setup lang="ts">
import { ref } from 'vue';

defineProps<{
  label: string;
}>();
const settings = {
  check1: ref(true),
  check2: ref(false),
  check3: ref(false),

  notif1: ref(true),
  notif2: ref(true),
  notif3: ref(false),

  volume: ref(6),
  brightness: ref(3),
  mic: ref(8),
}
</script>

<style lang="scss" scoped></style>
