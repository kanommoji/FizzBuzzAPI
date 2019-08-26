package api

import (
	"fizzbuzz/fizzbuzz/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FizzBuzzAPI struct {
	FizzBuzz service.FizzBuzzService
}

type FizzBuzz struct {
	number  int    `json:"number"`
	message string `json:"fizzbuzz"`
}

func (fizzBuzzAPI FizzBuzzAPI) ConvertToFizzBuzzAPI(context *gin.Context) {
	var fizzBuzz FizzBuzz
	err := context.BindJSON(&fizzBuzz)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fizzBuzz.message = fizzBuzzAPI.FizzBuzz.ConvertToFizzBuzz(fizzBuzz.number)
	context.String(http.StatusOK, fizzBuzz.message)
}
