package model

import "time"

type Manufacturer struct {
	ID              string    `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name"`
	OfficialWebsite string    `json:"official_website"`
	Description     string    `json:"description"`
	LogoURL         string    `json:"logo_url"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
