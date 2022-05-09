package main

import "fmt"

func main(){
	var R1,S int
	var R2 float64

	fmt.Scan(&R1,&S)
	R2 = (float64(S)-float64(R1)/2) * 2
	fmt.Print(int(R2))
}