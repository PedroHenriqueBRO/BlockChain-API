package repository

import (
	"block-chain/model"
	"database/sql"
	"errors"
	"fmt"
)

type Blockrepository struct {
	connection *sql.DB
}

func Newblockrepository(connection *sql.DB) Blockrepository {
	return Blockrepository{connection: connection}

}
func (br *Blockrepository) GetBlocks() ([]model.Block, error) {
	query := "SELECT previoushash,dados,timestampp,hash,nonce FROM Block"
	row, err := br.connection.Query(query)
	if err != nil {
		return []model.Block{}, err
	}
	var blockList []model.Block
	var blockOBJ model.Block
	for row.Next() {
		err := row.Scan(
			&blockOBJ.Previoushash,
			&blockOBJ.Data,
			&blockOBJ.Timestamp,
			&blockOBJ.Hash,
			&blockOBJ.Nonce,
		)
		if err != nil {
			return []model.Block{}, err
		}
		blockList = append(blockList, blockOBJ)

	}
	row.Close()
	return blockList, nil

}
func (br *Blockrepository) InsertBlock(block model.Block) (model.Block, error) {
	query, err := br.connection.Prepare("INSERT INTO Block" + "(previoushash,dados,timestampp,hash,nonce )" + "VALUES ($1,$2,$3,$4,$5)")
	if err != nil {
		fmt.Println("aqui query")
		return model.Block{}, err
	}
	_, err = query.Exec(fmt.Sprintf("%x", block.Previoushash), fmt.Sprintf("%x", block.Data), block.Timestamp, fmt.Sprintf("%x", block.Hash), block.Nonce)
	if err != nil {
		fmt.Println("aqui", err)
		return model.Block{}, errors.New("Erro")
	}
	query.Close()
	return block, nil

}
