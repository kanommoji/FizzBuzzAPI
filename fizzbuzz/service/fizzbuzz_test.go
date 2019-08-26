package service_test

import (
	"fizzbuzz/fizzbuzz/service"
	"testing"
)

func Test_ConvertToFizzBuzz_Input_2_Should_Be_2(t *testing.T) {
	expected := "2"
	number := 2
	fizzBuzz := service.FizzBuzzStruct{}

	actual := fizzBuzz.ConvertToFizzBuzz(number)

	if expected != string(actual) {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_ConvertToFizzBuzz_Input_3_Should_Be_Fizz(t *testing.T) {
	expected := "Fizz"
	number := 3
	fizzBuzz := service.FizzBuzzStruct{}

	actual := fizzBuzz.ConvertToFizzBuzz(number)

	if expected != string(actual) {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_ConvertToFizzBuzz_Input_5_Should_Be_Buzz(t *testing.T) {
	expected := "Buzz"
	number := 5
	fizzBuzz := service.FizzBuzzStruct{}

	actual := fizzBuzz.ConvertToFizzBuzz(number)

	if expected != string(actual) {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_ConvertToFizzBuzz_Input_6_Should_Be_Fizz(t *testing.T) {
	expected := "Fizz"
	number := 6
	fizzBuzz := service.FizzBuzzStruct{}

	actual := fizzBuzz.ConvertToFizzBuzz(number)

	if expected != string(actual) {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_ConvertToFizzBuzz_Input_15_Should_Be_FizzBuzz(t *testing.T) {
	expected := "FizzBuzz"
	number := 15
	fizzBuzz := service.FizzBuzzStruct{}

	actual := fizzBuzz.ConvertToFizzBuzz(number)

	if expected != string(actual) {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}
