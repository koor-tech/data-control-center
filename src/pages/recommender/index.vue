<script setup lang="ts">
import ContentCenterWrapper from '~/components/partials/ContentCenterWrapper.vue';

useHead({
    title: 'Recommender',
});
definePageMeta({
    title: 'Recommender',
    requiresAuth: true,
});
</script>

<template>
    <ContentCenterWrapper>
        <RecommenderList />
    </ContentCenterWrapper>
</template>
