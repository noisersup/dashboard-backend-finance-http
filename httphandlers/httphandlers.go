package httphandlers

import (
	"context"
	"net/http"
	"time"

	hu "github.com/noisersup/dashboard-backend-finance-http/httphandlers/utils"
	"github.com/noisersup/dashboard-backend-finance-http/logs"
	"github.com/noisersup/dashboard-backend-finance-http/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HttpHandlers struct {
	grpcClient	pb.FinanceServiceClient
} 

func CreateHandlers(financeClient pb.FinanceServiceClient) HttpHandlers{
	return HttpHandlers{financeClient}
}

func (h *HttpHandlers) GetGroups(w http.ResponseWriter, r *http.Request){
	var method string = "GetGroups"
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))
	defer cancel()

	groups, err := h.grpcClient.GetGroups(ctx,&emptypb.Empty{})
	if err != nil{
		logs.ErrorHandler(method,err)
		return
	}

	hu.SendResponse(w,groups,http.StatusOK)
}