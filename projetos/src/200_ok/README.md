# 200 OK

Verifica uma lista de hosts e realiza requisição para cada host, dependendo do STATUS CODE, salva o host em um arquivo específico.

Se o Status Code for 200, salva no arquivo 200.csv
Se o Status Code for 404, salva no arquivo 404.csv

```bash
go run 200-ok.go

Programa para separar urls pelo seu Status Code.

As urls que serão testados estão no arquivo lista.csv

Escolha uma das opções abaixo:

1 - Executar Programa
2 - Lista de aquivos Status Code 200
3 - Lista de aquivos Status Code 404
4 - Sair do Programa
```