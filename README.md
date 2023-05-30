# Stickerfy

[![codecov](https://codecov.io/gh/maxguzman/stickerfy/branch/main/graph/badge.svg?token=SWY3J7HWJ6)](https://codecov.io/gh/maxguzman/stickerfy)

Stickerfy is a testing oriented application that helps software engineers to build and test applications using the best practices when developing and deploying apps into a cloud native environment.

## Architecture

The application architecture is based on Docker containers, there is a `Makefile` that automate the local development environment installing all the dependencies that this application need:

![stickerfy-architecture](/static/stickerfy-architecture.png)

## Quick start

```bash
make docker.run
```

When it's ready you can execute the following:

- Access to the UI: [http://localhost:8080](http://localhost:8080)
- Test if the API is working: `curl http://localhost:8000/v1/products`
- Launch the API docs: [http://localhost:8000/swagger](http://localhost:8000/swagger)
- Launch Prometheus: [http://localhost:9090](http://localhost:9090)
- Launch [Grafana dashboard](http://localhost:3000/d/Du-6hDx4k/stickerfy-store?orgId=1) (user: admin, pass: admin)
