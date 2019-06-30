package services

import (
  "chain/data/repositories"
  "chain/domain"
)

type BlockService struct {
  Repository repositories.BlockRepository
}

func (blockService BlockService) Find(sequence int) domain.MinedBlock {
  block := blockService.Repository.Find(sequence)

  return block
}

func (blockService BlockService) Search(page int, size int) repositories.SearchResult {
  return blockService.Repository.Search(page, size)
}
