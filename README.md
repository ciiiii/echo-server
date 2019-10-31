# echo-server

[![Build Status](https://travis-ci.com/ciiiii/echo-server.svg?branch=master)](https://travis-ci.com/ciiiii/echo-server)

## grpc-server:

### Docker

```bash
docker run -p 8000:8000 killercai/grpc-echo:0.1.0
```

### Kubernetes

```bash
kubectl create secret tls grpc.example.com \
    --key your.key \
    --cert your.cert

kubectl apply -f grpc.yaml
```

---

## tcp-server:

```bash
docker run -p 8000:8000 -e version=v1 killercai/tcp-echo:0.1.0
```
