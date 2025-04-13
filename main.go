package main

import (
	"block-chain/funcs"
)

func main() {
	funcs.AddBlockGenesis()
	funcs.Addblockjson("Alice envia 10 moedas para Bob")
	funcs.Addblockjson("Bob envia 5 moedas para Carol")
	funcs.Addblockjson("Carol envia 2 moedas para Dave")
	funcs.Addblockjson("Eve envia 7 moedas para Frank")
	funcs.Addblockjson("Testando transação com dados maiores para ver o comportamento da blockchain em diferentes situações")
	funcs.Addblockjson("Usuário123 envia 999 moedas para EndereçoXYZ")
	funcs.Addblockjson("Bloco de teste com caracteres especiais: !@#$%^&*()_+")
	funcs.Addblockjson("Transação simulada para análise de performance")
	funcs.Addblockjson("Recompensa de mineração de 50 moedas")
	funcs.Addblockjson("Finalizando série de testes automáticos")
	funcs.Addblockjson("Nova transação: João -> Maria, 20 moedas")
	funcs.Addblockjson("Teste de transação sem valor monetário")
	funcs.Addblockjson("Atualização de saldo para carteira XYZ")
	funcs.Addblockjson("Transação revertida por inconsistência")
	funcs.Addblockjson("Registro de auditoria interna")

	funcs.Listblocks()

}
