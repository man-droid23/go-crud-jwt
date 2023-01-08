package migration

import (
	"fiber-rest/configs/db"
	"fiber-rest/models/entity"
)

func Migration() {
	err := db.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic("Failed to migrate database")
	}
}
