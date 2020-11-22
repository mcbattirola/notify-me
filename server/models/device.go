package models

type Device struct {
	ID         string `json:"id" gorm:"primary_key"`
	Is_Enabled bool   `json:"is_enabled"`
}
