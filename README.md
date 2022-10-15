# Example for deploying a simple Golang application to Kubernetes

Go source is in main.go. It is a simple http server which returns the http request headers.

# Requirements

Install go. For example `sudo snap install go --classic` for Ubuntu.

Install ko-build: https://github.com/ko-build/ko

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

