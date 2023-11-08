package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	InitDB()

	r := gin.Default()

	r.GET("/api/v1/users", func(c *gin.Context) {
		users, err := GetAllUsers()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, users)
	})

	r.GET("/api/v1/users/:id", func(c *gin.Context) {

		idStr := c.Param("id")
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		user, err := GetUserByID(uint(idInt))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	})

	r.POST("/api/v1/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := CreateUser(&user); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	})

	r.PUT("/api/v1/users/:id", func(c *gin.Context) {
		// URL orqali foydalanuvchi identifikatorini olish
		userID := c.Param("id")

		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(400, gin.H{"error": "Noto'g'ri foydalanuvchi identifikatori"})
			return
		}

		// Bazadan foydalanuvchi obyektini topish
		existingUser, err := GetUserByID(uint(userIDInt))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// JSON ma'lumotlarni yangi ma'lumotlar bilan qo'shish
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Yangi ma'lumotlarni mavjud ma'lumotlarga yangilash
		existingUser.LastName = newUser.LastName
		existingUser.FirstName = newUser.FirstName
		existingUser.Email = newUser.Email

		// Yangi ma'lumotlarni bazaga saqlash
		if err := UpdateUser(existingUser); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, existingUser)
	})

	r.DELETE("/api/v1/users/:id", func(c *gin.Context) {

		idStr := c.Param("id")
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}

		if err := DeleteUser(uint(idInt)); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "User deleted successfully"})
	})

	r.Run()

}
