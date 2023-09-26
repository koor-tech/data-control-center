<script setup lang="ts">
import { StatsDescription } from '~/composables/stats/types';

defineProps<{
    healthStats: StatsDescription;
}>();
</script>

<template>
    <li class="w-full flex justify-between gap-x-2 py-1">
        <div v-if="healthStats.title" class="flex min-w-0 gap-x-2">
            <div class="min-w-0 flex-auto">
                <p class="mt-1 truncate text-xs leading-5 text-gray-500">{{ healthStats.title }}</p>
            </div>
        </div>
        <div class="hidden sm:flex sm:flex-col sm:items-end">
            <p class="text-sm leading-6 text-gray-900">{{ healthStats.description }}</p>
        </div>
    </li>
</template>
