package gapi

import (
	"github.com/markitos-es/markitos-svc-boilerplates-grpc/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewGRPCBoilerplate(in *domain.Boilerplate) *Boilerplate {
	return &Boilerplate{
		Id:        in.Id,
		Name:      in.Name,
		CreatedAt: timestamppb.New(in.CreatedAt),
		UpdatedAt: timestamppb.New(in.UpdatedAt),
	}
}
