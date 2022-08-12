package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name string
	Description string
	StartDate time.Time
	EndDate time.Time
	Done bool
	WorkerID uint
}