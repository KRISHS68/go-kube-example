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

The navigate to the URL of the image (just add https:// before ghcr.io).
Then "Change package visibility" (in the box "Danger Zone") to public.

```
kubectl create deployment go-kube-example --image=ghcr.io/guettli/go-kube-example-d432...
```

TODO: expose port, first in the cluster, then via ingress.


