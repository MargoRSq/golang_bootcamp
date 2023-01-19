package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	Id      int64  `gorm:"primaryKey"`
	Author  string `gorm:"column:author"`
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func createConnection() (err error) {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable",
		credentials.DBLogin, credentials.DBPassword, "gopg")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(&Post{})
	return
}

func addNewPost(newPost Post) (err error) {
	err = DB.Create(&newPost).Error
	return
}

func fetchPosts() {
	DB.Model(&Post{}).Find(&Articles)
}
