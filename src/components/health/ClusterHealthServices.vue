<script setup lang="ts">
import type { DisplayStatsData } from '~/composables/stats/types';

defineProps<{
    stats: DisplayStatsData[];
}>();
</script>

<template>
    <div class="w-full relative isolate">
        <!-- Stats -->
        <dl class="mt-5 grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-5 px-2">
            <template v-for="item in stats">
                <HealthServices v-if="item.description.length" :key="item.title" :stats-container="item" />
            </template>
        </dl>
    </div>
</template>
