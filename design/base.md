---
title: "Base"
---

## Technologies

* Protobuf
    * GRPCweb
* Golang
* Quasar
    * Vue

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

## Points

* Use the Kubernetes API and Ceph Cluster as main sources for the stats and data.
* Connecting to a Prometheus/Alertmanager as a data source in the long term, not short term.
* OAuth2: For the start basic authentication is enough.
