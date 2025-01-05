# Development
Start the application:
```console
go run cmd/main.go
```

Send email with 


# Build docker image
```console
docker build -t kinodvor:0.1 .
```

```console
docker run -d --name kinodvor_container kinodvor:0.1
```

```console
docker tag kinodvor:0.1 europe-west3-docker.pkg.dev/kinodvor/kinodvor/kinodvor:0.1
```

```console
docker push europe-west3-docker.pkg.dev/kinodvor/kinodvor/kinodvor:0.1
```
