package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
	StartDate time.Time `form:"start_date" json:"start_date"`
	EndDate time.Time `form:"end_date" json:"end_date"`
	Done bool `form:"done" json:"done"`
	WorkerID uint `form:"worker_id" json:"worker_id"`
	Worker Worker `gorm:"foreignKey:WorkerID;references:ID"`
}