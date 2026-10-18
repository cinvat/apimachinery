# apimachinery

The shared API contract for the cinvat kloud: the envelope every resource API
serves, the metadata common to every object, and the discovery documents that
let a client turn a manifest's `kind` into the collection path it lives at.

`apimachinery` is to cinvat what `k8s.io/apimachinery` is to Kubernetes: a
neutral module that the aggregation origin and every downstream `apiService`
import. It stores no state and makes no decisions; it only names the shapes on
the wire.

## Why a separate module

The contract is spoken by many services:

```
                  github.com/cinvat/apimachinery
                       (the shared contract)
                              ^     ^     ^
                              |     |     |
                          kloudapi arsan alvand ... iam
                       (aggregation   (downstream apiServices)
                          origin)
```

It is deliberately **not** owned by any one of them. Putting it in kloudapi
would make every downstream depend on the aggregator that fronts it — the
thing that routes to you becoming your build dependency. A neutral module keeps
the graph a star: everyone imports it, nobody depends on anybody. It is also
the single fix for a contract that once drifted, with each service redefining
`ObjectMeta` and friends slightly differently.

## What is shared — and what is not

Shared, and only this:

- `TypeMeta` — `apiVersion` + `kind`, the object's type on the wire.
- `ObjectMeta` — the client-visible address (`name`, `namespace`).
- `APIResource`, `APIResourceList` — the groupVersion discovery document an
  `apiService` serves at its mount path joined with its version.

**Domain types never live here.** `Node`, `VirtualMachine`, `Flavor`, `Volume`
(arsan) and `Network`, `Subnet`, `Port`, `SecurityGroup` (alvand) belong to
the service that owns the records and their judgement.

## Use

```go
import corev1 "github.com/cinvat/apimachinery/api/core/v1"

type ObjectMeta = corev1.ObjectMeta
```

Envelope objects are a uniform `apiVersion`, `kind`, `metadata`, `spec`,
`status`; lists are JSON arrays; discovery is an `APIResourceList`. The full
routing, discovery and watch contract a downstream serves is documented in
kloudapi's `docs/apiService-development.md`, and the reasoning behind this
module is in its `docs/business-logics.md` §7.

## Layout

```
api/core/v1/types.go   the shared contract types
```

## Consuming

Consumers pin a released version:

```
require github.com/cinvat/apimachinery v0.1.0
```

While iterating on the contract, a consumer may temporarily point at a local
checkout with a `replace`:

```
replace github.com/cinvat/apimachinery => ../apimachinery
```

(kloudctl, one directory deeper inside the kloudapi repo, uses
`../../apimachinery`.) That is a transient aid, not the committed state.

## Build

```sh
go build ./... && go vet ./... && go test ./...
```
