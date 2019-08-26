package service

import "strconv"

type FizzBuzzService interface {
	ConvertToFizzBuzz(int) string
}

type FizzBuzzStruct struct {
}

func (fizzBuzz FizzBuzzStruct) ConvertToFizzBuzz(number int) string {
	var fizzBuzzString string
	mapFizzBuzz := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}
	for key, value := range mapFizzBuzz {
		if number%key == 0 {
			fizzBuzzString += value
		}
	}
	if fizzBuzzString != "" {
		return fizzBuzzString
	}
	return strconv.Itoa(number)
}
