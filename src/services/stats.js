import {createGrpcWebTransport } from "@connectrpc/connect-web";
import { createPromiseClient } from "@bufbuild/connect";
import { StatsService } from 'app/gen/ts/api/services/stats/stats_connect';


export default class {
    constructor(deps) {
        this.proto = deps.proto
    }

    async getStats () {

        const transport = createGrpcWebTransport({
            baseUrl: 'http://localhost:9000',
        });

        const client = createPromiseClient(StatsService, transport);

        const res = await client.getClusterStats(undefined, undefined);
        console.log(res);
        // client.getClusterStats({},  (err, res) => {
        //     if (!err) {
        //         console.log(res);
        //     }
        // });
    }
}
