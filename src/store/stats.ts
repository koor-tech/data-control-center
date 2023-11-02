import { ConnectError } from '@connectrpc/connect';
import { defineStore } from 'pinia';
import { ClusterRadar, ClusterStats, NodeInfo } from '~~/gen/ts/api/resources/stats/v1/stats_pb';
import { GetClusterResourcesResponse } from '~~/gen/ts/api/services/stats/v1/stats_pb';

export interface StatsState {}

export const useStatsStore = defineStore('stats', {
    state: () => ({}) as StatsState,
    persist: false,
    actions: {
        async getClusterNodes(): Promise<NodeInfo[]> {
            const { $grpc } = useNuxtApp();

            try {
                const resp = await $grpc.getStatsClient().getClusterNodes({});

                return resp.nodes;
            } catch (e) {
                if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                throw e;
            }
        },
        async getClusterRadar(): Promise<ClusterRadar> {
            const { $grpc } = useNuxtApp();

            try {
                const radar = await $grpc.getStatsClient().getClusterRadar({});

                return radar.radar!;
            } catch (e) {
                if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                throw e;
            }
        },
        async getClusterResources(): Promise<GetClusterResourcesResponse> {
            const { $grpc } = useNuxtApp();

            try {
                const resp = await $grpc.getStatsClient().getClusterResources({});

                return resp;
            } catch (e) {
                if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                throw e;
            }
        },
        async getClusterStats(): Promise<ClusterStats> {
            const { $grpc } = useNuxtApp();

            try {
                const stats = await $grpc.getStatsClient().getClusterStats({});

                return stats.stats!;
            } catch (e) {
                if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                throw e;
            }
        },
    },
});
