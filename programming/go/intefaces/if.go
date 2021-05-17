package main

import "fmt"

type Move interface {
	move()
}

type Car struct{
	name string
}

type Aircraft struct{
	name string
}

func drive(m Move){
	m.move()
}


func (c Car) move(){
	fmt.Println("Автомобиль ", c.name, " едет")
}

func (a Aircraft) move(){
	fmt.Println("Самолет ", a.name, "летит")
}

func main(){

	tesla := Car{"Tesla model 6"}
	var tesla2 Move = Car{"Tesla model 2"}
	boing := Aircraft{"Boing 777"}
	var mi2 Move = Aircraft{"Mi 2"}
	
	drive(tesla)
	drive(tesla2)
	drive(boing)
	drive(mi2)
}