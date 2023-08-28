<template>
  <div>
    <Knob :min="min" :max="max" v-model="activeValue" show-value size="100px" :thickness="0.22" color="teal"
      track-color="grey-3" class="q-ma-md" readonly />
    <div class="text-caption text-center text-bold">{{ caption }}</div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

const activeValue = ref(0);
const props = defineProps<{
  min: number;
  max: number;
  value: number;
  caption: string;
}>();

onMounted(() => {
  activeValue.value = props.value;
});
</script>
