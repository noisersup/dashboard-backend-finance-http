package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/noisersup/dashboard-backend-finance-http/httphandlers"
	"github.com/noisersup/dashboard-backend-finance-http/pb"
	"google.golang.org/grpc"
)

type ServiceConfig struct{
	GRPCAddress	string
	GRPCPort	int
	
	HttpPort	int
}

func main(){
	config := getVars()

	grpcPayload := fmt.Sprintf("%s:%d",config.GRPCAddress,config.GRPCPort)
	log.Printf("Connecting to %s", grpcPayload)

	conn,err := grpc.Dial(grpcPayload,grpc.WithInsecure()) //TODO: make grpc secure
	if err != nil { log.Fatal(err) }
	log.Print("gRPC Connection initialized")
	defer func () {
		if err = conn.Close(); err != nil{
			log.Fatal(err)
		}
	}()
	
	client := pb.NewFinanceServiceClient(conn)
	h := httphandlers.CreateHandlers(client)
	r := mux.NewRouter()

	r.HandleFunc("/finances",h.GetGroups).Methods("GET")

	log.Printf("Initializing http server on :%d",config.HttpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%d",config.HttpPort),r))
}


func getVars() *ServiceConfig {
	var config ServiceConfig

	config.GRPCAddress = os.Getenv("GRPC_ADDRESS")

	config.GRPCPort,_ = strconv.Atoi(os.Getenv("GRPC_PORT")) 
	config.HttpPort,_ = strconv.Atoi(os.Getenv("HTTP_PORT")) 

	if config.GRPCAddress=="" || config.GRPCPort==0 || config.HttpPort==0 { 
		log.Fatal("ENV variables did not set")
	}
	return &config
}