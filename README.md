# stock-checker

To run the go app in development mode use the `setup.sh` script to build and run the web app.

The app will be available at http://localhost:8080/

### Docker

To build and tag the docker image use the `docker.sh` script.

The docker container image has been pushed to the docker hub public repository, to access:

`docker pull teodorstandavid/stock-checker:latest`

The image has ENV variables set for testing purposes.

To run the docker image interactively use:

`docker run --rm -p 8080:8080 teodorstandavid/stock-checker`

Docker will remove the container when it's stopped.
### Kubernetes

To deploy the manifests from the base directory use:

`kubectl apply -f ./k8s/`

To access the web-app either `kubectl` or if using `minikube` the following 2 commands cam be used.

`kubectl port-forward service/stock-checker 8080:8080`

`minikube service stock-checker --url`
