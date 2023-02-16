package db

import (
	"dousheng/pkg/configs/sql"
	"time"

	"dousheng/pkg/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	gormlogrus := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)
	DB, err = gorm.Open(mysql.Open(consts.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt: false,
			Logger:      gormlogrus,
		},
	)
	if err != nil {
		panic(err)
	}

	if err := DB.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&sql.User{}, &sql.Video{})
	if err != nil {
		panic(err)
	}
}
