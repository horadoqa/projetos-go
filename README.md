# GO

Linguagem de programação criada pelo google em 2006.

## Instalando
    https://golang.org/dl/
    Fazer o Download do respectivo Sistema Operacional

1️⃣LINUX com ZSH

```bash
wget https://go.dev/dl/go1.25.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.25.5.linux-amd64.tar.gz
```

2️⃣ Corrigir PATH no zsh

Edite:

```bash
sudo vi ~/.zshrc
```

Adicione no final:

export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin


Recarregue:

```bash
source ~/.zshrc
```

3️⃣ Saber a versão

```bash
go version
go version go1.25.5 linux/amd64
```

## Criando o Go Workspace
```bash
mkdir go
cd go
mkdir pkg - Onde ficam os pacotes compartilhados das aplicações (o Go é uma linguagem bastante modular)
mkdir src - Onde ficam o sorce code
mkdir bin - Onde ficam os binários (compilados)
```

## Criando o primeiro script

Criar a pasta do projeto
```bash
scr/
mkdir hello
cd hello
```

## Criar o arquivo hello.go

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## Compilando o código
```bash
go build <nome do script>
```

## Executando o código
```bash
./hello
```

## Compilando e Executando o script
```bash
go run hello.go
```
