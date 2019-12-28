# Azure Red Hat OpenShift Resource Provider

## Notice

For information relating to the generally available Azure Red Hat OpenShift v3
service, please see the following links:

* https://azure.microsoft.com/en-us/services/openshift/
* https://www.openshift.com/products/azure-openshift
* https://docs.microsoft.com/en-us/azure/openshift/
* https://docs.openshift.com/aro/welcome/index.html


## Quickstarts

* If you have a whitelisted subscription and want to use `az aro` to create a
  cluster using the production RP, follow [using `az
  aro`](docs/using-az-aro.md).

* If you want to deploy a development RP, follow [deploy development
  RP](docs/deploy-development-rp.md).


## Contributing

This project welcomes contributions and suggestions. Most contributions require
you to agree to a Contributor License Agreement (CLA) declaring that you have
the right to, and actually do, grant us the rights to use your contribution. For
details, visit https://cla.microsoft.com.

When you submit a pull request, a CLA-bot will automatically determine whether
you need to provide a CLA and decorate the PR appropriately (e.g., label,
comment). Simply follow the instructions provided by the bot. You will only need
to do this once across all repositories using our CLA.

This project has adopted the [Microsoft Open Source Code of
Conduct](https://opensource.microsoft.com/codeofconduct/). For more information
see the [Code of Conduct
FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or contact
[opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional
questions or comments.


## Repository map

* .github/workflows: CI workflows using GitHub Actions.

* cmd/aro: RP entrypoint.

* deploy: ARM templates to deploy RP in development and production.

* docs: Documentation.

* hack: Build scripts and utilities.

* pkg: RP source code:

  * pkg/api: RP internal and external API definitions.

  * pkg/backend: RP backend workers.

  * pkg/client: Autogenerated ARO service Go client.

  * pkg/database: RP CosmosDB wrapper layer.

  * pkg/deploy: /deploy ARM template generation code.

  * pkg/env: RP environment-specific shims for running in production,
    development or test

  * pkg/frontend: RP frontend webserver.

  * pkg/install: OpenShift installer wrapper layer.

  * pkg/mirror: OpenShift release mirror tooling.

  * pkg/swagger: /swagger Swagger specification generation code.

  * pkg/util: Utility libraries.

* python: Autogenerated ARO service Python client and `az aro` client extension.

* swagger: Autogenerated ARO service Swagger specification.

* test: End-to-end tests.

* vendor: Vendored Go libraries.


## Basic architecture

* pkg/frontend is intended to become a spec-compliant RP web server.  It is
  backed by CosmosDB.  Incoming PUT/DELETE requests are written to the database
  with an non-terminal (Updating/Deleting) provisioningState.

* pkg/backend reads documents with non-terminal provisioningStates,
  asynchronously updates them and finally updates document with a terminal
  provisioningState (Succeeded/Failed).  The backend updates the document with a
  heartbeat - if this fails, the document will be picked up by a different
  worker.

* As CosmosDB does not support document patch, care is taken to correctly pass
  through any fields in the internal model which the reader is unaware of (see
  `github.com/ugorji/go/codec.MissingFielder`).  This is intended to help in
  upgrade cases and (in the future) with multiple microservices reading from the
  database in parallel.

* Care is taken to correctly use optimistic concurrency to avoid document
  corruption through concurrent writes (see `RetryOnPreconditionFailed`).

* The pkg/api architecture differs somewhat from
  `github.com/openshift/openshift-azure`: the intention is to fix the broken
  merge semantics and try pushing validation into the versioned APIs to improve
  error reporting.

* Everything is intended to be crash/restart/upgrade-safe, horizontally
  scaleable, upgradeable...


## Useful links

* https://github.com/Azure/azure-resource-manager-rpc

* https://github.com/microsoft/api-guidelines

* https://docs.microsoft.com/en-gb/rest/api/cosmos-db

* https://github.com/jim-minter/go-cosmosdb
