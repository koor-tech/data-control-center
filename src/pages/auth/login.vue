<script lang="ts" setup>
useHead({
    title: 'Login',
});
definePageMeta({
    title: 'Login',
    requiresAuth: false,
});

const { $grpc } = useNuxtApp();
$grpc.getStatsClient();
</script>

<template>
    <!-- TODO -->
</template>