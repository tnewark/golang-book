package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains(s, substr string) bool
	fmt.Println(strings.Contains("test", "es"))
	// => true

	// func Count(s, sep string) int
	fmt.Println(strings.Count("test", "t"))
	// => 2

	// func HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix("test", "te"))
	// => true

	// func HasSuffix(s, suffix string) bool
	fmt.Println(strings.HasSuffix("test", "st"))
	// => true

	// func Index(s sep string) int
	fmt.Println(strings.Index("test", "e"))
	// => 1
}
