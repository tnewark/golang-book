package main

import "fmt"

func main() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i = 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, " even")
		} else {
			fmt.Println(i, " odd")
		}
	}

	for j := 1; j <= 100; j++ {
		if (j % 3 == 0) {
			fmt.Println(j)
		}
	}

	for j := 1; j <= 100; j++ {
		if (j % 15 == 0) {
			fmt.Println("fizzbuzz")
		} else if (j % 5 == 0) {
			fmt.Println("buzz")
		} else if ( j % 3 == 0) {
			fmt.Println("fizz")
		} else {
			fmt.Println(j)
		}
	}
}