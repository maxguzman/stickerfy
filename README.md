# Stickerfy

[![codecov](https://codecov.io/gh/maxguzman/stickerfy/branch/main/graph/badge.svg?token=SWY3J7HWJ6)](https://codecov.io/gh/maxguzman/stickerfy)

Stickerfy is a testing oriented application that helps software engineers to build and test applications using the best practices when developing and deploying apps into a cloud native environment.

## Architecture

The architecture of this application is based on Docker containers, there is a `Makefile` that automate the local development environment installing all the dependencies that this application need:

![stickerfy-architecture](/static/stickerfy-architecture.png)

## Quick start

To make this application run on your local machine you need [Docker Engine](https://docs.docker.com/engine/install/) installed, once you have it, clone this repo and start the application using these commands:

```bash
make docker.dev-dependencies
make docker.stickerfy
```

Or just run the following command that make both

```bash
make docker.run
```

When it's ready you can execute the following:

- Test if the API is working: `curl http://localhost:8000/products`
- Launch the API docs: [http://localhost:8000/swagger](http://localhost:8000/swagger)
- Launch Prometheus: [http://localhost:9090](http://localhost:9090)
- Launch Grafana (there is an already configured dashboard): [http://localhost:3000](http://localhost:3000) (user: admin, pass: admin)

## Set up the development environment

For coding this app you don't want to deploy it to Docker, instead follow these steps:

- Setup the dev dependencies

```bash
make docker.dev-dependencies
```

- Set the local environment variables and run the app:

```bash
make dev
```
