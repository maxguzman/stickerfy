# Stickerfy API

## Quick start

To make this application run on your local machine you need [Docker Engine](https://docs.docker.com/engine/install/) installed, once you have it, clone this repo and start the application using these commands:

```bash
make docker.dev-dependencies
make dev
```

When it's ready you can execute the following:

- Test if the API is working: `curl http://localhost:8000/products`
- Launch the API docs: [http://localhost:8000/swagger](http://localhost:8000/swagger)
- Launch Prometheus: [http://localhost:9090](http://localhost:9090)
- Launch Grafana (there is an already configured dashboard): [http://localhost:3000](http://localhost:3000) (user: admin, pass: admin)

