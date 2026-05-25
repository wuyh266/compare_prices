package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type JSONTags []string

func (j JSONTags) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONTags) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	return json.Unmarshal(bytes, j)
}

type Product struct {
	ID              uint64         `gorm:"primaryKey;autoIncrement" json:"-"`
	ProductID       string         `gorm:"type:varchar(64);uniqueIndex:uk_product_id;not null" json:"id"`
	Title           string         `gorm:"type:varchar(255);not null" json:"title"`
	ImageURL        string         `gorm:"type:varchar(512);not null" json:"image_url"`
	Price           float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	Platform        string         `gorm:"type:varchar(32);index:idx_platform;not null" json:"-"`
	PlatformName    string         `gorm:"type:varchar(32);not null" json:"platform"`
	Sales           uint           `gorm:"not null;default:0" json:"sales"`
	Rating          float64        `gorm:"type:decimal(3,2);not null;default:5.00" json:"rating"`
	DetailURL       string         `gorm:"type:varchar(512);not null" json:"detail_url"`
	Category        string         `gorm:"type:varchar(128);index:idx_category;not null" json:"-"`
	Brand           string         `gorm:"type:varchar(128);index:idx_brand" json:"-"`
	Tags            JSONTags       `gorm:"type:json" json:"tags"`
	IsOfficialStore bool           `gorm:"not null;default:false" json:"-"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
}

type RecognizeTask struct {
	ID            uint64                 `gorm:"primaryKey;autoIncrement" json:"-"`
	TaskID        string                 `gorm:"type:varchar(64);uniqueIndex:uk_task_id;not null" json:"task_id"`
	Category      string                 `gorm:"type:varchar(128);not null" json:"category"`
	Attributes    map[string]interface{} `gorm:"type:json;not null" json:"attributes"`
	RawImageURL   string                 `gorm:"type:varchar(512);not null" json:"-"`
	CreatedAt     time.Time              `json:"-"`
}

type ProductPriceHistory struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"-"`
	ProductID  string    `gorm:"type:varchar(64);uniqueIndex:uk_product_date;not null" json:"-"`
	PriceDate  time.Time `gorm:"type:date;uniqueIndex:uk_product_date;not null" json:"-"`
	Price      float64   `gorm:"type:decimal(10,2);not null" json:"-"`
	CreatedAt  time.Time `json:"-"`
}

type SuggestCard struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type RecognizeResult struct {
	TaskID       string                 `json:"task_id"`
	Category     string                 `json:"category"`
	Attributes   map[string]interface{} `json:"attributes"`
	SuggestCards []SuggestCard          `json:"suggest_cards"`
}

type FilterQuery struct {
	MinPrice          float64 `json:"min_price"`
	MaxPrice          float64 `json:"max_price"`
	Brand             string  `json:"brand"`
	Color             string  `json:"color"`
	MinSales          uint    `json:"min_sales"`
	MinRating         float64 `json:"min_rating"`
	MustOfficialStore bool    `json:"must_official_store"`
	ExcludeSecondhand bool    `json:"exclude_secondhand"`
}

type ProductListResponse struct {
	PlatformMinPrice float64   `json:"platform_min_price"`
	PlatformAvgPrice float64   `json:"platform_avg_price"`
	Products         []Product `json:"products"`
}

type AutoMigrateFunc func(db *gorm.DB) error

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Product{},
		&RecognizeTask{},
		&ProductPriceHistory{},
	)
}
