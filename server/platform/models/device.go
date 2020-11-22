package models

type Device struct {
	Token      string `json:"token" gorm:"primaryKey"`
	Is_enabled bool   `json:"is_enabled"`
}
