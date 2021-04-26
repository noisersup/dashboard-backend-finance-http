package main

import (
	"log"

	"github.com/noisersup/dashboard-backend-finance-http/pb"
	"google.golang.org/grpc"
)

func main(){
	conn,err := grpc.Dial("localhost:9000",grpc.WithInsecure()) //TODO: make grpc secure
	if err != nil { log.Fatal(err) }
	
	defer func () {
		if err = conn.Close(); err != nil{
			log.Fatal(err)
		}
	}()
	
	client := pb.NewFinanceServiceClient(conn)
}