import {
    AlertBoxIcon,
    ArchiveArrowDownIcon,
    ChartPieIcon,
    ClipboardIcon,
    CubeIcon,
    FileDocumentIcon,
    GroupIcon,
    HarddiskIcon,
    HelpIcon,
    MonitorIcon,
    PoolIcon,
    SwapVerticalIcon,
    TableIcon,
} from 'mdi-vue3';
import { type DefineComponent } from 'vue';
import { ClusterHealth } from '~~/gen/ts/api/resources/stats/v1/stats_pb';

export type StatsDescription = { title?: string; description: string };

export type DisplayStatsData = {
    title: string;
    icon: DefineComponent;
    color?: string;
    description: StatsDescription[];
};

export type ClusterHealthStats = {
    id: string;
    health: ClusterHealth;
    stats: DisplayStatsData[];
};

export type ServiceInfo = {
    [key: string]: {
        title: string;
        color?: string;
        icon: DefineComponent;
    };
};

export const meta: ServiceInfo = {
    default: {
        title: 'Unknown',
        color: 'bg-gray-500',
        icon: markRaw(HelpIcon),
    },
    alerts: {
        title: 'Alerts',
        color: 'bg-warn-500',
        icon: markRaw(AlertBoxIcon),
    },
    mon: {
        title: 'MONs',
        icon: markRaw(MonitorIcon),
    },
    mgr: {
        title: 'MGRs',
        icon: markRaw(ClipboardIcon),
    },
    mds: {
        title: 'MDS',
        icon: markRaw(FileDocumentIcon),
    },
    osd: {
        title: 'OSDs',
        icon: markRaw(TableIcon),
    },
    rgw: {
        title: 'RADOS Gateway (RGW)',
        icon: markRaw(ArchiveArrowDownIcon),
    },
    volumes: {
        title: 'Volumes',
        icon: markRaw(HarddiskIcon),
    },
    pools: {
        title: 'Pools',
        icon: markRaw(PoolIcon),
    },
    objects: {
        title: 'Objects',
        icon: markRaw(CubeIcon),
    },
    pgs: {
        title: 'Placement Groups',
        icon: markRaw(GroupIcon),
    },
    iops: {
        title: 'Input/Output',
        icon: markRaw(SwapVerticalIcon),
    },
    usage: {
        title: 'Usage',
        icon: markRaw(ChartPieIcon),
    },
};
