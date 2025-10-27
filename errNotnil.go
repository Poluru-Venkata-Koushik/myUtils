package myUtils

import (
	"log"
)

const (
	Red   = "\033[31m"
)

func ErrNotNil(err error){
	if err!=nil{
		log.Fatalln(Red, "ERROR :", err)
	}
}



