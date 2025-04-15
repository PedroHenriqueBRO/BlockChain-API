package main

import (
	"block-chain/controller"
	"block-chain/db"
	"block-chain/repository"
	"block-chain/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbconnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	blockrepository := repository.Newblockrepository(dbconnection)
	blockusecase := usecase.Newblockusecase(blockrepository)
	blockController := controller.NewblockController(blockusecase)

	server.GET("/ping", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "pong"}) })
	server.GET("/blocks", blockController.GetBlocks)
	server.POST("/block", blockController.InsertBlock)
	server.DELETE("/delete", blockController.Deleteall)
	server.Run(":8000")

}
