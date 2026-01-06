# Generate UUID google

## Repositório

https://github.com/google/uuid

## Inicializar o modulo uuid

```bash
go mod init github.com/codeedu/myuuid
```

## Install

Instala pacotes e binários

```bash
go get github.com/google/uuid@latest
```
a partir da versão 1.16 usar o `go install`

```bash
go install github.com/google/uuid@latest
```

## Executando

```bash
go run main.go
    917137d5-74a7-4edc-92c6-9cdbbae41db9
```

## Verificando os módulos

```bash
go mod tidy
```

## Verificando as dependências

```bash
go mod graph
    github.com/codeedu/myuuid github.com/google/uuid@v1.6.0
    github.com/codeedu/myuuid go@1.22.2
    go@1.22.2 toolchain@go1.22.2
```

## Vendor

Deixa as dependências/módulos isolados no projeto

```bash
go mod vendor
```
