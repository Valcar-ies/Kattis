package main

import "fmt"

func main(){
	var n,p int
	var str string

	fmt.Scan(&n,&p)
	for i := 1; i <= n; i++{
		fmt.Scan(&str)
	}
	fmt.Print(p)
}