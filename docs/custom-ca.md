---
title: "Custom CA Certificate"
---

This guide assumes that you already have your custom CA certificate available in Kubernetes as a ConfigMap, Secret, or whatever classifies as a possible `volumes:` for a Pod's container.

1. Add the CA certificate to the data-control-center by using the `additionalVolumes:` and `additionalVolumeMounts:`.
    1. This is an example for a CA certificate being in a ConfigMap named `custom-ca-cert` under the `data` key called `ca.pem`.
        ```yaml
        additionalVolumes:
          - name: ca-cert
            configMap:
              name: custom-ca-cert

        additionalVolumeMounts:
          - name: ca-cert
            mountPath: /certs
            readOnly: true
        ```
    2. Example for a CA certificate being in a Secret:
       ```yaml
       additionalVolumes:
          - name: ca-cert
            secret:
              secretName: custom-ca-cert

        additionalVolumeMounts:
          - name: ca-cert
            mountPath: /certs
            readOnly: true
       ```
2. Add the CA cert to the list of (custom) certificates for the data-control-center to load in your `values.yaml` like this.
   1. Find the `config:` section and in the `certs:`  sub section add the certs path in the container/Pod like this:
        ```yaml
        config:
            # [...]
          certs:
            caCerts:
              - /certs/ca.pem
        ```
