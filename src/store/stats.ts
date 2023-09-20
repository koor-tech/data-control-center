import { defineStore } from 'pinia';
import { ClusterRadar, ClusterStats, NodeInfo } from '~~/gen/ts/api/resources/stats/stats_pb';
import { GetClusterResourcesResponse } from '~~/gen/ts/api/services/stats/stats_pb';

export interface StatsState {
    stats?: ClusterStats;
    clusterResources?: GetClusterResourcesResponse;
    clusterNodes?: NodeInfo;
    clusterRadar?: ClusterRadar;
}

export const useStatsStore = defineStore('stats', {
    state: () => ({}) as StatsState,
    persist: false,
    actions: {
        // TODO
    },
});
