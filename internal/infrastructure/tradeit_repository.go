package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/toannguyen3105/ytb-tradeit-crawler/internal/domain"
	"github.com/toannguyen3105/ytb-tradeit-crawler/internal/repository"
)

type TradeitRepository struct {
	BaseURL string
}

func NewTradeitRepository() repository.ItemRepository {
	return &TradeitRepository{
		BaseURL: "https://tradeit.gg/api/v2/inventory/data",
	}
}

type apiResponse struct {
	Items []struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
	} `json:"items"`
}

func (t *TradeitRepository) FetchItems(offset, limit int) ([]domain.Item, error) {
	params := url.Values{}
	params.Set("gameId", "730")
	params.Set("offset", strconv.Itoa(offset))
	params.Set("limit", strconv.Itoa(limit))
	params.Set("sortType", "Price - high")
	params.Set("minPrice", "100")
	params.Set("maxPrice", "5000")
	params.Set("minFloat", "0")
	params.Set("maxFloat", "1")
	params.Set("showTradeLock", "true")
	params.Set("onlyTradeLock", "false")
	params.Set("showUserListing", "true")
	params.Set("context", "trade")
	params.Set("fresh", "false")
	params.Set("isForStore", "0")

	url := fmt.Sprintf("%s?%s", t.BaseURL, params.Encode())
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result apiResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	var items []domain.Item
	for _, apiItem := range result.Items {
		items = append(items, domain.Item{
			Name:  apiItem.Name,
			Price: apiItem.Price,
		})
	}

	return items, nil
}
