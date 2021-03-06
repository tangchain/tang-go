# Horizon
[![Build Status](https://travis-ci.org/tang/horizon.svg?branch=master)](https://travis-ci.org/tang/horizon)

Horizon is the [client facing API](/docs) server for the Tang ecosystem.  It acts as the interface between tang-core and applications that want to access the Tang network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Tang ecosystem](https://www.tang.org/developers/guides/get-started/) for more details.

## Downloading the server
[Prebuilt binaries](https://github.com/tang/go/releases) of horizon are available on the 
[releases page](https://github.com/tang/go/releases).

See [the old releases page](https://github.com/tang/horizon/releases) for prior releases

| Platform       | Binary file name                                                                         |
|----------------|------------------------------------------------------------------------------------------|
| Mac OSX 64 bit | [horizon-darwin-amd64](https://github.com/tang/go/releases/download/horizon-v0.12.0-testing/horizon-v0.12.0-testing-darwin-amd64.tar.gz)      |
| Linux 64 bit   | [horizon-linux-amd64](https://github.com/tang/go/releases/download/horizon-v0.12.0-testing/horizon-v0.12.0-testing-linux-amd64.tar.gz)       |
| Windows 64 bit | [horizon-windows-amd64.exe](https://github.com/tang/go/releases/download/horizon-v0.12.0-testing/horizon-v0.12.0-testing-windows-amd64.zip) |

Alternatively, you can [build](#building) the binary yourself.

## Dependencies

Horizon requires go 1.9 or higher to build. See (https://golang.org/doc/install) for installation instructions.

## Building

[dep](https://golang.github.io/dep/) is used for building horizon.

Please, follow the [dep installation guide](https://golang.github.io/dep/docs/installation.html) to get the `dep` on your
system. 

Next, you must download the source for packages that horizon depends upon. From within the project directory, run:

```bash
dep ensure -v
```

Then, simply run `go install github.com/tang/go/services/horizon`.  After successful
completion, you should find `horizon` is present in your `$GOPATH/bin` directory.

More detailed intructions and [admin guide](internal/docs/reference/admin.md). 

## Developing Horizon

See [the development guide](internal/docs/developing.md).

## Contributing
Please see the [CONTRIBUTING.md](./CONTRIBUTING.md) for details on how to contribute to this project.
