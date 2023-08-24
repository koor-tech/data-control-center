<template>
        <div class="bg-gray-100 min-h-screen p-8">
            <h1 class="text-1xl font-semibold mb-4">Cluster Health Dashboard</h1>
                    <!-- Services -->
                    <div class="bg-white rounded p-4">
                        <h2 class="text-lg font-semibold">Cluster id:{{ data.id }}</h2>
                    </div>
            <ClusterHealthServices :data="data"></ClusterHealthServices>
        </div>
</template>
<script setup>
import ClusterHealthServices from "~/components/health/ClusterHealthServices.vue";
// import {createInitialClusterData} from "~/components/factories/initialClusterDataFactory.ts";
const initialClusterData = {
    id: null,
    daemon_crashes: [],
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
            osd_count: null,
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
};

const { $grpc } = useNuxtApp();
const data = ref(initialClusterData);


// Define an async function to fetch cluster stats
const fetchClusterStats = async () => {
    try {
        const statsClient = $grpc.getStatsClient();
        const emptyRequest = {}; // Create an instance of the request message
        const ClusterStatusResponse = {}; // Create an instance of the request message

        const response = await statsClient.getClusterStats(emptyRequest, ClusterStatusResponse);
        console.log('Cluster stats:', response);
        data.value = response;
    } catch (error) {
        console.error('Error fetching cluster stats:', error);
        throw error;
    }
};

// Fetch data when the component is mounted
onMounted(() => {
    fetchClusterStats();
});
// Call the async function to fetch cluster stats
</script>


