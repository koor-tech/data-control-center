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


// async function main() {
//     // console.log("starting the call");
//     // const client = createPromiseClient(StatsService, transport);
//     // const res = await client.getClusterStats({}, {});
//     // console.log(res);
//
//     fetch("http://localhost:8080/stats.StatsService/GetClusterStats", {
//         "method": "POST",
//         "headers": {"Content-Type": "application/json"},
//         "body": JSON.stringify({})
//     })
//         .then(response => { return response.json() })
//         .then(data => { console.log(data) })
//
//     // fetch("https://demo.connectrpc.com/connectrpc.eliza.v1.ElizaService/Say", {
//     //     "method": "POST",
//     //     "headers": {"Content-Type": "application/json"},
//     //     "body": JSON.stringify({"sentence": "I feel happy."})
//     // })
//     //     .then(response => { return response.json() })
//     //     .then(data => { console.log(data) })
//
// }
// main();
// // Define an async function to fetch cluster stats
const fetchClusterStats = async () => {
    try {
        const statsClient = await $grpc.getStatsClient();
        const emptyRequest = {}; // Create an instance of the request message
        const ClusterStatusResponse = {}; // Create an instance of the request message

        const response = await statsClient.getClusterStats(emptyRequest, ClusterStatusResponse);
        console.log('Cluster stats:', response);
        return response;
    } catch (error) {
        console.error('Error fetching cluster stats:', error);
        throw error;
    }
};

// Call the async function to fetch cluster stats
fetchClusterStats();
</script>

<template>
    <!-- TODO -->
</template>
