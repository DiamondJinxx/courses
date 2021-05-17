package main

import "fmt"



func op(x,y int, operation func(int, int) int) int{
	return operation(x,y)
}

func action(x int) func(int,int)int{
	switch(x){
		case 1:
			 return func(x,y int) (val int) {val = x + y; return}
		case 2: 
			return func(x,y int) (val int) {val = x - y; return}
		case 3: 
			return func(x,y int) (val int) {val = x * y; return}
	}
	// fmt.Println("Дефолтное")
	return func(x, y int)int{return x/y}
}


func square() func() int{
	x := 2
	return func() int{
		x ++ 
		return x * x
	}
}


func main(){
	// var sum func(x,y float32)float32 = func(x,y float32) float32{return x + y}

	fmt.Println(op(3,4, action(1)))
	fmt.Println(op(3,4, action(2)))
	fmt.Println(op(3,4, action(3)))
	fmt.Println(op(3,4, action(0)))

	fmt.Println("А теперь квадрат!")
	var f func() int = square()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())


}
	