<script lang="ts" setup>
import ScrubbingSchedule from '~/components/controls/ScrubbingSchedule.vue';

useHead({
    title: 'OSD Controls',
});
definePageMeta({
    title: 'OSD Controls',
    requiresAuth: true,
});
</script>

<template>
    <div class="p-2">
        <h2 class="text-xl q-mt-lg text-center">OSD Controls</h2>
        <div class="flex">
            <ScrubbingSchedule />
        </div>
    </div>
</template>
