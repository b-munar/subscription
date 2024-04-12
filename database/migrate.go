package database

import (
	"subscription/model"
)

func Migrate() {
	ConnectDB()

	db := DB

	db.AutoMigrate(&model.Subscription{})

}
