package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/toannguyen3105/ytb-tradeit-crawler/internal/infrastructure"
	"github.com/toannguyen3105/ytb-tradeit-crawler/internal/usecase"
)

func main() {
	start := time.Now()

	repo := infrastructure.NewTradeitRepository()
	uc := usecase.NewCrawlItemsUsecase(repo)

	items, err := uc.FetchAllItems()
	if err != nil {
		log.Fatalf("Error fetching items: %v", err)
	}

	fmt.Printf("Total items fetched: %d\n", len(items))

	// Save to CSV
	file, err := os.Create("items.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// header
	writer.Write([]string{"Name", "Price"})

	for _, item := range items {
		writer.Write([]string{item.Name, fmt.Sprintf("%d", item.Price)})
	}

	fmt.Println("Saved to items.csv")

	elapsed := time.Since(start)
	fmt.Printf("Done! Execution time is %.2fs\n", elapsed.Seconds())
}
