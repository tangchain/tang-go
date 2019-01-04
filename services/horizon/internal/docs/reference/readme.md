---
title: Overview
---

Horizon is an API server for the Tang ecosystem.  It acts as the interface between [tang-core](https://github.com/tang/tang-core) and applications that want to access the Tang network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Tang ecosystem](https://www.tang.org/developers/guides/) for details of where Horizon fits in. You can also watch a [talk on Horizon](https://www.youtube.com/watch?v=AtJ-f6Ih4A4) by Tang.org developer Scott Fleckenstein:

[![Horizon: API webserver for the Tang network](https://img.youtube.com/vi/AtJ-f6Ih4A4/sddefault.jpg "Horizon: API webserver for the Tang network")](https://www.youtube.com/watch?v=AtJ-f6Ih4A4)

Horizon provides a RESTful API to allow client applications to interact with the Tang network. You can communicate with Horizon using cURL or just your web browser. However, if you're building a client application, you'll likely want to use a Tang SDK in the language of your client.
SDF provides a [JavaScript SDK](https://www.tang.org/developers/js-tang-sdk/learn/index.html) for clients to use to interact with Horizon.

SDF runs a instance of Horizon that is connected to the test net: [https://horizon-testnet.tang.org/](https://horizon-testnet.tang.org/) and one that is connected to the public Tang network:
[https://horizon.tang.org/](https://horizon.tang.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/tang/js-tang-sdk)
- [Java](https://github.com/tang/java-tang-sdk)
- [Go](https://github.com/tang/go)

Community maintained libraries (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Ruby](https://github.com/tang/ruby-tang-sdk)
- [Python](https://github.com/TangCN/py-tang-base)
- [C#](https://github.com/QuantozTechnology/csharp-tang-base)
