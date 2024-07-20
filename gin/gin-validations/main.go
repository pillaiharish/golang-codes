package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID          uint   `json:"id" uri:"details"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Phone       string `json:"phone" binding:"required,e164"`
	CountryCode string `json:"countryCode" binding:"required,iso3166_1_alpha2"`
}

func main() {
	ginNew := gin.New()

	ginNew.POST("/user", func(c *gin.Context) {
		user := User{}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})

	})

	ginNew.Run(":3002")
}
