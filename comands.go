package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const MONITORAMENTOS = 3
const DELAY = 1

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	codes := []int{200, 300, 400, 500}
	randomIndex := rand.Intn(len(codes))
	
	sites := capturaSites()
	if sites[0] != "" {
		for i := 0; i < MONITORAMENTOS; i++ {
			for _, site := range sites {
				if site == "https://httpbin.org/status/" {
					endpoint := fmt.Sprint(site, codes[randomIndex])
					validaSite(endpoint)
				} else {
					validaSite(site)
				}
			}
			time.Sleep(DELAY * time.Second)
			cmd := exec.Command("echo", "Iniciando novo monitoramento....")
			stdout, err := cmd.Output()
			
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(string(stdout))
		}
	} else {
		fmt.Println("Arquivo sites.txt vazio")
		fmt.Println("Voltando ao menu...")
	}

	main()
}

func validaSite(endpoint string) {
	resp, err := http.Get(endpoint)

	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", endpoint, "foi carregado com suecesso")
		registraLog(endpoint, true)
	} else {
		fmt.Println("O site", endpoint, "está com problemas")
		registraLog(endpoint, false)
	}
}

func exibirLogs() {
	arquivo, err := os.ReadFile("files/log.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(arquivo))
	main()
}

func capturaSites() []string {
	var sites []string
	arquivo, err := os.OpenFile("files/sites.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	leitor := bufio.NewReader(arquivo)
	for{
		linha, err := leitor.ReadString('\n') //Bytes são representados com aspas simples
		sites = append(sites, strings.TrimSpace(linha))
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	mensagem_status := "offline"
	arquivo, err := os.OpenFile("files/log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	
	if status {
		mensagem_status = "online"
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 - 15:04:05") + " - " + site + " - " + mensagem_status + "\n")

	arquivo.Close()
}


// Slices em Go eles são Arrays
// Quando é necessário colocar mais elementos do que sua capacidade atual, o slice dobra a capacidade
//  func exibiNome() {
// 	nomes := []string{"Enzo", "Bruna", "Naruto"}
// 	fmt.Println("aaaaaaa",len(nomes))
// 	nomes = append(nomes, "Batman")
// 	fmt.Println(len(nomes))
// 	fmt.Println(reflect.TypeOf(nomes))
// 	fmt.Println(nomes[0])
// 	fmt.Println(cap(nomes))
// }
