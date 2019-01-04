---
title: Horizon
---

Horizon is the server for the client facing API for the Tang ecosystem.  It acts as the interface between [tang-core](https://www.tang.org/developers/learn/tang-core) and applications that want to access the Tang network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Tang ecosystem](https://www.tang.org/developers/guides/) for more details.

You can interact directly with horizon via curl or a web browser but SDF provides a [JavaScript SDK](https://www.tang.org/developers/js-tang-sdk/learn/) for clients to use to interact with Horizon.

SDF runs a instance of Horizon that is connected to the test net [https://horizon-testnet.tang.org/](https://horizon-testnet.tang.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/tang/js-tang-sdk)
- [Java](https://github.com/tang/java-tang-sdk)
- [Go](https://github.com/tang/go)

Community maintained libraries (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Ruby](https://github.com/tang/ruby-tang-sdk)
- [Python](https://github.com/TangCN/py-tang-base)
- [C# .NET 2.0](https://github.com/QuantozTechnology/csharp-tang-base)
- [C# .NET Core 2.x](https://github.com/elucidsoft/dotnetcore-tang-sdk)
- [C++](https://bitbucket.org/bnogal/tangqore/wiki/Home)
