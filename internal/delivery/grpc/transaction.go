package grpc

import (
	"context"
	"fmt"

	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/grpc/transaction"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/customstatus"
	pb "github.com/dedihartono801/go-clean-architecture-v2/pkg/protobuf"
)

type Service struct {
	Service transaction.Service
}

func (h *Service) TransactionDetail(ctx context.Context, req *pb.TransactionId) (*pb.TransactionResponse, error) {
	fmt.Println(req.Id)
	trx, statuscode, err := h.Service.Find(req.Id)
	if err != nil {
		return &pb.TransactionResponse{
			Status: int32(statuscode),
			Error:  customstatus.ErrNotFound.Message,
		}, nil
	}

	data := &pb.FindOneData{
		Id:               trx.ID,
		AdminId:          trx.AdminID,
		TotalQuantity:    int32(trx.TotalQuantity),
		TotalTransaction: int32(trx.TotalTransaction),
	}
	return &pb.TransactionResponse{
		Status: int32(statuscode),
		Data:   data,
	}, nil
}
