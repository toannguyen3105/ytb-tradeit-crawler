package repository

import "github.com/toannguyen3105/ytb-tradeit-crawler/internal/domain"

type ItemRepository interface {
	FetchItems(offset, limit int) ([]domain.Item, error)
}
