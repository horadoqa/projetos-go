package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const vezes = 5
const delay = 5

func main() {

	introducao()

	for {

		opcoes()
		
		comando := commando()

		switch comando {
		case 1:
			monitoramento()
			fmt.Println("")
		case 2: 
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 3:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func introducao() {
	nome := "Fahham"
	versao := 1.0
	
	fmt.Println("----------------------------------------------")
	fmt.Println(" Olá", nome)
	fmt.Println(" A versão desta app é", versao)
	fmt.Println("----------------------------------------------")
}

func opcoes() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")
	fmt.Println("")
	fmt.Println("Digite uma opção...")
	fmt.Println("----------------------------------------------")
}

func commando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi:", comando)
	return comando
}

func monitoramento() {
	fmt.Println("Monitorando...")
	fmt.Println("______________")
	fmt.Println("")

	// sites := []string{
	// 	"https://report.apps.tsuru.gcp.i.globo",
	// 	"https://www.globo.com",
	// }

	sites := leArquivo()

	for i := 0; i < vezes ; i++ {
		fmt.Println("Tentativa:", i)
		for i, site := range sites {
			fmt.Println("testando o site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)
	// fmt.Println(resp)

	if err != nil {
		fmt.Println("Ocorreu um erro na função testaSite: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!!!")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "não está respondendo, Status Code:", resp.StatusCode )
		registraLog(site, false)

	} 
}

func leArquivo() []string {

	var sites []string

	// Imprime o binário do arquivosites
	// arquivoUrls, err := os.Open("urls.csv")

	// Imprime o bloco do arquivo
	// arquivoUrls, err := ioutil.ReadFile("sites.txt")

	// Abre o arquivo para leitura
	arquivoUrls, err := os.Open("sites.txt")

	
	if err != nil {
		fmt.Println("Ocorreu um erro na função leArquivo: ", err)
	}

	leitor := bufio.NewReader(arquivoUrls)

	for {	// Delimitador da linha, quebra de linha
		
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		
		// fmt.Println(string(linha))

		sites = append(sites, linha)
		
		// Tratando erros
		
		// if err != nil {
		// 	fmt.Println("Ocorreu um erro ao ler a linha: ", err)
		// }

		if err == io.EOF {
			break
		}

	}
	
	// Fecha o arquivo para leitura
	arquivoUrls.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND , 0666);

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 - 15:04:05")+ " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}