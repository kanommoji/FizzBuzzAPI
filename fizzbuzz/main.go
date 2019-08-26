package main

import (
	"fizzbuzz/fizzbuzz/api"
	"fizzbuzz/fizzbuzz/service"

	"github.com/gin-gonic/gin"
)

func main() {
	api := api.FizzBuzzAPI{
		FizzBuzz: &service.FizzBuzzStruct{},
	}
	route := gin.Default()
	route.POST("/fizzbuzz/convert_to_fizzbuzz", api.ConvertToFizzBuzzAPI)
	route.Run(":8080")
}
