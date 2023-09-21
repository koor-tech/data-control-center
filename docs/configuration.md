---
title: "Configuration"
---

The data-control-center will by default try look for a `config.yaml` file in the working directory and `/config`.

For an example config file, checkout [`config.example.yaml`](/config.example.yaml).

Changes to the config file are not auto detected, they require a (manual) restart of the application.

## Reference

* `logLevel`: Can be `DEBUG`, `INFO`, `WARN` and `ERROR`. Defaults to `INFO`.
* `mode`: Can be `debug` or `release`. Default: `release`.

### `jwt` Section

* `secret`: Used to generate safe [JWT tokens](https://jwt.io/).

### `ceph` Section

* `api`: Ceph Dashboard API config
    * `url`: Ceph Dashboard API URL, if run inside the same cluster your Rook/Koor Ceph cluster is running use `https://rook-ceph-mgr-dashboard:8443/api` (assumes that SSL is enabled).
    * `username`: Username of Ceph Dashboard user to use.
    * `password`: Password of Ceph Dashboard user to use.
    * `insecureSSL`: If the certificate of the Ceph Dashboard API should be verified or not.

### `oauth2` Section

* `providers`: List of OAuth2/OpenID providers.
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

### `users` Section

List of user logins, an entry looks like this:

* `username`: Username of the user.
* `password`: Password of the user.
