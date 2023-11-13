<script setup lang="ts">
import { HelpIcon } from 'mdi-vue3';
import HealthLiItem from './HealthLiItem.vue';
import type { DisplayStatsData } from '~/composables/stats/types';

defineProps<{
    statsContainer: DisplayStatsData;
}>();
</script>

<template>
    <div class="relative overflow-hidden rounded-lg bg-white px-4 py-4 shadow">
        <dt>
            <div
                class="flex justify-center rounded-md p-3 text-white"
                :class="[statsContainer.color ? statsContainer.color : 'bg-accent-500']"
            >
                <component :is="statsContainer.icon ?? HelpIcon" class="h-6 w-6" aria-hidden="true" />
                <p class="ml-4 col-span-2 text-sm font-medium">{{ statsContainer.title }}</p>
            </div>
        </dt>
        <dd class="flex">
            <ul role="list" class="w-full divide-y divide-gray-100">
                <HealthLiItem
                    v-for="healthStats in statsContainer.description"
                    :key="healthStats.title"
                    :health-stats="healthStats"
                />
            </ul>
        </dd>
    </div>
</template>
