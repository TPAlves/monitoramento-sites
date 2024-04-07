package main

import (
	"fmt"
	"os/user"
	"strings"
)

// func main(){
// 	// fmt.Scanf("%d", &comand) // o "%d" aceita somente digitos e o & diz respeito ao ponteiro/endereço da variavel
// 	// Faz a mesma coisa que o Scanf, porém não é necessario passar nenhum modificado, pois o tipo foi declarado na criação da variavel
// 	comand := optionsComands()

// 	fmt.Println("O comando escolhido foi",comand)
// }

func opcaoMenu() int {
	var currentUser, err = user.Current()
	var comand int = 0

	if err != nil {
		fmt.Println(err)
	}

	username := strings.ToLower(currentUser.Username)

	fmt.Println("Olá, Sr/Sra", username)

	for {
		fmt.Println("---------------------- MENU ----------------------")
		fmt.Println("1 - Iniciar Monitoramento")
		fmt.Println("2 - Exibir logs")
		fmt.Println("3 - Sair")

		fmt.Scan(&comand)

		if comand < 1 || comand > 3 {
			fmt.Println("Opção inválida. Por favor, digite uma opção entre 1 a 3")
		} else {
			break
		}

	}
	return comand
}
