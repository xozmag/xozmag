package entities

import "database/sql"

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
