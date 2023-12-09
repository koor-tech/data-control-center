<script setup lang="ts">
import AncienttControl from '~/components/troubleshooting/AncienttControl.vue';
import TroubleshootReport from '~/components/troubleshooting/TroubleshootReport.vue';

useHead({
    title: 'Troubleshoot',
});
definePageMeta({
    title: 'Troubleshoot',
    requiresAuth: true,
});
</script>

<template>
    <div>
        <TroubleshootReport />
        <AncienttControl />
    </div>
</template>
