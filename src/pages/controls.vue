<script setup lang="ts">
import ScrubbingSchedule from '~/components/cluster/ScrubbingSchedule.vue';

useHead({
    title: 'Controls',
});
definePageMeta({
    title: 'Controls',
    requiresAuth: true,
});
</script>

<template>
    <div class="p-2">
        <h2 class="text-xl q-mt-lg text-center">Control Panel</h2>
        <div class="flex">
            <ScrubbingSchedule />
        </div>
    </div>
</template>
