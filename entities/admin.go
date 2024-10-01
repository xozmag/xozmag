package entities

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
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
	ID string `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Photo string `json:"photo" gorm:"column:photo"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type SubCategory struct {
	ID string `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Photo string `json:"photo" gorm:"column:photo"`
	CategoryId string `json:"categoryId" gorm:"category_id"`
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