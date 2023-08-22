<template>
  <q-item clickable tag="a" target="_blank" :href="link">
    <q-item-section v-if="icon" avatar>
      <q-icon :name="icon" />
    </q-item-section>

    <q-item-section>
      <li>{{ title }}</li>
      <li caption>{{ caption }}</li>
    </q-item-section>
  </q-item>
</template>

<script setup lang="ts">
withDefaults(defineProps<{
  title: string;
  caption: string;
  link: string;
  icon: string;
}>(), {
  caption: '',
  link: '',
  icon: '',
});
</script>
