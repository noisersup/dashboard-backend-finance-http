package logs

import "log"

func ErrorHandler(method string,err error){
	log.Printf("ERROR[%s()]: %s",method,err.Error())
}