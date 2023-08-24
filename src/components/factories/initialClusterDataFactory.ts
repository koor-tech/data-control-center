
import { ClusterData} from "~/components/types/types";// Import the ClusterData type

export const createInitialClusterData = (): ClusterData => ({
    id: null,
    daemon_crashes: [
        {
            description:'',
        },
    ],
    health: null,
    services: {
        mon: {
            daemon_count: null,
            quorum: null,
            age: null
        },
        mgr: {
            standbys: [],
            active: null,
            since: null
        },
        mds: {
            daemons_up: null,
            hot_standby_count: null
        },
        osd: {
            osd_count: 0,
            osd_up: null,
            osd_in: null,
            since: null
        },
        rgw: {
            active_daemon: null,
            host_count: null,
            zone_count: null
        }
    },
    data: {
        volumes: {
            description: null
        },
        pools: {
            pool_count: null,
            pgs: null
        },
        objects: {
            object_count: null,
            size: null
        },
        usage: {
            used: null,
            available: null,
            total: null
        },
        pgs: {
            active_clean: null
        }
    },
    io: {
        client_read: null,
        client_read_ops: null,
        client_write_ops: null
    }
});
