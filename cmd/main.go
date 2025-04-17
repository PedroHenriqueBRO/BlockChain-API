package main

import (
	"block-chain/controller"
	"block-chain/db"
	"block-chain/model"
	"block-chain/repository"
	"block-chain/usecase"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbconnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	bloco := model.Block{Timestamp: time.Now().Unix(), Data: []byte("Genesis"), Previoushash: []byte("")}
	pow := model.NewProofOfWork(&bloco)
	bloco.Nonce, bloco.Hash = pow.Run()
	blockrepository := repository.Newblockrepository(dbconnection)
	blockusecase := usecase.Newblockusecase(blockrepository)
	blockController := controller.NewblockController(blockusecase)
	_, err = blockrepository.InsertBlock(bloco)
	if err != nil {
		fmt.Println(err)
	}
	server.GET("/blocks", blockController.GetBlocks)
	server.POST("/block", blockController.InsertBlock)
	server.DELETE("/delete", blockController.Deleteall)
	server.GET("/getbyhash", blockController.GetByHash)
	server.Run(":8000")

}
