# Koor Data Control Center

Nice dashboard and controls for managing Rook Ceph data storage.

## Usage

You can install/run the Koor data-control-center as a [binary](https://github.com/koor-tech/data-control-center/releases) or via [the Helm chart](charts/data-control-center/README.md) in Kubernetes.

Using the Helm chart to run in Kubernetes is the recommended way.

### Configuration

Checkout the [configuration documentation here](docs/configuration.md).

## Development

### Requirements

* `make` installed
* Golang version `1.21` or higher
* Node version of at least `v16.x`
* `yarn` for, e.g., `yarn dev` to preview the frontend.
* [`buf` cli](https://buf.build/docs/installation) installed (you can use `make buf` to install it to your Go bin path)
* `go-licenses`: Run `go install github.com/google/go-licenses@latest`.
* `prettier` is installed when running `yarn`.

### Tasks

* Generate Protobuf code: `make gen-proto`
* Run server: `make run-server`
* Installing the frontend dependencies: `yarn`
* Run frontend dev server: `yarn dev`

### VSCode/Codium Users

* You can use the built-in debugger to start the "compound" debug group called "All" to have everything ready to start development.

### Contributing

To run the server and everything you need to create a copy of the `config.example.yaml` called `config.yaml` in the root of the repository.
Depending on what you want to develop/contribute to, you might want to update the `config.yaml` as you need. It is recommended to add an user to the `users:` list so you can login.

For running the data-control-center a kubeconfig file is currently required. A [minikube](https://kubernetes.io/de/docs/tasks/tools/install-minikube/) cluster ([extra configuration might be needed](https://github.com/rook/rook/blob/master/Documentation/Contributing/development-environment.md#minikube)) with a [Rook Ceph test cluster (manifest)](https://github.com/rook/rook/blob/master/deploy/examples/cluster-test.yaml) instealld is enough for simple development and testing.

#### Guidelines

* Must run `prettier` format on any changed files. For Codium/VSCode users that should automatically be configured due to the settings in [`.vscode/` folder](.vscode/).

## License

The project is licensed under the [MIT License](/LICENSE).

Licenses of used libraries, code and media can be found in the [`src/public/licenses/` folder](/src/public/licenses/).
