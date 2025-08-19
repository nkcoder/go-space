package topics

import "fmt"

func ForLoop() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n % 2 != 0: ", n)
	}

}
