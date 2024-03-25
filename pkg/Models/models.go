package models

import "gorm.io/gorm"

type AppStatistic struct {
	gorm.Model
	Id             int    `json:"id"`
	FullName       string `json:"full_name"`
	Shop           string `json:"shop"`
	Position       string `json:"position"`
	UserID         int    `json:"user_id"`
	ShopID         int    `json:"shop_id"`
	Region         string `json:"region"`
	InstalledCount int    `json:"instelled_count"`
}

type Regions struct {
	gorm.Model
	Id       int    `json:"id"`
	Region   string `json:"region"`
	RegionId int    `json:"region_id"`
}
