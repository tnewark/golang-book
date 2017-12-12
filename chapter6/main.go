package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _,v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func sum( nums []int) int {
	total := 0
	for _,v := range nums {
		total += v
	}
	return total
}

func add(args ...int) int {
	total := 0
	for _,v := range args {
		total += v
	}
	return total
}

func max(args ...int) int {
	maxNum := args[0]
	for _,v := range args {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}

func half(num int) (int, bool) {
	return num /2 , (num  % 2 == 0)
}


func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return 
	}
}

func fib(n uint) uint {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1 
	}

	return fib(n-1) + fib(n-2)
}

func square(x *float64) {
	*x = *x * *x
}

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	ys := []int{98, 93, 77, 82, 83}
	fmt.Println("avg: ",average(xs))
	fmt.Println("max: ", max(ys...))
	fmt.Println(add(1,2,3,4))
	for _, v := range ys {
		fmt.Println(half(v))
	}

	nextOdd := makeOddGenerator() 

	for i:=0; i< 10 ; i++ {
		fmt.Println(nextOdd())
	}

	fmt.Println("10th fib number is ", fib(10))

	sVal := new(float64)
	*sVal = 1.5
	square(sVal)

	fmt.Println("Square of 1.5 is ", *sVal)
	a :=1 
	b := 2
	swap(&a, &b)
	fmt.Println("Swap of 1 and 2 : ", a, b)
}

