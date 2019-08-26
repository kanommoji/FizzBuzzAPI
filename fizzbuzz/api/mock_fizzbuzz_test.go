package api_test

type MockService struct {
}

func (mockService MockService) ConvertToFizzBuzz(number int) string {
	return "Fizz"
}
