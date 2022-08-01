package main

import "fmt"

func main(){
	var n int
	var kalimat string

	fmt.Scan(&n)
	for i := 1; i <= n; i++{
		fmt.Scan(&kalimat)
		if i % 2 == 1{
			fmt.Println(kalimat)
		}
	}
}