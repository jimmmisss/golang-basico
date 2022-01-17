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

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		command := lerComando()

		switch command {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("[Exibindo logs]")
			reportLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este programa")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	name := "Wesley"
	version := 1.1
	fmt.Println("Olá sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")
	return comandoLido
}

func iniciarMonitoramento() {

	/*sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
	}*/

	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func testaSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", nil)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
		registerLog(site, true)
	} else {
		fmt.Println("Site", site, "está cpm problemas. Status code:", response.StatusCode)
		registerLog(site, false)
	}
}

func lerSitesDoArquivo() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", nil)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - ONLINE: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func reportLogs() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
