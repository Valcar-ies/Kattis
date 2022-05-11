package main

import "fmt"

func main(){
	var hasil string
	var n int

	fmt.Scan(&n)
	if n % 2 == 0{
		hasil = "Bob"
	}else{
		hasil = "Alice"
	}
	fmt.Print(hasil)
}