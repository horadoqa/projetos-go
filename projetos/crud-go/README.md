# crud-go

## Criado o módulo

Usando o endereço do repositório do projeto

```bash
go mod init github.com/rfahham/crud-go
```

## Biblioteca do go para gerar uma api

https://github.com/gin-gonic/gin

Para baixar a biblioteca

```bash
go get github.com/gin-gonic/gin
```
Esta biblioteca fica na máquina e não dentro do projeto

## Listar as bibliotecas instaladas

```bash
ls $(go env GOPATH)/pkg/mod/github.com/google/
```

## Criando o Hello World

    package main

    import (
        "github.com/gin-gonic/gin"
    )

    func main(){
        g := gin.Default()
        g.GET("/", func(ctx *gin.Context) {
            ctx.JSON(200, gin.H{
                "message": "Hello World !",
            })
        })
        g.Run(":3000")
    }


## Chamando a API

```bash
go run cmd/api/main.go
```

http://localhost:3000

{
    "message": "Hello World !"
}

## Conectando ao banco

```bash
go get github.com/jackc/pgx/v4
```



## Para rodar o projeto

make up

## Para rodar os testes

go test ./...

