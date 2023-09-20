import {
    AlertBoxIcon,
    ArchiveArrowDownIcon,
    ClipboardIcon,
    CubeIcon,
    FileDocumentIcon,
    GroupIcon,
    HarddiskIcon,
    MonitorIcon,
    PoolIcon,
    SwapVerticalIcon,
    TableIcon,
} from 'mdi-vue3';
import { DefineComponent } from 'vue';

export type StatsDescription = { title?: string; description: string };

export type DisplayStatsData = {
    title: string;
    icon?: DefineComponent;
    color?: string;
    description: StatsDescription[];
};

export type ClusterHealthStats = {
    id: string;
    health: string;
    stats: DisplayStatsData[];
};

export type ServiceInfo = {
    [key: string]: {
        title?: string;
        color?: string;
        icon?: DefineComponent;
    };
};

export const meta: ServiceInfo = {
    alerts: {
        title: 'Alerts',
        color: 'orange',
        icon: AlertBoxIcon,
    },
    mon: {
        title: 'Monitors',
        color: 'blue',
        icon: MonitorIcon,
    },
    mgr: {
        title: 'Managers',
        color: 'blue',
        icon: ClipboardIcon,
    },
    mds: {
        title: 'Metadata',
        color: 'blue',
        icon: FileDocumentIcon,
    },
    osd: {
        title: 'Object Storage',
        color: 'blue',
        icon: TableIcon,
    },
    rgw: {
        title: 'RADOS Gateway (RGW)',
        color: 'blue',
        icon: ArchiveArrowDownIcon,
    },
    volumes: {
        title: 'Volumes',
        color: 'blue',
        icon: HarddiskIcon,
    },
    pools: {
        title: 'Pools',
        color: 'blue',
        icon: PoolIcon,
    },
    objects: {
        title: 'Objects',
        color: 'blue',
        icon: CubeIcon,
    },
    pgs: {
        title: 'Placement Groups',
        icon: GroupIcon,
    },
    io: {
        title: 'Input/Output',
        icon: SwapVerticalIcon,
    },
};

export const statuses: { [key: string]: string } = {
    HEALTH_OFFLINE: 'text-gray-500 bg-gray-100/10',
    HEALTH_OK: 'text-green-400 bg-green-400/10',
    HEALTH_WARN: 'text-orange-400 bg-orange-400/10',
};
