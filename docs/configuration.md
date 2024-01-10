---
title: "Configuration"
---

The data-control-center will by default try look for a `config.yaml` file in the working directory and `/config`.
For an example config file, checkout [`config.example.yaml`](/config.example.yaml).

Changes to the config file are not auto detected, they require a (manual) restart of the application.

## Reference

* `namespace`: The namespace the Rook Ceph cluster is running in. In most cases, this is `rook-ceph`. Defaults to `rook-ceph`.
* `logLevel`: Can be `DEBUG`, `INFO`, `WARN` and `ERROR`. Defaults to `INFO`.
* `mode`: Can be `debug` or `release`. Default: `release`.

### `jwt` Section

* `secret`: Used to generate safe [JWT tokens](https://jwt.io/).

### `ceph` Section

* `api`: Ceph dashboard API config
    * `url`: Ceph dashboard API URL, if run inside the data-control-center in same cluster your Rook/Koor Ceph cluster is running use `https://rook-ceph-mgr-dashboard:8443/api` (assumes that SSL is enabled on the dashboard, if not change `https` to `http`). The Ceph dashboard is part of the Ceph MGR component exposed on the configured port.
    * `username`: Username of Ceph dashboard user to use.
    * `password`: Password of Ceph dashboard user to use.
    * `insecureSSL`: If the certificate of the Ceph dashboard API should be verified or not.

### `oauth2` Section

* `providers`: List of OAuth2/OpenID providers. If one or more providers are given, the user login via password will be disabled.
    * `name`: "Computer"-Name of the provider. E.g., `keycloak`.
    * `label`: Human friendly Label. E.g., `Keycloak`
    * `homepage`: Homepage of the OAuth2/OpenID provider. E.g., `https://keycloak.example.com`
    * `type`: Type of the OAuth2/OpenID provider. Possible values: `generic`.
    * `redirectURL`: Redirect URL passed to the OpenID provider. Should be the data-control-center URL like this `https://chart-example.local/api/oauth2/callback/keycloak`.
    * `clientID`: OpenID client ID.
    * `clientSecret`: OpenID client secret.
    * `scopes`: List of OpenID scopes to request, minimum should be `["openid"]` scope for most OpenID providers.
    * `endpoints`: Open ID endpoints.
        * `authURL`: OpenID authentication endpoint. Example for Keycloak: 'https://keycloak.example.com/auth/realms/yourrealm/protocol/openid-connect/auth'
        * `tokenURL`: OpenID token endpoint. Example for Keycloak: 'https://keycloak.example.com/auth/realms/yourrealm/protocol/openid-connect/token'
        * `userInfoURL`: OpenID user info endpoint. Example for Keycloak: 'https://keycloak.example.com/auth/realms/yourrealm/protocol/openid-connect/userinfo'
    * `mapping`: User info fields mapping.
        * `id`: User ID to use. Example for Keycloak: `sub`.
        * `username`: Username to use. Example for Keycloak: `preferred_username`.

For more information on creating the appropriate OAuth2/OpenID clients in your identity provider, please checkout [SSO client config](./sso.md).

### `users` Section

List of user logins, an entry looks like this:

* `username`: Username of the user.
* `password`: Password of the user.

### `updateCheck` Section

Update check settings:

* `enabled`: Default: `true`.
* `interval`: Update check interval. Default: `24h`.

The update check requires internet access to `https://api.github.com`.

### Misc Settings

* `ancienttCmd`: Path to the `ancientt` command binary.

## Kubernetes Cluster Access

The data-control-center also requires access to a Kubernetes cluster to function.
The data-control-center tries to access a kubeconfig via the following "config precedence":

> * --kubeconfig flag pointing at a file
> * KUBECONFIG environment variable pointing at a file
> * In-cluster config if running in cluster
> * $HOME/.kube/config if exists.

(Copied from the [`GetConfig()` method docs of `sigs.k8s.io/controller-runtime`](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/client/config#GetConfig))
