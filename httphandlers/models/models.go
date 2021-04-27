package models

import "github.com/noisersup/dashboard-backend-finance-http/pb"


type GetResponse struct{
	Groups	[]pb.Groups `json:"groups"`
	Error	string `json:"error"`
}

type ErrorResponse struct{
	Error	string `json:"error"`
}