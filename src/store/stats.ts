import { ConnectError } from '@connectrpc/connect';
import { defineStore } from 'pinia';
import { ClusterRadar, ClusterStats, NodeInfo } from '~~/gen/ts/api/resources/stats/stats_pb';
import { GetClusterResourcesResponse } from '~~/gen/ts/api/services/stats/stats_pb';

export interface StatsState {}

export const useStatsStore = defineStore('stats', {
    state: () => ({}) as StatsState,
    persist: false,
    actions: {
        async getClusterNodes(): Promise<NodeInfo[]> {
            return new Promise(async (res, rej) => {
                const { $grpc } = useNuxtApp();

                try {
                    const resp = await $grpc.getStatsClient().getClusterNodes({});

                    return res(resp.nodes);
                } catch (e) {
                    if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                    return rej(e);
                }
            });
        },
        async getClusterRadar(): Promise<ClusterRadar> {
            return new Promise(async (res, rej) => {
                const { $grpc } = useNuxtApp();

                try {
                    const radar = await $grpc.getStatsClient().getClusterRadar({});
                    if (!radar.radar) {
                        return rej();
                    }

                    return res(radar.radar);
                } catch (e) {
                    if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                    return rej(e);
                }
            });
        },
        async getClusterResources(): Promise<GetClusterResourcesResponse> {
            return new Promise(async (res, rej) => {
                const { $grpc } = useNuxtApp();

                try {
                    const resp = await $grpc.getStatsClient().getClusterResources({});

                    return res(resp);
                } catch (e) {
                    if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                    return rej(e);
                }
            });
        },
        async getClusterStats(): Promise<ClusterStats> {
            return new Promise(async (res, rej) => {
                const { $grpc } = useNuxtApp();

                try {
                    const stats = await $grpc.getStatsClient().getClusterStats({});
                    if (!stats.stats) return rej();

                    return res(stats.stats);
                } catch (e) {
                    if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                    return rej(e);
                }
            });
        },
    },
});
