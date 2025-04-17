package controller

import (
	"block-chain/model"
	"block-chain/usecase"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Stringaux struct {
	Aux string `json:"Data"`
}
type BlockController struct {
	Blockusecase usecase.Blockusecase
}

func NewblockController(usecase usecase.Blockusecase) BlockController {

	return BlockController{usecase}
}

func (b *BlockController) GetBlocks(ctx *gin.Context) {
	blocos, err := b.Blockusecase.GetBlocks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, blocos)

}
func (b *BlockController) InsertBlock(ctx *gin.Context) {

	defer time.Sleep(time.Second * 2)

	var block model.Block
	aux := Stringaux{}
	err := ctx.BindJSON(&aux)
	if err != nil {
		fmt.Println("Nao foi possível passar de JSON para Block")
	}
	bloco, err := b.Blockusecase.GetLastBlock()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	block.Data = []byte(aux.Aux)
	block.Previoushash = bloco.Hash
	block.NewBlock()
	insertedblock, err := b.Blockusecase.InsertBlock(block)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

	}
	ctx.JSON(http.StatusCreated, insertedblock)

}

func (b *BlockController) Deleteall(ctx *gin.Context) {
	err := b.Blockusecase.Deleteall()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

	}
}
func (bu *BlockController) GetByHash(ctx *gin.Context) {
	aux := Stringaux{}
	err := ctx.BindJSON(&aux)
	if err != nil {
		fmt.Println("Nao foi possível passar de JSON para Block")
	}
	blocos, err := bu.Blockusecase.GetByHash(aux.Aux)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, blocos)

}

func (bu *BlockController) GetLastBlock(ctx *gin.Context) {
	bloco, err := bu.Blockusecase.GetLastBlock()
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, bloco)

}
