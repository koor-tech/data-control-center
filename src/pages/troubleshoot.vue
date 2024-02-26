<script setup lang="ts">
import GenericContainer from '~/components/partials/GenericContainer.vue';
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
    <div class="p-2">
        <GenericContainer>
            <TroubleshootReport />
        </GenericContainer>

        <GenericContainer>
            <AncienttControl />
        </GenericContainer>
    </div>
</template>
