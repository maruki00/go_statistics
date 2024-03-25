package models

import "gorm.io/gorm"

type AppStatistic struct {
	gorm.Model
	Id             int
	FullName       string
	Shop           string
	Position       string
	UserID         int
	ShopID         int
	Region         string
	InstalledCount int
}

type Regions struct {
	gorm.Model
	Id       int
	Region   string
	RegionId int
}
