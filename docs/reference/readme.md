---
title: Overview
---

The Go SDK contains packages for interacting with most aspects of the tang ecosystem.  In addition to generally useful, low-level packages such as [`keypair`](https://godoc.org/github.com/tang/go/keypair) (used for creating tang-compliant public/secret key pairs), the Go SDK also contains code for the server applications and client tools written in go.

## Godoc reference

The most accurate and up-to-date reference information on the Go SDK is found within godoc.  The godoc.org service automatically updates the documentation for the Go SDK everytime github is updated.  The godoc for all of our packages can be found at (https://godoc.org/github.com/tang/go).

## Client Packages

The Go SDK contains packages for interacting with the various tang services:

- [`horizon`](https://godoc.org/github.com/tang/go/clients/horizon) provides client access to a horizon server, allowing you to load account information, stream payments, post transactions and more.
- [`tangtoml`](https://godoc.org/github.com/tang/go/clients/tangtoml) provides the ability to resolve Tang.toml files from the internet.  You can read about [Tang.toml concepts here](../../guides/concepts/tang-toml.md).
- [`federation`](https://godoc.org/github.com/tang/go/clients/federation) makes it easy to resolve a tang addresses (e.g. `scott*tang.org`) into a tang account ID suitable for use within a transaction.

