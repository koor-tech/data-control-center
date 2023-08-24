<template>
  <tr>
    <td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0">{{ node.name }}</td>
    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ node.description }}</td>
    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ node.capacity }}</td>
    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ node.health }}</td>
    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{{ node.lastPing }}</td>
  </tr>
</template>

<script setup lang="ts">
type Node = {
  name: string;
  description: string;
  capacity: string;
  health: number;
  lastPing: string;
}

defineProps<{
  node: Node,
}>();
</script>
