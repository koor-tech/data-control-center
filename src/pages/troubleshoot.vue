<script setup lang="ts">
import Container from '~/components/partials/Container.vue';
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
        <Container>
            <TroubleshootReport />
        </Container>

        <Container>
            <AncienttControl />
        </Container>
    </div>
</template>
