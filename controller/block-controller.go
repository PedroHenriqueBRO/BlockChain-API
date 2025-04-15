package controller

import (
	"block-chain/model"
	"block-chain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	var block model.Block
	err := ctx.BindJSON(&block)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
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
