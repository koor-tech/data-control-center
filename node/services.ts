import { ConnectRouter } from "@connectrpc/connect";
import {StatsService} from "~~/gen/ts/api/services/stats/stats_connect";
import {
    ClusterStatusResponse,
    DaemonCrash,
    Services,
    MonService,
    MgrService,
    MdsService,
    OsdService,
    RgwService,
    Data,
    VolumeStatus,
    PoolStatus,
    ObjectStatus,
    UsageStatus,
    PGs,
    IoStatus,
    HealthStatus,
} from "~~/gen/ts/api/services/stats/stats_pb";


export default (router: ConnectRouter) =>
    router.service(StatsService, {
        async getClusterStats(req): Promise<ClusterStatusResponse> {

            const daemonCrashes: DaemonCrash[] = [
                {
                    description: '1 daemons have recently crashed',
                },
                {
                    description: '1 mgr modules have recently crashed',
                },
            ];

            const svc: Services = {
                mon: {
                    daemonCount: 3,
                    quorum: 'a,b,c',
                    age: '13d',
                } as MonService,
                mgr: {
                    active: 'a',
                    standbys: ['b'],
                    since: '6h',
                } as MgrService,
                mds: {
                    daemonsUp: 1,
                    hotStandbyCount: 1,
                } as MdsService,
                osd: {
                    osdCount: 5,
                    osdUp: 4,
                    osdIn: 4,
                    since: '3d',
                } as OsdService,
                rgw: {
                    activeDaemon: 1,
                    hostCount: 1,
                    zoneCount: 1,
                } as RgwService,
            };


            const data: Data = {
                volumes: {
                    description: '1/1 healthy',
                } as VolumeStatus,
                pools: {
                    poolCount: 13,
                    pgs: 217,
                } as PoolStatus,
                objects: {
                    objectCount: '15.67k',
                    size: '239 MB',
                } as ObjectStatus,
                usage: {
                    used: '3.3 GiB',
                    available: '117 GiB',
                    total: '120 GiB',
                } as UsageStatus,
                pgs: {
                    activeClean: 217,
                } as PGs,
            };

            const io: IoStatus = {
                clientRead: '852 B/s',
                clientReadOps: '1 op/s',
                clientWriteOps: '0 op/s',
            };


            return {
                id: '60d54b9e-87c1-41ec-bc14-9e9925338e59',
                health: HealthStatus.HEALTH_WARN,
                daemonCrashes: daemonCrashes,
                services: svc,
                data: data,
                io: io,
            }
        },
    });
