package myutils

import (
	"log"
)

const (
	Red   = "\033[31m"
)

func ErrNotnil(err *error){
	if err!=nil{
		log.Fatalln(Red, "ERROR :", err)
	}
}



