package funcs

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"os"
	"time"
)

type Block struct {
	Previoushash []byte `json:"Previoushash"`
	Data         []byte `json:"Data"`
	Timestamp    int64  `json:"Timestamp"`
	Hash         []byte `json:"Hash"`
	Nonce        int    `json:"Nonce"`
	/*Proximo      *Block*/
}

const targetBits = 23

type BlockChain struct {
	Block []Block
}
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

const maxNonce = math.MaxInt64

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)                  //gera um big int de valor
	target.Lsh(target, uint(256-targetBits)) //shifta 1 por 256 - targetbits

	pow := &ProofOfWork{b, target} //armazena um ProofOfWork em pow que contem um bloco e um target para ele bloco

	return pow //retorn pow
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte //gera um slice de tamanho 32 de byte para armazenar hash
	nonce := 0        //um contador para ir somando com os dados do bloco pow

	for nonce < maxNonce {
		data := pow.Gethash(nonce) //pega os dados do pow e faz dados+nonce=hash e armazena em data
		hash = sha256.Sum256(data) //hash recebe um checksum de data para armazenar em forma de um slice de byte com 32 de tamanho
		hashInt.SetBytes(hash[:])  //recebe o slice de hash e o interpreta em um big int que nesse caso é o hashInt que vai ter o valor interpreta de [32]byte em forma de big.Int

		if hashInt.Cmp(pow.Target) == -1 { //se o valor de hashInt for menor que pow.targe o hash é valido logo da break e retorna nonce e hash[:]
			break
		} else {
			nonce++ //caso nao seja incrementa nonce faz o mesmo processo dados + o nonce novo
		}
	}

	return nonce, hash[:]
}

func NewBlock(Data string) Block {
	//função usada para criar um novo bloco usando o ProofOfWork e pow.Run() para gerar o nonce e hash do bloco para retornar e armazenar no blocos.json
	arq, err := os.Open("arquivo/blocos.json")
	if err != nil {
		fmt.Println("Arquivo nao aberto")
	}
	blocos := BlockChain{}
	err = json.NewDecoder(arq).Decode(&blocos.Block)
	if err != nil {
		fmt.Println("Blocos.json vazio, adicionando o primeiro bloco!!!")
	}
	if len(blocos.Block) != 0 {
		block := Block{Timestamp: time.Now().Unix(), Previoushash: blocos.Block[len(blocos.Block)-1].Hash, Data: []byte(Data)}
		pow := NewProofOfWork(&block)
		nonce, hash := pow.Run()

		block.Hash = hash[:]
		block.Nonce = nonce
		return block
	} else {
		block := Block{Timestamp: time.Now().Unix(), Previoushash: []byte(""), Data: []byte(Data)}
		pow := NewProofOfWork(&block)
		nonce, hash := pow.Run()

		block.Hash = hash[:]
		block.Nonce = nonce
		return block
	}

}

func AddBlockGenesis() {
	//adiciona o primeiro bloco da blockchain em blocos.json
	arq, err := os.Open("arquivo/blocos.json")
	if err != nil {
		fmt.Println("Arquivo nao aberto")
	}
	blocos := BlockChain{}
	err = json.NewDecoder(arq).Decode(&blocos.Block)
	if err != nil {
		fmt.Println("Erro")
	}
	b := NewBlock("Data Genesis")
	blocos.Block = append(blocos.Block, b)
	vetor, err := json.MarshalIndent(blocos.Block, "", " ")
	if err != nil {
		fmt.Println("Erro")
	}
	err = os.WriteFile("arquivo/blocos.json", vetor, 0644)
	if err != nil {
		fmt.Println("erro")
	}
}
func Addblockjson(Data string) {
	//adiciona os demais blocos em blocos.json
	b := NewBlock(Data)
	arq, err := os.Open("arquivo/blocos.json")
	if err != nil {
		fmt.Println("Arquivo nao aberto")
	}
	blocos := BlockChain{}
	err = json.NewDecoder(arq).Decode(&blocos.Block)
	if err != nil {
		fmt.Println("Blocos.json vazio, adicionando o primeiro bloco!!!")
	}
	blocos.Block = append(blocos.Block, b)
	vetor, err := json.MarshalIndent(blocos.Block, "", " ")
	if err != nil {
		fmt.Println("Erro")
	}
	err = os.WriteFile("arquivo/blocos.json", vetor, 0644)
	if err != nil {
		fmt.Println("erro")
	}

}
func (pow *ProofOfWork) Gethash(nonce int) []byte {
	//gera o hash que é a junção dos campos data, PreviousHash , timestamp , nonce e o targetbits que é retornado como um []byte
	data := bytes.Join(
		[][]byte{
			pow.Block.Previoushash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data

}
func IntToHex(n int64) []byte {
	//transforma um int 64 em []byte
	buff := make([]byte, 8)                     //faz um []byte de tamanho 8 , logo [8]byte
	binary.BigEndian.PutUint64(buff, uint64(n)) //pega o buff que vai armazenar o uint64(int64 n) em forma de []byte
	return buff
}
func Listblocks() {
	//lista os blocos do json
	arq, err := os.Open("arquivo/blocos.json")
	if err != nil {
		fmt.Println("Arquivo nao aberto")
	}
	var blocos BlockChain
	err = json.NewDecoder(arq).Decode(&blocos.Block)
	if err != nil {
		fmt.Println("Erro")
	}
	for _, b := range blocos.Block {
		fmt.Printf("PreviousHash:%x,Data:%s,Timestamp:%v,Hash:%x\n", b.Previoushash, string(b.Data), b.Timestamp, b.Hash)

	}

}

//Daqui para baixo é uma versao antiga utilizando lista ligada simples de Blocks

/*func (b *BlockChain) Addblock(Index int, Data string, Timestamp string) {
	bn := &Block{Index: Index, Data: Data, Timestamp: Timestamp, Proximo: nil}
	if b.Initialblock == nil {
		bn.Previoushash = ""
		bn.Gethash()
		b.Initialblock = bn
		return
	}
	var aux *Block
	aux = b.Initialblock
	for {
		if aux.Proximo == nil {
			bn.Previoushash = aux.Hash
			bn.Gethash()
			aux.Proximo = bn
			break
		}
		aux = aux.Proximo
	}

}*/
/*func (b *BlockChain) Listblocks() {
	var aux *Block
	aux = b.Initialblock
	for {

		fmt.Printf("Index:%d,PreviousHash:%x,Data:%s,Timestamp:%s,Hash:%x\n", aux.Index, aux.Previoushash, aux.Data, aux.Timestamp, aux.Hash)
		if aux.Proximo == nil {
			break
		} else {
			aux = aux.Proximo
		}

	}
}*/
/*func (b *BlockChain) Init() {
	(b).Initialblock = nil

}*/
