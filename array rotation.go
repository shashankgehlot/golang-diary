package main

import "fmt"

func main() {
	var Testcase int
	var Arraylength int
	var k int
	fmt.Println("Enter the no of test Cases")
	fmt.Scanln(&Testcase)
	fmt.Println("no OF TEST case", Testcase)

	fmt.Println("Enter Array length and K (rotation var)")
	fmt.Scanln(&Arraylength)
	fmt.Scanln(&k)
	fmt.Println("Array length and K (rotation var)", Arraylength, k)

	a := make([]int, Arraylength)

	for i := 1; i < len(a); i++ {
		a[i] = i
	}
	// b := shift(a)
	for i := 0; i < k; i++ {
		a = shift(a)
	}
	fmt.Println(a)
}

func shift(arr []int) []int {
	tem := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = tem
	return arr
}
