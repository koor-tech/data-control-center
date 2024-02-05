<script setup lang="ts">
import ContentWrapper from '~/components/partials/ContentWrapper.vue';

useHead({
    title: 'Recommender',
});
definePageMeta({
    title: 'Recommender',
    requiresAuth: true,
});
</script>

<template>
    <ContentWrapper>
        <RecommenderList />
    </ContentWrapper>
</template>
