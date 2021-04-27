package logs

import "log"

func ErrorHandler(method string,err error){
	log.Printf("ERROR[%s()]: %s",method,err.Error())
}
func FatalErrorHandler(method string,err error){
	log.Fatalf("ERROR[%s()]: %s",method,err.Error())
}