<template>
    <div class="w-full h-full flex">
        <div class="w-full mx-4 my-4 rounded-lg border border-gray-300 py-4">
            <header class="inset-x-0 top-0 z-50 flex h-16 border-b border-gray-900/10">
                <div class="mx-auto flex w-full max-w-7xl items-center justify-between">
                    <div class="flex items-center gap-x-6">
                        Cluster Health stats
                    </div>
                </div>
            </header>
            <!-- Your content goes here -->
            <div class="w-full  relative isolate">
                <!-- Secondary navigation -->
                <header class="w-full border-b border-b-gray-300">
                    <div class="mx-auto flex max-w-7xl flex-wrap items-center gap-6 sm:flex-nowrap">
                        <h1 class="text-base font-semibold leading-7 text-gray-900">Cluster id: {{ clusterStats?.id }}</h1>
                        <div :class="[statuses['warn'], 'flex-none rounded-full p-1']">
                            <div class="h-2 w-2 rounded-full bg-current" />
                        </div>
                        <p class="text-gray-700 text-sm">
                            HEALTH WARN
                        </p>
                        <div
                            class="order-last flex w-full gap-x-8 text-sm font-semibold leading-6 sm:order-none sm:w-auto sm:border-l sm:border-gray-200 sm:pl-6 sm:leading-7">
                            <a v-for="item in secondaryNavigation" :key="item.name" :href="item.href"
                                :class="item.current ? 'text-indigo-600' : 'text-gray-700'">{{ item.name }}</a>
                        </div>
                    </div>
                </header>

                <!-- Stats -->
                <dl class="mt-5 grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-5 px-2">
                    <HealthServices v-if="clusterStats" v-for="item in clusterStats.stats" :statsContainer="item" />

                    <!-- <Doughnut :data="data" :options="options" /> -->
                </dl>

            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
const { $grpc } = useNuxtApp();

const { data: clusterStats } = useLazyAsyncData('clusterStats', async () => transformData(await $grpc.getStatsClient().getClusterStats({})));

type TransformedData = {
    title: string;
    icon?: string,
    color?: string,
    description: { title: string; description?: string }[];
};

type ClusterStats = {
    id: string;
    stats: TransformedData[];
}

type ServiceInfo = {
    [key: string]: {
        title?: string;
        color?: string,
        icon?: string;
    };
};

const meta: ServiceInfo = {
    Alerts: {
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

function transformDescriptions(title: string, data: any): TransformedData[] {
    return Object.keys(data).map(serviceName => {
        const serviceData = data[serviceName];
        const description = Object.entries(serviceData)
            .map(([key, value]) => `${key}: ${value}`)
            .join(', ');

        console.log(meta[serviceName]?.icon);

        return {
            title: meta[serviceName]?.title || serviceName,
            icon: meta[serviceName]?.icon,
            color: meta[serviceName]?.color,
            description: description.split(', ').map(entry => {
                const [key, value] = entry.split(': ');
                return { title: key, description: value };
            }),
        };
    });
}

function transformData(clusterStats: any): ClusterStats {
    const transformedArray: TransformedData[] = [];

    transformedArray.push({
        title: 'Alerts',
        icon: meta['Alerts'].icon,
        color: meta['Alerts']?.color,

        description: clusterStats.daemonCrashes.map((daemon: any) => daemon.description)
    });

    transformedArray.push(...transformDescriptions('services', clusterStats.services));
    transformedArray.push(...transformDescriptions('data', clusterStats.data));


    transformedArray.push({
        title: 'Input / Output',
        icon: meta['io'].icon,
        color: meta['io']?.color,
        description: Object.keys(clusterStats.io).map(serviceName => ({
            title: serviceName, description: String(clusterStats.io[serviceName])
        })),
    });

    console.log(clusterStats);

    return {
        id: clusterStats.id, stats: transformedArray
    };
}

const secondaryNavigation = [
    { name: 'Cluster 2', href: '#', current: true },
    { name: 'Cluster 3', href: '#', current: false },
    { name: 'Cluster 4', href: '#', current: false },
]


const statuses = {
    offline: 'text-gray-500 bg-gray-100/10',
    health: 'text-green-400 bg-green-400/10',
    warn: 'text-orange-400 bg-orange-400/10',
}

</script>
