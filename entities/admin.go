package entities

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Xozmak struct {
	ID        string         `json:"id" gorm:"column:id"`
	Name      string         `json:"name" gorm:"column:name"`
	CreatedBy sql.NullString `gorm:"column:created_by"`
	UpdatedBy sql.NullString `gorm:"column:updated_by"`
	Location  Location       `json:"location" gorm:"column:location;type:json"`
}

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Category struct {
	ID        string    `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Photo     string    `json:"photo" gorm:"column:photo"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type SubCategory struct {
	ID         string    `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`
	Photo      string    `json:"photo" gorm:"column:photo"`
	CategoryId string    `json:"categoryId" gorm:"category_id"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (l *Location) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan Location, unexpected type %T", value)
	}
	if err := json.Unmarshal(bytes, l); err != nil {
		return fmt.Errorf("failed to unmarshal Location JSON: %w", err)
	}
	return nil
}

func (l Location) Value() (driver.Value, error) {
	return json.Marshal(l)
}

type Product struct {
	ID             int64           `json:"id" gorm:"column:id"`
	ProductID      uuid.UUID       `json:"product_id" gorm:"column:product_id"`
	CategoryID     uuid.UUID       `json:"category_id" gorm:"column:category_id"`
	SubCategoryID  *uuid.UUID      `json:"sub_category_id" gorm:"column:sub_category_id"`
	CreatedAt      time.Time       `json:"created_at" gorm:"column:created_at"`
	CreatedBy      *uuid.UUID      `json:"created_by" gorm:"column:created_by"`
	UpdatedAt      *time.Time      `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy      *uuid.UUID      `json:"updated_by" gorm:"column:updated_by"`
	NameKr         *string         `json:"name_kr" gorm:"column:name_kr"`
	NameLt         string          `json:"name_lt" gorm:"column:name_lt"`
	NameRu         *string         `json:"name_ru" gorm:"column:name_ru"`
	ShortInfoKr    *string         `json:"short_info_kr" gorm:"column:short_info_kr"`
	ShortInfoLt    string          `json:"short_info_lt" gorm:"column:short_info_lt"`
	ShortInfoRu    *string         `json:"short_info_ru" gorm:"column:short_info_ru"`
	DescriptionKr  *string         `json:"description_kr" gorm:"column:description_kr"`
	DescriptionLt  string          `json:"description_lt" gorm:"column:description_lt"`
	DescriptionRu  *string         `json:"description_ru" gorm:"column:description_ru"`
	Status         *string         `json:"status" gorm:"column:status"`
	ProductDetails []ProductDetail `json:"product_details" gorm:"-"`
	Files          []File          `json:"files" gorm:"-"`
}

type File struct {
	ID        int64      `json:"id" gorm:"column:id"`
	FileID    uuid.UUID  `json:"file_id" gorm:"column:file_id"`
	Name      *string    `json:"name" gorm:"column:name"`
	FilePath  *string    `json:"file_path" gorm:"column:file_path"`
	Size      *int64     `json:"size" gorm:"column:size"`
	Extension *string    `json:"extension" gorm:"column:extension"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	CreatedBy *uuid.UUID `json:"created_by" gorm:"column:created_by"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy *uuid.UUID `json:"updated_by" gorm:"column:updated_by"`
	IsMain    bool       `json:"is_main" gorm:"column:is_main"`
	AttachID  *uuid.UUID `json:"attach_id" gorm:"column:attach_id"`
}

type ProductDetail struct {
	ID                     int64      `json:"id" gorm:"column:id"`
	ProductID              uuid.UUID  `json:"product_id" gorm:"column:product_id"`
	Color                  *int       `json:"color" gorm:"column:color"`
	Weight                 *float64   `json:"weight" gorm:"column:weight"`
	Capacity               *float64   `json:"capacity" gorm:"column:capacity"`
	TwoDimensionalHeight   *float64   `json:"two_dimensional_height" gorm:"column:two_dimensional_height"`
	TwoDimensionalWidth    *float64   `json:"two_dimensional_width" gorm:"column:two_dimensional_width"`
	ThreeDimensionalHeight *float64   `json:"three_dimensional_height" gorm:"column:three_dimensional_height"`
	ThreeDimensionalWidth  *float64   `json:"three_dimensional_width" gorm:"column:three_dimensional_width"`
	ThreeDimensionalThick  *float64   `json:"three_dimensional_thick" gorm:"column:three_dimensional_thick"`
	Amount                 int64      `json:"amount" gorm:"column:amount"`
	Price                  float64    `json:"price" gorm:"column:price"`
	DiscountPrice          *float64   `json:"discount_price" gorm:"column:discount_price"`
	DiscountPercent        *float64   `json:"discount_percent" gorm:"column:discount_percent"`
	CreatedAt              time.Time  `json:"created_at" gorm:"column:created_at"`
	CreatedBy              *uuid.UUID `json:"created_by" gorm:"column:created_by"`
	UpdatedAt              *time.Time `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy              *uuid.UUID `json:"updated_by" gorm:"column:updated_by"`
}
