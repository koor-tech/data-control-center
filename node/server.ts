import { fastify } from "fastify";
import { fastifyConnectPlugin } from "@connectrpc/connect-fastify";
import { fastifyCors } from "@fastify/cors";
import routes from "./services";

async function main() {
    const server = fastify();
    await server.register(fastifyConnectPlugin, {
        routes,
    });
    await server.register(fastifyCors, (instance) => {
        return (req, callback) => {
            const corsOptions = {
                // This is NOT recommended for production as it enables reflection exploits
                origin: true
            };

            // do not include CORS headers for requests from localhost
            if (/^localhost$/m.test(req.headers.origin)) {
                corsOptions.origin = false
            }

            // callback expects two parameters: error and options
            callback(null, corsOptions)
        }
    });

    server.get("/", (_, reply) => {
        reply.type("text/plain");
        reply.send("Hello World!");
    });
    await server.listen({ host: "localhost", port: 8080 });
    console.log("server is listening at", server.addresses());
}
// You can remove the main() wrapper if you set type: module in your package.json,
// and update your tsconfig.json with target: es2017 and module: es2022.
void main();
