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
}
