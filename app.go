package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

func main() {

	consultaDiretorio()

	comand := opcaoMenu()

	switch comand {
	case 1:
		iniciarMonitoramento()
	case 2:
		exibirLogs()
	case 3:
		fmt.Println("Tchau!")
		os.Exit(0)
	}
}

func consultaDiretorio() {
	logger := slog.Default()

	cmd := "ls | grep -e ^files$"

	_, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		logger.Warn("Diretório files inexistente")
		logger.Info("Criando diretório....")
		err = os.Mkdir("files", 0700)
		if err != nil {
			logger.Error(err.Error())
			logger.Error("Não foi possível criar o diretório")
			os.Exit(2)
		}
	}

}
