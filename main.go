package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/noisersup/dashboard-backend-finance-http/httphandlers"
	"github.com/noisersup/dashboard-backend-finance-http/pb"
	"google.golang.org/grpc"
)

func main(){
	conn,err := grpc.Dial("localhost:9000",grpc.WithInsecure()) //TODO: make grpc secure
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

	log.Print("Initializing http server")

	log.Fatal(http.ListenAndServe(":8000",r))
}