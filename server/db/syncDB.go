package db

import "server-ssm/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{}, &models.Post{})
}