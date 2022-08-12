package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_config "github.com/wildanie12/dts-task-planner/config"
)

func Connect(config *_config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Pass,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:              time.Second,   	// Slow SQL threshold
				LogLevel:                   logger.Silent, 	// Log level
				IgnoreRecordNotFoundError: 	true,           // Ignore ErrRecordNotFound error for logger
				Colorful:                  	false,          // Disable color
			},
		),
	})
	if err != nil {
		panic(err)
	}
	return db
}