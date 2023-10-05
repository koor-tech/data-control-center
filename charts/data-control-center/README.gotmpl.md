---
title: Data Control Center Helm Chart
---
{{ template "generatedDocsWarning" . }}

Installs [Koor Data Control Center](https://github.com/koor-tech/data-control-center) to inspect, configure, and manage your Koor Storage Cluster on Kubernetes.

## Introduction

This chart sets up [Koor Data Control Center](https://github.com/koor-tech/data-control-center) deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

* Kubernetes 1.19+
* Helm 3.x
* [Koor Storage Cluster](https://docs.koor.tech/v1.11/Getting-Started/quickstart/#create-a-ceph-cluster) or [Rook Ceph Cluster](https://rook.io/) v1.11.x+

See the [Helm support matrix](https://helm.sh/docs/topics/version_skew/) for more details.

## Installing

The Ceph Operator helm chart will install the basic components necessary to create a storage data control center for your Rook Ceph cluster in Kubernetes.

1. Add the Koor Helm repo
2. Install the Helm chart
3. [Create a Koor Storage cluster](https://docs.koor.tech/v1.11/Getting-Started/quickstart/#create-a-ceph-cluster).

The `helm install` command deploys the Koor Operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation. It is recommended that the Koor Operator be installed into the same namespace as the Rook Ceph cluster (e.g., `rook-ceph` namespace used as an example).

```console
helm repo add data-control-center https://koor-tech.github.io/data-control-center
helm repo update
helm install --create-namespace --namespace rook-ceph data-control-center data-control-center/data-control-center -f values.yaml
```

For example settings, see the next section or [values.yaml](/charts/data-control-center/values.yaml).

### Exposing via Ingress

To expose the data-control-center you need to use the `ingress:` config section.

An example for a common Ingress setup with NGINX ingress controller and that uses the cert-manager to retrieve a certificate:

```yaml
ingress:
  enabled: true
  # -- Ingress class name
  className: "nginx"
  annotations:
    cert-manager.io/cluster-issuer: your-cluster-issuer
  hosts:
    - host: datacontrolcenter.example.com
      paths:
        - path: /
          pathType: ImplementationSpecific
  tls:
    - secretName: data-control-center-tls
      hosts:
        - datacontrolcenter.example.com
```

## Configuration

The following table lists the configurable parameters of the rook-operator chart and their default values.

{{ template "chart.valuesTable" . }}

## Uninstalling the Chart

To see the currently installed Rook chart:

```console
helm ls --namespace rook-ceph
```

To uninstall/delete the `data-control-center` deployment:

```console
helm delete --namespace rook-ceph data-control-center
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## License

Copyright 2023 Koor Technologies, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
