# REST API

## Requisitos

- GO

```bash
go version
go version go1.18.1 linux/amd64
```

- Docker

```bash
docker --version
Docker version 26.1.4, build 5650f9b
```

- Extensaão do GO para vscode

## Inicializando projeto

```bash
go mod init go-api
```

Criar a pasta cmd

Criar o arquivo main.go

Baixar o pacote do framework web

```bash
go get github.com/gin-gonic/gin
```


Realizar o teste pelo postman

localhost:8000/ping

## Criar o banco de Dados com o Docker Compose

```bash
docker compose up -d go_db

 ✔ Container go_db  Running 
```

```bash
docker container ls
CONTAINER ID   IMAGE         COMMAND                  CREATED         STATUS         PORTS                    NAMES
e89c8ea0cad8   postgres:12   "docker-entrypoint.s…"   2 minutes ago   Up 2 minutes   0.0.0.0:5432->5432/tcp   go_db
```

## Baixar o DBeaver para acessar ao banco

https://dbeaver.io/download/

Executar o dbeaver como Administrador no WINDOWS

 Solucionei assim: https://github.com/dbeaver/dbeaver/issues/34775

Preferências/conexão/desmarcar Usar armazenamento de confiança do Windows - Reiniciar o DBeaver e repetir o processo de conexão.

## Ao instalar realizar um impor

_ "github.com/lib/pq"

rodar o 

```bash
go mod tidy
```

## Executar 

```bash
go run main.go
```

## Criando a imagem docker a partir do Dockerfile

```bash
docker build -t rest-api-go:v1 .
```

## Visualizaar a imagem

 ```bash
docker image ls
REPOSITORY                 TAG            IMAGE ID       CREATED          SIZE
rest-api-go                latest         8e5214e37e4d   11 seconds ago   1.21GB
```
## Docker compose

Alterando o Docker compose para subir com a aplicação e o banco de dados (go-app e go-db)
Utilizando a imagem que foi criada anteriormente `rest-api-go`.

```bash
docker compose up -d
```

## Dockerhub

Fazer o login

```bash
docker login
```

Subindo a imagem para o 

```bash
docker push rfahham/rest-api-go:v1
docker push rfahham/rest-api-go #latest
```