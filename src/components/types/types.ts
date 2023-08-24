export interface ClusterData {
    id: string | null;
    daemon_crashes: DaemonCrash[];
    health: string | null;
    services: ServicesData;
    data: DataData;
    io: IoData;
}

export interface DaemonCrash {
    description: string;
}

export interface ServicesData {
    mon: MonService;
    mgr: MgrService;
    mds: MdsService;
    osd: OsdService;
    rgw: RgwService;
}

export interface MonService {
    daemon_count: number | null;
    quorum: string | null;
    age: string | null;
}

export interface MgrService {
    standbys: string[];
    active: string | null;
    since: string | null;
}

export interface MdsService {
    daemons_up: number | null;
    hot_standby_count: number | null;
}

export interface OsdService {
    osd_count: number | null;
    osd_up: number | null;
    osd_in: number | null;
    since: string | null;
}

export interface RgwService {
    active_daemon: number | null;
    host_count: number | null;
    zone_count: number | null;
}

export interface DataData {
    volumes: VolumesData;
    pools: PoolsData;
    objects: ObjectsData;
    usage: UsageData;
    pgs: PgsData;
}

export interface VolumesData {
    description: string | null;
}

export interface PoolsData {
    pool_count: number | null;
    pgs: number | null;
}

export interface ObjectsData {
    object_count: string | null;
    size: string | null;
}

export interface UsageData {
    used: string | null;
    available: string | null;
    total: string | null;
}

export interface PgsData {
    active_clean: number | null;
}

export interface IoData {
    client_read: string | null;
    client_read_ops: string | null;
    client_write_ops: string | null;
}
