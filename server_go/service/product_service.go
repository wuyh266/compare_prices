package service

import (
	"compare_prices/server_go/dao"
	"compare_prices/server_go/model"
	"sort"
	"sync"
)

type ProductService struct {
	productDAO *dao.ProductDAO
	taskDAO    *dao.RecognizeTaskDAO
}

func NewProductService(productDAO *dao.ProductDAO, taskDAO *dao.RecognizeTaskDAO) *ProductService {
	return &ProductService{
		productDAO: productDAO,
		taskDAO:    taskDAO,
	}
}

func (s *ProductService) AggregateProducts(taskID, filterType, sortBy string) (*model.ProductListResponse, error) {
	task, err := s.taskDAO.GetByTaskID(taskID)
	if err != nil {
		return nil, err
	}

	var products []model.Product

	switch filterType {
	case "card_lowest_price":
		products, err = s.productDAO.GetByCategory(task.Category)
	case "card_official_store":
		products, err = s.productDAO.GetOfficialStoreProducts(task.Category)
	case "card_hot_recommend":
		products, err = s.productDAO.GetHotProducts(task.Category, 50)
	default:
		products, err = s.productDAO.GetByCategory(task.Category)
	}

	if err != nil {
		return nil, err
	}

	sortProducts(products, sortBy)

	return s.calculatePriceStats(products), nil
}

func (s *ProductService) FilterProductsByQuery(taskID string, filterQuery *model.FilterQuery) (*model.ProductListResponse, error) {
	task, err := s.taskDAO.GetByTaskID(taskID)
	if err != nil {
		return nil, err
	}

	products, err := s.productDAO.GetByFilter(filterQuery, task.Category)
	if err != nil {
		return nil, err
	}

	sortProducts(products, "")

	return s.calculatePriceStats(products), nil
}

func (s *ProductService) calculatePriceStats(products []model.Product) *model.ProductListResponse {
	if len(products) == 0 {
		return &model.ProductListResponse{
			PlatformMinPrice: 0,
			PlatformAvgPrice: 0,
			Products:         []model.Product{},
		}
	}

	var totalPrice float64
	minPrice := products[0].Price

	for _, p := range products {
		totalPrice += p.Price
		if p.Price < minPrice {
			minPrice = p.Price
		}
	}

	avgPrice := totalPrice / float64(len(products))

	return &model.ProductListResponse{
		PlatformMinPrice: minPrice,
		PlatformAvgPrice: avgPrice,
		Products:         products,
	}
}

func sortProducts(products []model.Product, sortBy string) {
	sort.Slice(products, func(i, j int) bool {
		switch sortBy {
		case "price_asc":
			return products[i].Price < products[j].Price
		case "price_desc":
			return products[i].Price > products[j].Price
		case "sales":
			return products[i].Sales > products[j].Sales
		case "rating":
			return products[i].Rating > products[j].Rating
		default:
			return products[i].Price < products[j].Price
		}
	})
}

func (s *ProductService) ParallelFetchFromPlatforms(category string) ([]model.Product, error) {
	platforms := []string{"jd", "taobao", "pdd"}
	var wg sync.WaitGroup
	var mu sync.Mutex
	var allProducts []model.Product

	for _, platform := range platforms {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			products, err := s.simulatePlatformFetch(p, category)
			if err == nil {
				mu.Lock()
				allProducts = append(allProducts, products...)
				mu.Unlock()
			}
		}(platform)
	}

	wg.Wait()
	return allProducts, nil
}

func (s *ProductService) simulatePlatformFetch(platform, category string) ([]model.Product, error) {
	return s.productDAO.GetByCategory(category)
}
