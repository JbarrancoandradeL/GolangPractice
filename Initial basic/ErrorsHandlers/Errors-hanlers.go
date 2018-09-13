package main

import (
	"github.com/pkg/errors"
	"fmt"
)

func isEnable(enable bool) (bool,error){
	if enable {
		return false,errors.New("la variable es un booleano")
	}

	return true,nil
}



func main (){
	i, err := isEnable(true)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(i)
	}

}
