# Snohomish Tribe

Website for the Snohomish Tribe

## Starting the app

```sh
go run cmd/web/*.go
```

## Go package used

-[Gomap](https://pkg.go.dev/github.com/cwinters8/gomap#section-readme)

## Push docker image to OCI Container Registry

### Prerequisite

[Authenticate with Docker CLI to Oracle Cloud image registry](https://docs.oracle.com/en-us/iaas/Content/Registry/Tasks/registrypushingimagesusingthedockercli.htm#Pushing_Images_Using_the_Docker_CLI)

Once you have an auth token, login with Docker:

```sh
docker login ocir.us-sanjose-1.oci.oraclecloud.com
```

Username will be in the format: axihvv9biq8w/your-username

If successful, you'll get the message: `Login Succeeded`

### Build

Get the output for the base image tag from Terraform

```sh
IMAGE_TAG=$(terraform -chdir=terraform output container_repo_image_tag | tr -d '"')
```

Build and tag the Docker image

```sh
TAG="${IMAGE_TAG}:latest"
docker build -t "${TAG}" .
```

### Push

```sh
docker push "${TAG}"
```

TODO:

- Pass environment variables to container definition
