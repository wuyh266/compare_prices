package service

import (
	"compare_prices/server_go/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortProducts(t *testing.T) {
	products := []model.Product{
		{Title: "Product A", Price: 100, Sales: 1000, Rating: 4.5},
		{Title: "Product B", Price: 50, Sales: 5000, Rating: 4.9},
		{Title: "Product C", Price: 200, Sales: 100, Rating: 4.7},
	}

	t.Run("sort by price asc", func(t *testing.T) {
		sortProducts(products, "price_asc")
		assert.Equal(t, 50.0, products[0].Price)
		assert.Equal(t, 100.0, products[1].Price)
		assert.Equal(t, 200.0, products[2].Price)
	})

	t.Run("sort by price desc", func(t *testing.T) {
		sortProducts(products, "price_desc")
		assert.Equal(t, 200.0, products[0].Price)
		assert.Equal(t, 100.0, products[1].Price)
		assert.Equal(t, 50.0, products[2].Price)
	})

	t.Run("sort by sales", func(t *testing.T) {
		sortProducts(products, "sales")
		assert.Equal(t, uint(5000), products[0].Sales)
		assert.Equal(t, uint(1000), products[1].Sales)
		assert.Equal(t, uint(100), products[2].Sales)
	})

	t.Run("sort by rating", func(t *testing.T) {
		sortProducts(products, "rating")
		assert.Equal(t, 4.9, products[0].Rating)
		assert.Equal(t, 4.7, products[1].Rating)
		assert.Equal(t, 4.5, products[2].Rating)
	})
}

func TestCalculatePriceStats(t *testing.T) {
	products := []model.Product{
		{Title: "P1", Price: 100},
		{Title: "P2", Price: 200},
		{Title: "P3", Price: 300},
	}

	svc := &ProductService{}
	result := svc.calculatePriceStats(products)

	assert.Equal(t, 100.0, result.PlatformMinPrice)
	assert.Equal(t, 200.0, result.PlatformAvgPrice)
	assert.Len(t, result.Products, 3)
}

func TestCalculatePriceStatsEmpty(t *testing.T) {
	products := []model.Product{}
	svc := &ProductService{}
	result := svc.calculatePriceStats(products)

	assert.Equal(t, 0.0, result.PlatformMinPrice)
	assert.Equal(t, 0.0, result.PlatformAvgPrice)
	assert.Empty(t, result.Products)
}
