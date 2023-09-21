# Koor Data Control Center

Nice dashboard and controls for managing Rook Ceph data storage.

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

### Contributing

* Must run `prettier` format on any changed files. For Codium/VSCode users that should automatically be configured due to the settings in [`.vscode/` folder](.vscode/).

## License

The project is licensed under the [MIT License](/LICENSE).

Licenses of used libraries, code and media can be found in the [`src/public/licenses/` folder](/src/public/licenses/).
