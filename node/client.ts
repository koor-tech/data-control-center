import { createPromiseClient } from "@connectrpc/connect";
import {StatsService} from "~~/gen/ts/api/services/stats/stats_connect";
import { createConnectTransport } from "@connectrpc/connect-node";

const transport = createConnectTransport({
    baseUrl: "http://localhost:8080", // this works
    //baseUrl: "http://localhost:8282", // this doesn't work

});

async function main() {
    const client = createPromiseClient(StatsService, transport);
    const res = await client.getClusterStats({},{});
    console.log(res);
}
void main();
