package models

import "gorm.io/gorm"

type Worker struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	Position string `json:"position" form:"position"`
	Gender string `json:"gender" form:"gender"`
	Address string `json:"address" form:"address"`
	Tasks []Task `gorm:"foreignKey:WorkerID;references:ID"`
}