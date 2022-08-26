# BEMS API

Swagger + Gin: https://github.com/swaggo/gin-swagger

## Setup

```bash
helm install target-exporter charts/target-exporter --namespace=target-exporter --create-namespace
```

to uninstall:

```bash
helm uninstall -n target-exporter target-exporter
```

## Build

```bash
docker build -t bems-api .
docker login
docker tag bems-api cristianohelio/bems-api:0.0.1
docker push cristianohelio/bems-api:0.0.1
```

Note: repo was moved to a private GitHub repo (https://github.com/criscola/bems-api)
