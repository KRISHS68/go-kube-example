# Example for a deploying a simple Golang application to Kubernetes

install go

install ko https://github.com/ko-build/ko

```
export GITHUB_TOKEN=... https://github.com/settings/tokens give Permission "write:packages"
export KO_DOCKER_REPO=ghcr.io/YOUR_GITHUB_ACCOUNT

ko build .
```

Last line of output is something like:

```
ghcr.io/guettli/go-kube-example-d4328...@sha256:0091a7...
```

