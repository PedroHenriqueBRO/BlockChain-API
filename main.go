package main

import (
	"block-chain/funcs"
	"fmt"
)

func main() {
	var E string
	fmt.Println("Bem vindo a Block Chain!!")
	for {
		var v bool
		v = true
		fmt.Println("Deseja adicionar blocos ou listar? 1(Adicionar),2(Listar),3(stop)!!!!")
		fmt.Scan(&E)
		switch E {
		case "1":
			var data string
			fmt.Println("Digite o dado do bloco!!!!")
			fmt.Scan(&data)
			funcs.Addblockjson(data)

		case "2":
			funcs.Listblocks()
		case "3":
			v = false

		}
		if !v {
			break
		}
	}

}
