# Dashkoor App (dashkoor-on-quasar)

Nice dashboard and controls for managing Rook Ceph data storage

## Install the dependencies

```bash
yarn
# or
npm install
```

### Start the app in development mode (hot-code reloading, error reporting, etc.)

```bash
quasar dev
```

### Lint the files

```bash
yarn lint
# or
npm run lint
```

### Format the files

```bash
yarn format
# or
npm run format
```

### Build the app for production

```bash
quasar build
```

### Customize the configuration

See [Configuring quasar.config.js](https://v2.quasar.dev/quasar-cli-vite/quasar-config-js).

## Development

### Requirements

* `make` installed
* `npm` with a Node version of at least `v16.x`
* [`buf` cli](https://buf.build/docs/installation) installed

### Tasks

* Re-Generate Protobuf generated code: `make gen-proto`
