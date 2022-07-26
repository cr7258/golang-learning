package service

import (
	"context"
	"grpcpro/src/pbfiles"
)

type ProdService struct {
	pbfiles.UnimplementedProdServiceServer
}

func (*ProdService) GetProd(ctx context.Context, req *pbfiles.ProdRequest) (*pbfiles.ProdResponse, error) {
	rsp := &pbfiles.ProdResponse{
		Result: &pbfiles.ProdModel{Id: req.ProdId, Name: "chengzw"},
	}
	return rsp, nil
}

func NewProdService() *ProdService {
	return &ProdService{}
}
