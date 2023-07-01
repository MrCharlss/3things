package dal

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(){
    fmt.Println("Starting Connection")
    db, err := gorm.Open(sqlite.Open("3thingstoday.db"),&gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	} )
    if err != nil {
        fmt.Println("[DATABASE]::CONNECTION_ERROR")
        panic(err)
    }

	DB = db

	fmt.Println("[DATABASE]::CONNECTED")
}

func Migrate(tables ...interface{}) error  {
    return DB.AutoMigrate(tables...)
}
