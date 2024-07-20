package main

import (
	"gin-binding/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin_new := gin.New()

	gin_new.POST("/body", func(c *gin.Context) {
		body := utils.Body{}
		err := utils.GetOnlyBody(c, body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body for Body"})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"message": "Successful post to body"})
	})
	gin_new.POST("/data", func(c *gin.Context) {
		var data map[string]interface{}
		err := utils.GetRequest(c, data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body for Data"})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"message": "Successful post to data"})
	})

	gin_new.Run(":3000")
}
