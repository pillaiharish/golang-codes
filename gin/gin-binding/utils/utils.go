package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Head struct {
	City string `json:"city"`
}

type Body struct {
	// json tag to de-serialize json body
	Name string `json:"name"`
}

func GetOnlyBody(c *gin.Context, body Body) error {
	// using BindJson method to serialize body with struct
	err := c.BindJSON(&body)
	if err != nil {
		return err
	}
	fmt.Println("Only body structure: ", body)
	return nil
}

func GetRequest(c *gin.Context, data map[string]interface{}) error {
	// using ShouldBindBodyWith which helps to de-serialize body to different structures
	err := c.ShouldBindBodyWith(&data, binding.JSON)
	if err != nil {
		return err
	}
	fmt.Println("Entire Request: ", data)
	return nil
}
