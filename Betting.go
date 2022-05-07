package main

import "fmt"

func main(){
	var n int
	var hasil1, hasil2 float64

	fmt.Scanln(&n)
	hasil1 = 100 / float64(n)
	hasil2 = 100 / (100-float64(n))
	fmt.Printf("%.3f\n",hasil1)
	fmt.Printf("%.3f",hasil2)

}