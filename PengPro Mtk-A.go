package main 

import "fmt"

func main () { 
    
    // Soal Matematika A Output 1
	var x1 float64 = 2 // ubah angka di sini
	var y1 float64 = 2 // ubah angka di sini

	fmt.Print("Output 1: ")
	fmt.Scan(&x1, &y1)

	var fx1 = (1 / ((3 * x1 * x1) + 10 )) + ((10 * y1) + 7)
	fmt.Println(fx1)
	
	// Soal Matematika A Output 2
	
    var x2 float64 = 70 // ubah angka di sini
	var y2 float64 = 20 // ubah angka di sini

	fmt.Print("Output 2: ")
	fmt.Scan(&x2, &y2)

	var fx2 = (1 / ((3 * x2 * x2) + 10 )) + ((10 * y2) + 7)
	fmt.Println(fx2)
	
	// Soal Matematika A Output 3
	
    var x3 float64 = 5 // ubah angka di sini
	var y3 float64 = 46 // ubah angka di sini

	fmt.Print("Output 3: ")
	fmt.Scan(&x3, &y3)

	var fx3 = (1 / ((3 * x3 * x3) + 10 )) + ((10 * y3) + 7)
	fmt.Println(fx3)
    
} 
