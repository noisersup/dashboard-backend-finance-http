package httphandlers

import (
	"context"
	"net/http"
	"time"

	"github.com/noisersup/dashboard-backend-finance-http/httphandlers/models"
	"github.com/noisersup/dashboard-backend-finance-http/httphandlers/utils"
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

	response := models.GetResponse{}

	groups, err := h.grpcClient.GetGroups(ctx,&emptypb.Empty{})
	if err != nil{
		logs.ErrorHandler(method,err)
		response.Error = err.Error()
		utils.SendResponse(w,response,http.StatusInternalServerError)
		return
	}

	for _,pointer := range groups.Groups {
		response.Groups = append(response.Groups, *pointer)
	}

	utils.SendResponse(w,response,http.StatusOK)
}