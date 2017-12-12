package main

import "fmt"

func main() {
	{
		var x [5]int
		x[4] = 100
		fmt.Println(x)
	}

	{
		x := [5]float64{
			98, 
			93, 
			77, 
			82, 
			83, 
		}
		var total float64 
		for _, value := range x {
			total += value
		}
		fmt.Println(total / float64(len(x)))

		slice1 := []int{1,2,3}
		slice2 := append(slice1, 4, 5)
		fmt.Println(slice1, slice2)

		slice3 := []int{1,2,3}
		slice4 := make([]int,2)
		copy(slice4, slice3)
		fmt.Println(slice3, slice4)

		xx := make(map[string]int)
		xx["key"] = 10
		fmt.Println(xx["key"])

		elements := map[string]string {
			"H": "Hydrogen",
			"He": "Helium",
			"Li": "Lithium",
			"Be": "Beryllium",
			"B": "Boron",
			"C": "Carbon",
			"N": "Nitrogen",
			"O": "Oxygen",
			"F": "Fluorine",
			"Ne": "Neon",
		}
		fmt.Println(elements["Li"])

		if name, ok := elements["Un"]; ok {
			fmt.Println(name, ok)
		}

		nums := []int {
			48,96,86,68,
			57,82,63,70,
			37,34,83,27,
			19,97, 9,17,
		}

		maxElem := 0
		for _, val := range nums {
			if val > maxElem {
				maxElem = val
			}
		}
		fmt.Println("Max element is: ", maxElem)
	}
}