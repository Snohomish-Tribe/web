# Snohomish Tribe

Website for the Snohomish Tribe

## Starting the app

```sh
go run cmd/web/*.go
```

## Go package used

-[Gomap](https://pkg.go.dev/github.com/cwinters8/gomap#section-readme)

## Push docker image to OCI Container Registry

Get the output for the base image tag from Terraform

```sh
IMAGE_TAG=$(terraform -chdir=terraform output container_repo_image_tag | tr -d '"')
```

Build and tag the Docker image

```sh
docker build -t "${IMAGE_TAG}"
```

TODO:

- Authenticate with Docker CLI to Oracle Cloud registry
- Build and push image
- Terraform for container instance
