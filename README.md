# Multicluster Usage of envtest Package

This repo shows how to setup multiple Kubernetes clusters for testing purposes. It also makes sure to cleanup those clusters on failure or early exit.

## Guide

Install `setup-envtest`.

```sh
go get sigs.k8s.io/controller-runtime/tools/setup-envtest@latest
```

Run tests with kubernetes version `1.23`.

```sh
KUBEBUILDER_ASSETS="$(setup-envtest use 1.23 -p path)" go test ./... -v
```

