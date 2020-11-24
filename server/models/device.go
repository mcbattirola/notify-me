package models

type Device struct {
	ID        string `json:"id" gorm:"primary_key"`
	IsEnabled bool   `json:"is_enabled"`
}
