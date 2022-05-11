package main

import "fmt"

func main(){
	var h,b int
	var hasil float64

	fmt.Scanln(&h,&b)
	hasil = float64(h)*float64(b) / 2

	fmt.Printf("%.7f",hasil)
}