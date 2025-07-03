package usecase

import (
	"fmt"
	"sync"

	"github.com/toannguyen3105/ytb-tradeit-crawler/internal/domain"
	"github.com/toannguyen3105/ytb-tradeit-crawler/internal/repository"
)

const maxPages = 20

type CrawlItemsUsecase struct {
	Repo repository.ItemRepository
}

func NewCrawlItemsUsecase(repo repository.ItemRepository) *CrawlItemsUsecase {
	return &CrawlItemsUsecase{Repo: repo}
}

func (uc *CrawlItemsUsecase) FetchAllItems() ([]domain.Item, error) {
	var allItems []domain.Item
	limit := 160

	var wg sync.WaitGroup
	itemsChan := make(chan []domain.Item, maxPages)
	errChan := make(chan error, maxPages)

	for i := 0; i < maxPages; i++ {
		wg.Add(1)
		offset := i * limit
		go func(offset int) {
			defer wg.Done()
			fmt.Printf("Fetching offset %d...\n", offset)
			items, err := uc.Repo.FetchItems(offset, limit)
			if err != nil {
				errChan <- err
				return
			}
			if len(items) > 0 {
				itemsChan <- items
			}
		}(offset)
	}

	wg.Wait()
	close(itemsChan)
	close(errChan)

	for err := range errChan {
		return nil, err
	}

	for items := range itemsChan {
		allItems = append(allItems, items...)
	}

	return allItems, nil
}
