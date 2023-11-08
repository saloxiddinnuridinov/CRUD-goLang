package main

import "gorm.io/gorm"

func seedUsers(db *gorm.DB) {
	Users := []User{
		{FirstName: "Salohiddin", LastName: "Nuridinov", Email: "admin@gmail.com"},
		{FirstName: "Salohiddin", LastName: "Nuridinov", Email: "admin@gmail.com"},
		{FirstName: "Salohiddin", LastName: "Nuridinov", Email: "admin@gmail.com"},
		{FirstName: "Salohiddin", LastName: "Nuridinov", Email: "admin@gmail.com"},
		{FirstName: "Salohiddin", LastName: "Nuridinov", Email: "admin@gmail.com"},
		{FirstName: "Salohiddin", LastName: "Nuridinov", Email: "admin@gmail.com"},
		{FirstName: "Salohiddin", LastName: "Nuridinov", Email: "admin@gmail.com"},
		{FirstName: "Salohiddin", LastName: "Nuridinov", Email: "admin@gmail.com"},
	}

	db.Exec("TRUNCATE TABLE Users")

	for _, Users := range Users {
		db.Create(&Users)
	}
}
