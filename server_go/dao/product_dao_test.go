package dao

import (
	"compare_prices/server_go/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockProductDAO struct {
	products []model.Product
}

func NewMockProductDAO() *MockProductDAO {
	return &MockProductDAO{
		products: []model.Product{},
	}
}

func (m *MockProductDAO) GetByCategory(category string) ([]model.Product, error) {
	var result []model.Product
	for _, p := range m.products {
		if p.Category == category {
			result = append(result, p)
		}
	}
	return result, nil
}

func (m *MockProductDAO) GetByFilter(query *model.FilterQuery, category string) ([]model.Product, error) {
	var result []model.Product
	for _, p := range m.products {
		if p.Category != category {
			continue
		}
		if query.MinPrice > 0 && p.Price < query.MinPrice {
			continue
		}
		if query.MaxPrice > 0 && p.Price > query.MaxPrice {
			continue
		}
		if query.Brand != "" && p.Brand != query.Brand {
			continue
		}
		if query.MinSales > 0 && p.Sales < query.MinSales {
			continue
		}
		if query.MinRating > 0 && p.Rating < query.MinRating {
			continue
		}
		result = append(result, p)
	}
	return result, nil
}

func (m *MockProductDAO) BatchCreate(products []model.Product) error {
	m.products = append(m.products, products...)
	return nil
}

func TestMockProductDAO_CRUD(t *testing.T) {
	dao := NewMockProductDAO()

	testProducts := []model.Product{
		{ProductID: "p1", Title: "Nike Air Max", Price: 899, Category: "运动鞋", Brand: "Nike", Sales: 12560, Rating: 4.9},
		{ProductID: "p2", Title: "Adidas Ultraboost", Price: 1299, Category: "运动鞋", Brand: "Adidas", Sales: 8500, Rating: 4.8},
		{ProductID: "p3", Title: "Jordan 1", Price: 1599, Category: "运动鞋", Brand: "Nike", Sales: 25000, Rating: 4.95},
	}

	t.Run("batch create products", func(t *testing.T) {
		err := dao.BatchCreate(testProducts)
		assert.NoError(t, err)
		assert.Len(t, dao.products, 3)
	})

	t.Run("get by category", func(t *testing.T) {
		products, err := dao.GetByCategory("运动鞋")
		assert.NoError(t, err)
		assert.Len(t, products, 3)
	})
}

func TestMockProductDAO_Filter(t *testing.T) {
	dao := NewMockProductDAO()
	testProducts := []model.Product{
		{ProductID: "p1", Title: "P1", Price: 500, Category: "运动鞋", Brand: "Nike", Sales: 100, Rating: 4.5},
		{ProductID: "p2", Title: "P2", Price: 1500, Category: "运动鞋", Brand: "Adidas", Sales: 5000, Rating: 4.9},
		{ProductID: "p3", Title: "P3", Price: 800, Category: "运动鞋", Brand: "Nike", Sales: 2000, Rating: 4.7},
	}
	dao.BatchCreate(testProducts)

	t.Run("filter by max price", func(t *testing.T) {
		query := &model.FilterQuery{MaxPrice: 1000}
		products, err := dao.GetByFilter(query, "运动鞋")
		assert.NoError(t, err)
		assert.Len(t, products, 2)
	})

	t.Run("filter by brand", func(t *testing.T) {
		query := &model.FilterQuery{Brand: "Nike"}
		products, err := dao.GetByFilter(query, "运动鞋")
		assert.NoError(t, err)
		assert.Len(t, products, 2)
	})

	t.Run("filter by min sales", func(t *testing.T) {
		query := &model.FilterQuery{MinSales: 1000}
		products, err := dao.GetByFilter(query, "运动鞋")
		assert.NoError(t, err)
		assert.Len(t, products, 2)
	})
}
