import { ConnectError } from '@connectrpc/connect';
import { defineStore, type StoreDefinition } from 'pinia';
import { KoorCluster } from '~~/gen/ts/api/resources/koor/v1/koor_pb';
import { GetTroubleshootReportResponse } from '~~/gen/ts/api/services/cluster/v1/cluster_pb';

export interface ClusterState {}

export const useClusterStore = defineStore('cluster', {
    state: () => ({}) as ClusterState,
    persist: false,
    actions: {
        async getKoorCluster(): Promise<KoorCluster> {
            const { $grpc } = useNuxtApp();

            try {
                const resp = await $grpc.getClusterClient().getKoorCluster({});

                return resp.koorCluster!;
            } catch (e) {
                if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                throw e;
            }
        },
        async getTroubleshootReport(): Promise<GetTroubleshootReportResponse> {
            const { $grpc } = useNuxtApp();

            try {
                return await $grpc.getClusterClient().getTroubleshootReport({});
            } catch (e) {
                if (e instanceof ConnectError) $grpc.handleError(e as ConnectError);
                throw e;
            }
        },
    },
});

if (import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useClusterStore as unknown as StoreDefinition, import.meta.hot));
}
