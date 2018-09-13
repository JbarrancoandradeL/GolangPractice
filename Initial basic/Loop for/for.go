package main

import (
	"strconv"
	"fmt"
)

func main(){
	var x string
	x = ""
	for i:= 10; i>=1;i--{

		if i==10 {
           x = strconv.Itoa(i)

		}else {
			x =  x + "," + strconv.Itoa(i)
		}

	}
	fmt.Print(x)
}
