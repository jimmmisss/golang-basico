package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibeIntroducao()
	for {
		exibeMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
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
	nome := "Wesley"
	versao := 1.1
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)
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

	sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
	}

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
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site", site, "está cpm problemas. Status code:", resp.StatusCode)
	}
}
