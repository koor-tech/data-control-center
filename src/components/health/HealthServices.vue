<script setup lang="ts">
import { HelpIcon } from 'mdi-vue3';
import { DisplayStatsData } from '~/composables/stats/types';
import HealthLiItem from './HealthLiItem.vue';

defineProps<{
    statsContainer: DisplayStatsData;
}>();
</script>

<template>
    <div class="relative overflow-hidden rounded-lg bg-white px-4 py-4 shadow">
        <dt class="grid grid-cols-3 gap-6 pb-2">
            <div class="rounded-md p-3" :class="[statsContainer.color ? statsContainer.color : 'bg-accent-500']">
                <component :is="statsContainer.icon ?? HelpIcon" class="h-6 w-6 text-white" aria-hidden="true" />
            </div>
            <p class="col-span-2 text-sm font-medium text-gray-500">{{ statsContainer.title }}</p>
        </dt>
        <dd class="flex">
            <ul role="list" class="w-full divide-y divide-gray-100">
                <template v-for="healthStats in statsContainer.description">
                    <HealthLiItem :healthStats="healthStats" />
                </template>
            </ul>
        </dd>
    </div>
</template>
