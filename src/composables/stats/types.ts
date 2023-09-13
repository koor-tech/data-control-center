export type DisplayStatsData = {
    title: string;
    icon?: string,
    color?: string,
    description: { title: string; description: string }[];
};

export type ClusterHealthStats = {
    id: string;
    health: string;
    stats: DisplayStatsData[];
}

export type ServiceInfo = {
    [key: string]: {
        title?: string;
        color?: string,
        icon?: string;
    };
};

export const meta: ServiceInfo = {
    alerts: {
        title: 'Alerts',
        color: 'orange',
        icon: 'ShieldExclamationIcon'
    },
    mon: {
        title: 'Monitors',
        color: 'blue',
        icon: 'ComputerDesktopIcon'
    },
    mgr: {
        title: 'Managers',
        color: 'blue',
        icon: 'ClipboardIcon'
    },
    mds: {
        title: 'Metadata',
        color: 'blue',
        icon: 'DocumentTextIcon'
    },
    osd: {
        title: 'Object Storage',
        color: 'blue',
        icon: 'TableCellsIcon'
    },
    rgw: {
        title: 'RADOS Gateway (RGW)',
        color: 'blue',
        icon: 'ArchiveBoxArrowDownIcon'
    },
    volumes: {
        title: 'Volumes',
        color: 'blue',
        icon: 'CircleStackIcon'
    },
    pools: {
        title: 'Pools',
        color: 'blue',
        icon: 'Square3Stack3DIcon'
    },
    objects: {
        title: 'Objects',
        color: 'blue',
        icon: 'CubeIcon'
    },
    pgs: {
        title: 'Placement Groups',
        icon: 'ArchiveBoxIcon'
    },
    io: {
        title: 'Input/Output',
        icon: 'ArrowsUpDownIcon'
    },
}

export const statuses: { [key: string]: string } = {
    HEALTH_OFFLINE: 'text-gray-500 bg-gray-100/10',
    HEALTH_OK: 'text-green-400 bg-green-400/10',
    HEALTH_WARN: 'text-orange-400 bg-orange-400/10',
};
