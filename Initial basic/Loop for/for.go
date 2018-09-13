package main

import (
	"fmt"
	"strconv"
)

func main(){
	var x string
	x = ""
	for i:= 1; i<=10;i++{

		if i==10 {
           x = strconv.Itoa(i)

		}else {
			x =  x + "," + strconv.Itoa(i)
		}

	}
	fmt.Print(x)
}
