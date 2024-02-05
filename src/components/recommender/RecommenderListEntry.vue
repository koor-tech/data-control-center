<script lang="ts" setup>
import { type ClusterRecommendation, RecommendationLevel } from '~~/gen/ts/api/resources/ceph/v1/recommendations_pb';
import Container from '~/components/partials/Container.vue';

defineProps<{
    recommendation: ClusterRecommendation;
}>();
</script>

<template>
    <Container>
        <h3 class="text-xl">{{ recommendation.title }}</h3>
        <p class="text-base">{{ recommendation.description }} - {{ RecommendationLevel[recommendation.level] }}</p>
    </Container>
</template>
