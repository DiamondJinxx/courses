package main

import (
	"fmt"
)

func printIntArray(p *[3]*int){
	for i,v := range p{
		fmt.Println(i, "\t", v,"\t", *v)
	}
}


func main(){
	var (
		a int = 3
		b int = 4
		c int = 5
	)
	p := &[3]*int{&a, &b, &c}
	// var mass [5]int = [...]int{1,2,3,4,5}		
	fmt.Println(p[0:2])
	printIntArray(p)
	// printIntArray(mass)
}