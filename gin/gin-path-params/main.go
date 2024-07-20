package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type URI struct {
	Name string `json:"name" uri:"details"`
}

func main() {
	gin_new := gin.New()

	gin_new.GET("/test/:details", func(c *gin.Context) {

		uri := URI{}
		c.BindUri(&uri)
		fmt.Println(uri)
		c.JSON(http.StatusAccepted, &uri)
	})

	gin_new.Run(":3001")
}
