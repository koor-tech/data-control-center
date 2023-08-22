<template>
  <q-item clickable tag="a" :to="{ name: to }" exact>
    <q-item-section v-if="icon" avatar>
      <q-icon :name="icon" />
    </q-item-section>

    <q-item-section>
      <li>{{ title }}</li>
      <li>{{ caption }}</li>
    </q-item-section>
  </q-item>
</template>

<script setup lang="ts">
withDefaults(defineProps<{
  title: string;
  caption: string;
  to: string;
  icon: string;
  exact?: boolean;
}>(), {
  caption: '',
  to: '',
  icon: '',
  default: false,
});
</script>
