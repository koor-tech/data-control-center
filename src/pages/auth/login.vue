<script lang="ts" setup>
import { StatsService} from "~~/gen/ts/api/services/stats/stats_connect";
import { createPromiseClient } from "@connectrpc/connect";
// import { ElizaService } from "./gen/eliza_connect";
//import { createConnectTransport } from "@connectrpc/connect-node";
import { createConnectTransport  } from '@connectrpc/connect-web'



useHead({
    title: 'Login',
});
definePageMeta({
    title: 'Login',
    requiresAuth: false,
});

const { $grpc } = useNuxtApp();


const transport = createConnectTransport({
    baseUrl: "http://localhost:8282",
});

async function main() {
    console.log("starting the call");
    const client = createPromiseClient(StatsService, transport);
    const res = await client.getClusterStats({}, {});
    console.log(res);
}
main();
// // Define an async function to fetch cluster stats
// const fetchClusterStats = async () => {
//     try {
//         const statsClient = await $grpc.getStatsClient();
//         const emptyRequest = {}; // Create an instance of the request message
//         const ClusterStatusResponse = {}; // Create an instance of the request message
//
//         const response = await statsClient.getClusterStats(emptyRequest, ClusterStatusResponse);
//         console.log('Cluster stats:', response);
//         return response;
//     } catch (error) {
//         console.error('Error fetching cluster stats:', error);
//         throw error;
//     }
// };
//
// // Call the async function to fetch cluster stats
// fetchClusterStats();
</script>

<template>
    <!-- TODO -->
</template>
