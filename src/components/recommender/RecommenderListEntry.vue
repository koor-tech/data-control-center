<script lang="ts" setup>
import { InformationIcon } from 'mdi-vue3';
import { type ClusterRecommendation, RecommendationLevel } from '~~/gen/ts/api/resources/ceph/v1/recommendations_pb';
import GenericContainer from '~/components/partials/GenericContainer.vue';
import { getRecommendationLevelTextColor } from '~/components/recommender/helpers';
import GenericBadge from '~/components/partials/GenericBadge.vue';

defineProps<{
    recommendation: ClusterRecommendation;
}>();
</script>

<template>
    <GenericContainer>
        <h3 class="text-xl inline-flex items-center gap-1">
            <GenericBadge class="gap-1">
                <InformationIcon class="h-5 w-auto" :class="getRecommendationLevelTextColor(recommendation.level)" />
                {{ RecommendationLevel[recommendation.level] }}
            </GenericBadge>
            {{ recommendation.title }}
        </h3>
        <p class="text-base">{{ recommendation.description }}</p>
        <template v-if="recommendation.extraData && recommendation.extraData.case !== undefined">
            <hr class="my-1" />
            <p v-if="recommendation.extraData.case === 'recommendedValue'" class="text-sm">
                Current: {{ recommendation.extraData.value.current }} - Expected: {{ recommendation.extraData.value.expected }}
            </p>
        </template>
    </GenericContainer>
</template>
