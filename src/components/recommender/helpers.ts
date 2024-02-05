import { RecommendationLevel } from '~~/gen/ts/api/resources/ceph/v1/recommendations_pb';

export function getRecommendationLevelTextColor(level: RecommendationLevel): string {
    switch (level) {
        case RecommendationLevel.MEDIUM:
            return 'text-warn-500';
        case RecommendationLevel.HIGH:
            return 'text-warn-500';
        case RecommendationLevel.CRITICAL:
            return 'text-error-500';
        case RecommendationLevel.UNSPECIFIED:
        case RecommendationLevel.LOW:
        default:
            return 'text-base-500';
    }
}
