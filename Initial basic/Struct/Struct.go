package main

import "fmt"

type Car struct {
	ID    int
	doors int
	tires int
	color string
	model int
	motor int
	name  string

}


func (car *Car) showName() string{
	return car.name
}

func (car *Car) setColor(color string) {
	car.color = color
}

func (car *Car) getColor() string {
	return car.color
}

func newCar (ID int,doors int,tires int,color string,model int,motor int,name string) *Car{
	return &Car{
		ID: ID,
		doors:doors,
		tires:tires,
		color:color,
		model:model,
		motor:motor,
		name:name,
	}
}

func main(){
   car := newCar(1,4,4,"black",2018,1600,"Ford Fusion")

   car.setColor("Red")

   fmt.Printf("%v",car.showName())



}