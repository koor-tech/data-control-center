---
title: "Base"
---

## Technologies

* Backend:
    * Golang
    * Protobuf
    * Gin or similar HTTP server that works for us
* Frontend:
    * Quasar
    * Vue3
    * GRPCweb

## Architecture

```mermaid
---
title: "Flow"
---
flowchart LR
    User --> GRPC
    User --> HTTP
    GRPC --> EnvoyProxy
    EnvoyProxy --> Backend
    HTTP --> Backend
    Backend --- Kubernetes
    Backend --- CephCluster["Ceph Cluster"]
```

* The HTTP server is used to serve the frontend app files, provide the OAuth2/OpenID login endpoints and providing the "app config".
* The GRPC server is used to provide the APIs for, e.g., retrieving/streaming cluster stats, other information about the cluster, and actions that the dashboard will offer.

## Points

* Use the Kubernetes API and Ceph Cluster as main sources for the stats and data.
    * Kubernetes API: Node health, Pods, CephClusters, Logs, etc.
    * Ceph Cluster: Ceph health state, OSD tree, OSD stats, Storage Usages, etc.
* Connecting to a Prometheus/Alertmanager as a data source in the long term, not short term.
* OAuth2/OpenID: For the start basic oauth2 login is enough.
    * Due to nature of OAuth2/OpenID, that process will be using HTTP and not GRPC, though we could potentially get away with using a simple auth proxy in front of "everything" that then passes the user info to the backend in the future.
    * Future expansion: Using oauth2 attributes/userinfo for different roles in the dashboard.
