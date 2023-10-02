package main 

import "fmt"

func main () {
    
    // x1 itu menandakan x output 1
    // Soal Matematika B Output 1
	var x1 float64 = 1.6 // ubah nilai di sini
	
	fmt.Println("Output 1: ")
	fmt.Scan(&x1)

	var fx1 = ((x1 * x1 * x1) + (3 * x1)) / ((x1 * x1 * x1 * x1) - (3 * x1 * x1) + 4)
	fmt.Println(fx1)
	
	// x2 itu menandakan x output 2
	// Soal Matematika B Output 2
	var x2 float64 = -1.4 // ubah nilai disini
	
	fmt.Println("Output 2: ")
	fmt.Scan(&x2)

	var fx2 = ((x2 * x2 * x2) + (3 * x2)) / ((x2 * x2 * x2 * x2) - (3 * x2 * x2) + 4)
	fmt.Println(fx2)
	
	// x3 itu menandakan x output 3
	// Soal Matematika B Output 3
	var x3 float64 = 0 // ubah nilai disini
	
	fmt.Println("Output 3: ")
	fmt.Scan(&x3)

	var fx3 = ((x3 * x3 * x3) + (3 * x3)) / ((x3 * x3 * x3 * x3) - (3 * x3 * x3) + 4)
	fmt.Println(fx3)
}