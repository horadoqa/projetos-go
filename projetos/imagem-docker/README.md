# Docker image

Reduzindo o tamanho das imagens docker

```bash
docker images
REPOSITORY            TAG       IMAGE ID       CREATED             SIZE
rest-api-go           v1        db5c1e24292e   About an hour ago   1.21GB
rfahham/rest-api-go   latest    db5c1e24292e   About an hour ago   1.21GB
rfahham/rest-api-go   v1        db5c1e24292e   About an hour ago   1.21GB
rest-api-go           latest    4249ba14b123   11 hours ago        1.21GB
postgres              12        94245d7f8bda   2 weeks ago         419MB
```

## Inicializando projeto

```bash
go mod init go-image-docker

```

# Instalar dependências

```bash
go mod tidy
```

## Gerando imagem padrao a partir do DOCKERFILE

    FROM golang:1.20-alpine

    WORKDIR /go/scr/app

    COPY . .

    EXPOSE 8000

    RUN go build -o main main.go

    CMD ["./main"]

```bash
docker build -t rfahham/imagem-padrao .
```


## Gerando imagem otimizada a partir do DOCKERFILE

    FROM golang:1.20-alpine as stage1

    WORKDIR /app

    COPY go.mod go.sum ./
    RUN go mod download

    COPY . .
    RUN CGO_ENABLE=0 GOOS=linux go build -o meuExecutavel

    FROM scratch

    COPY --from=stage1 /app/meuExecutavel /

    ENTRYPOINT [ "/meuExecutavel"]

```bash
docker build -t rfahham/imagem-otimizada .
```

## Listando as imagens

```bash
docker images

REPOSITORY                 TAG       IMAGE ID       CREATED          SIZE
rfahham/imagem-padrao      latest    1f6e5753af86   39 seconds ago   581MB
rfahham/imagem-otimizada   latest    796620068d10   16 minutes ago   11.4MB
```

## Executando o container

```bash
docker container run -d -p 8080:8080 rfahham/imagem-otimizada
```

## Parando a execução do container

```bash
docker container stop <CONTAINER ID>
```

## removendo a imagem

```bash
docker rmi <IMAGE ID>
```

Fonte: https://www.youtube.com/watch?v=cGYfMIKHC30