package usecase

import (
	"block-chain/model"
	"block-chain/repository"
)

type Blockusecase struct {
	Repository repository.Blockrepository
}

func Newblockusecase(Repository repository.Blockrepository) Blockusecase {
	return Blockusecase{Repository: Repository}
}

func (bu *Blockusecase) GetBlocks() ([]model.Block, error) {
	return bu.Repository.GetBlocks()
}
func (bu *Blockusecase) InsertBlock(b model.Block) (model.Block, error) {
	block, err := bu.Repository.InsertBlock(b)
	if err != nil {
		return model.Block{}, err
	}
	return block, nil

}
