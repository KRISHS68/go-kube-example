# Example for deploying a simple Golang application to Kubernetes

Go source is in main.go. It is a simple http server which returns the http request headers.

# Requirements

Install go. For example `sudo snap install go --classic` for Ubuntu.

Install ko-build: https://github.com/ko-build/ko

You need a running kubernetes (or minikube) and `kubectl`.

In this text `k` is an alias for `kubectl`

# Config

The file `dot-envrc-example` contains environment variables which are
needed. I use [direnv](https://direnv.net/) to automatically enable them. Choose
your favorite way.

Modify the ko-image URL in deployment.yaml, if you forked the repo.

```
ko apply -f deployment.yaml
```

Above command will build a container image of the Go code. Then the image will be pushed to the container registry
which is defined in KO_DOCKER_REPO. Then it will use deployment.yaml, replace the ko:// URL to the URL of
the new image, and then create the deployment via your current `kubectl` config.

If you do this for the first time, it will fail, since the image is not public yet.

The navigate to the URL of the image (just add https:// before ghcr.io).

Then "Change package visibility" (in the box "Danger Zone") to public.

Execute above `ko apply` command again, and then the deployment will be available.

# Use the application

You can use k8slens, k9s or kubectl to have a look at your container.

```
k port-forward deployments/go-kube-example-deployment 8080:8080
```

Now you can see the running go code. It shows you the http request headers:

```
curl http://localhost:8080/headers
```

Up to now the port 8080 of our application is only easily accessible
for other containers which run in this pod. To make the port available
to other pods in the cluster we need a service.

# Service

To make the app available for other pods, we need to create a service.

```
k apply -f service.yaml
```

Start a temporary [netshoot](https://github.com/nicolaka/netshoot) container:

```
k run -it --rm --image=nicolaka/netshoot foo

foo> curl http://go-kube-example-service:8080/headers

foo> nslookup go-kube-example-service

Server:         10.96.0.10
Address:        10.96.0.10#53

Name:   go-kube-example-service.default.svc.cluster.local
Address: 10.103.211.158
```

