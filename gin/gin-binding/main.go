package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	// json tag to de-serialize json body
	Name string `json:"name"`
}

func main() {
	gin_new := gin.New()

	gin_new.POST("/test", func(c *gin.Context) {
		body := Body{}
		// using BindJson method to serialize body with struct
		err := c.BindJSON(&body)

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body)
		c.JSON(http.StatusAccepted, &body)
	})

	gin_new.Run(":3000")
}
