package fizz_buzz

import "fmt"

func FizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i <= n; i++ {
		result = append(result, mapResponse(i))
	}
	return result
}

func mapResponse(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%d", n)
}
