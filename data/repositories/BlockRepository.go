package repositories

import (
  "chain/domain"
  "math"
)

type SearchResult struct {
  Total int
  Count int
  From int
  Results []domain.MinedBlock
}

type BlockRepository struct {
  BlockChain []domain.MinedBlock
}

func (r BlockRepository) Find(sequence int) domain.MinedBlock  {
  var block domain.MinedBlock

  if len(r.BlockChain) <= sequence || sequence == 0 {
    return block
  }

  return r.BlockChain[sequence - 1]
}

func (r BlockRepository) Search(page int, size int) SearchResult {
  total := len(r.BlockChain)

  pages := int(math.Ceil(float64(total) / float64(size)))

  if page >= pages {
    return SearchResult{Results: []domain.MinedBlock{}}
  }

  start := page * (size - 1)
  end := start + size; if end > total {
    end = total
  }

  blocks := r.BlockChain[start:end]

  return SearchResult{
    Total: total,
    Results: blocks,
    From: start,
    Count: len(blocks),
  }
}
