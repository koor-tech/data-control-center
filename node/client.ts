import { createPromiseClient } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-node';
import { StatsService } from '~~/gen/ts/api/services/stats/stats_connect';

const transport = createConnectTransport({
    baseUrl: 'http://localhost:8282', // this works
    //baseUrl: "http://localhost:8282", // this doesn't work
});

async function main() {
    const client = createPromiseClient(StatsService, transport);
    const res = await client.getClusterStats({}, {});
    console.log(res);
}
void main();
