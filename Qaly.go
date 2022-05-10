package main

import "fmt"

func main(){
	var n int
	var hasil,q,y,count float64

	fmt.Scan(&n)
	for i := 1; i <= n; i++{
		fmt.Scan(&q,&y)
		hasil = q * y
		count += hasil
	}
	fmt.Printf("%.3f",count)
}