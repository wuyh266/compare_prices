package dao

import (
	"compare_prices/server_go/model"

	"gorm.io/gorm"
)

type ProductDAO struct {
	db *gorm.DB
}

func NewProductDAO(db *gorm.DB) *ProductDAO {
	return &ProductDAO{db: db}
}

func (d *ProductDAO) GetByCategory(category string) ([]model.Product, error) {
	var products []model.Product
	err := d.db.Where("category = ?", category).Find(&products).Error
	return products, err
}

func (d *ProductDAO) GetByCategoryAndBrand(category, brand string) ([]model.Product, error) {
	var products []model.Product
	err := d.db.Where("category = ? AND brand = ?", category, brand).Find(&products).Error
	return products, err
}

func (d *ProductDAO) GetByFilter(query *model.FilterQuery, category string) ([]model.Product, error) {
	var products []model.Product
	db := d.db.Model(&model.Product{}).Where("category = ?", category)

	if query.MinPrice > 0 {
		db = db.Where("price >= ?", query.MinPrice)
	}
	if query.MaxPrice > 0 {
		db = db.Where("price <= ?", query.MaxPrice)
	}
	if query.Brand != "" {
		db = db.Where("brand = ?", query.Brand)
	}
	if query.MinSales > 0 {
		db = db.Where("sales >= ?", query.MinSales)
	}
	if query.MinRating > 0 {
		db = db.Where("rating >= ?", query.MinRating)
	}
	if query.MustOfficialStore {
		db = db.Where("is_official_store = ?", true)
	}

	err := db.Find(&products).Error
	return products, err
}

func (d *ProductDAO) GetOfficialStoreProducts(category string) ([]model.Product, error) {
	var products []model.Product
	err := d.db.Where("category = ? AND is_official_store = ?", category, true).Find(&products).Error
	return products, err
}

func (d *ProductDAO) GetHotProducts(category string, limit int) ([]model.Product, error) {
	var products []model.Product
	err := d.db.Where("category = ?", category).Order("sales DESC").Limit(limit).Find(&products).Error
	return products, err
}

func (d *ProductDAO) Create(product *model.Product) error {
	return d.db.Create(product).Error
}

func (d *ProductDAO) BatchCreate(products []model.Product) error {
	return d.db.Create(&products).Error
}
