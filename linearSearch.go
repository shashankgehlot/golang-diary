package main

import (
	"fmt"
	"time"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	var first int
	var boolean bool
	fmt.Println("Enter the no you want to search: ")
	fmt.Scanln(&first)
	fmt.Println(first)
	t1 := time.Now()
	// fmt.Println(t1)
	for i := 0; i < len(primes); i++ {
		fmt.Println(primes[i])
	}

	for i := 0; i < len(primes); i++ {
		if primes[i] == first {
			fmt.Println("No found at position: ", i, primes[i])
			boolean = true
		}
	}
	if !boolean {
		fmt.Println(" does not exist")
	}
	t2 := time.Now()
	// fmt.Println(t2)
	diff := t2.Sub(t1)
	fmt.Println(diff)
}
