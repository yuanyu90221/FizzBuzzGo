# FizzBuzzGo

write a function fizzBuzz to generate a list of string based on the input number

## 我的解法

```golang
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
```