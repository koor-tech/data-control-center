---
title: "Auth"
---

## Points

* Use the HTTP server for handling the OAuth2/OpenID process, because GRPC can't be used directly for it.
* Try to keep it as stateless as possible. Preferably the dashboard doesn't need a database to run.

## Security Aspects

* For the beginning we'll simply use short lived JWT tokens, in the future as we further explore the buf connect protocol.
