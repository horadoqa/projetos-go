# Trabalhando com sqlite3

## Youtubee

https://www.youtube.com/watch?v=4w1pXqJMoA0


## Instalando o SQLite3

```bash
sudo apt-get install sqlite
```

## Criando o banco

```bash
✗ sqlite3 users.db
SQLite version 3.37.2 2022-01-06 13:25:41
Enter ".help" for usage hints.
sqlite> 
```

## Criando a Tabela

```bash
sqlite> create table users (id interger, name text, email text);
```

## Verificando a tabela

```bash
sqlite> .tables
users
```

## Inserindo dados

```bash
sqlite> insert into users values(1,'Ricardo Fahham','rfahham@hotmail.com');
```

## Verificando dados

```bash
sqlite> select * from users;
```

## Apagando todos os dados

```bash
sqlite> delete from users;
```

## Versinando o projeto

```bash
go mod init github.com/projetos-go/src/db
```

## Adicionando dependências do projeto

```bash
go mod tidy
```

## ERROR

Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub

CGO_ENABLED=1 go install github.com/mattn/go-sqlite3
go build

## Criando usuário

