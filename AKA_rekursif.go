package main

import "fmt"

func Rekursif(n int) string {
	if n == 0 {
		return "0"
	}
	if n == 1 {
		return "1"
	}
	return Rekursif(n/2) + string('0'+n%2) // Rekursi + tambah digit
}

func main() {
	fmt.Println("Rekursif:")
	fmt.Println("5 ->", Rekursif(5))   // Output: 101
	fmt.Println("10 ->", Rekursif(10)) // Output: 1010
	fmt.Println("0 ->", Rekursif(0))   // Output: 0
}