package usecase

import (
	"fmt"

	"github.com/toannguyen3105/ytb-tradeit-crawler/internal/domain"
	"github.com/toannguyen3105/ytb-tradeit-crawler/internal/repository"
)

type CrawlItemsUsecase struct {
	Repo repository.ItemRepository
}

func NewCrawlItemsUsecase(repo repository.ItemRepository) *CrawlItemsUsecase {
	return &CrawlItemsUsecase{Repo: repo}
}

func (uc *CrawlItemsUsecase) FetchAllItems() ([]domain.Item, error) {
	var allItems []domain.Item
	offset := 0
	limit := 160

	for {
		fmt.Printf("Fetching offset %d...\n", offset)
		items, err := uc.Repo.FetchItems(offset, limit)
		if err != nil {
			return nil, err
		}

		if len(items) == 0 {
			break
		}

		allItems = append(allItems, items...)
		offset += limit
	}

	return allItems, nil
}
