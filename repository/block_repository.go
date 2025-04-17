package repository

import (
	"block-chain/model"
	"database/sql"
	"encoding/hex"
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
	var prevHashStr, dataStr, hashStr string

	for row.Next() {
		err := row.Scan(
			&prevHashStr,
			&dataStr,
			&blockOBJ.Timestamp,
			&hashStr,
			&blockOBJ.Nonce,
		)
		blockOBJ.Previoushash, _ = hex.DecodeString(prevHashStr)
		blockOBJ.Hash, _ = hex.DecodeString(hashStr)
		blockOBJ.Data, _ = hex.DecodeString(dataStr)
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
		return model.Block{}, err
	}
	_, err = query.Exec(fmt.Sprintf("%x", block.Previoushash), fmt.Sprintf("%x", block.Data), block.Timestamp, fmt.Sprintf("%x", block.Hash), block.Nonce)
	if err != nil {
		return model.Block{}, errors.New("Erro")
	}
	query.Close()
	return block, nil

}
func (br *Blockrepository) Deleteall() error {
	_, err := br.connection.Exec("Delete from Block")
	if err != nil {
		return err
	}
	return nil

}

func (br *Blockrepository) GetByHash(aux string) ([]model.Block, error) {
	query := "SELECT previoushash,dados,timestampp,hash,nonce FROM Block WHERE dados=" + fmt.Sprintf("'%x'", []byte(aux))
	row, err := br.connection.Query(query)
	if err != nil {
		return []model.Block{}, err
	}
	var blockList []model.Block
	var blockOBJ model.Block
	var prevHashStr, dataStr, hashStr string
	for row.Next() {
		err = row.Scan(
			&prevHashStr,
			&dataStr,
			&blockOBJ.Timestamp,
			&hashStr,
			&blockOBJ.Nonce,
		)
		if err != nil {
			fmt.Println("Erro no scan")
		}
		blockOBJ.Previoushash, _ = hex.DecodeString(prevHashStr)
		blockOBJ.Hash, _ = hex.DecodeString(hashStr)
		blockOBJ.Data, _ = hex.DecodeString(dataStr)
		blockList = append(blockList, blockOBJ)
	}
	row.Close()
	return blockList, nil

}
func (br *Blockrepository) GetLastBlock() (model.Block, error) {
	query := "SELECT * FROM Block ORDER BY timestampp DESC LIMIT 1"
	row := br.connection.QueryRow(query)
	var blockOBJ model.Block
	var prevHashStr, dataStr, hashStr string

	row.Scan(
		&prevHashStr,
		&dataStr,
		&blockOBJ.Timestamp,
		&hashStr,
		&blockOBJ.Nonce,
	)
	blockOBJ.Previoushash, _ = hex.DecodeString(prevHashStr)
	blockOBJ.Hash, _ = hex.DecodeString(hashStr)
	blockOBJ.Data, _ = hex.DecodeString(dataStr)

	return blockOBJ, nil
}
