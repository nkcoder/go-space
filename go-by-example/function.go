package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// If the types of the parameters are the same, you can omit the type from all but the last parameter
func sum(a, b, c int) int {
	return a + b + c
}

// Multiple return values
func vals() (int, int) {
	return 4, 5
}

// Variadic functions can be called with any number of trailing arguments
func variadicSum(numbers ...int) int {
	fmt.Println("numbers:", numbers)
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Anonymous functions are functions that are defined without a name, and they can be used to create closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func Function() {
	sum1 := plus(3, 4)
	sum2 := sum(4, 5, 7)
	fmt.Println("sum1:", sum1, "sum2:", sum2)

	x, y := vals()
	fmt.Println("x:", x, "y:", y)

	// The blank identifier _ can be used to ignore a value
	_, z := vals()
	fmt.Println("z:", z)

	sum3 := variadicSum(1, 2, 3, 4, 5)
	fmt.Println("sum3:", sum3)

	// If you already have a slice, you can pass it to a variadic function using the ... operator
	nums := []int{2, 4, 6, 8, 10}
	sum4 := variadicSum(nums...)
	fmt.Println("sum4:", sum4)

	nextInt := intSeq()
	fmt.Println("nextInt:", nextInt())
	fmt.Println("nextInt:", nextInt())
	fmt.Println("nextInt:", nextInt())

	// The state in the closure is unique to each function
	nextInt2 := intSeq()
	fmt.Println("nextInt2:", nextInt2())
	fmt.Println("nextInt2:", nextInt2())

	fmt.Println(factorial(8))

	// Anonymous functions can also be recursive, but this requires explicitly declaring a variable with `var` to store the function before it's defined
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(8))
}
