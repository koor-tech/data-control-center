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

### Tasks

* Installing the frontend dependencies: `yarn`
* Run frontend dev server: `yarn dev`
* Generate Protobuf code: `make gen-proto`

## License

The project is licensed under the [MIT License](/LICENSE).

Licenses of used libraries, code and media can be found in the [`src/public/licenses/` folder](/src/public/licenses/).
