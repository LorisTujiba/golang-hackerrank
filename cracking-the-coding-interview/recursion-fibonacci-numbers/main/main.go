package main
import "fmt"

func fibonacci(n int) int {
	var lh = 0
	var rh = 1
	var result int

	for i := 0; i<n;i++{
		result = lh + rh
		lh = rh
		rh = result
	}
	return lh
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Println(fibonacci(n))
}
