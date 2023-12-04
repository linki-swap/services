package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/linki-swap/services/pkg/apigateway"
)

type Set struct {
	NetworksEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc apigateway.Service) Set {
	return Set{
		NetworksEndpoint: MakeNetworksEndpoint(svc),
	}
}

func MakeNetworksEndpoint(svc apigateway.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetNetworksRequest)
		networks, err := svc.Networks(ctx, req.NetworkType)
		if err != nil {
			return GetNetworksResponse{networks, err.Error()}, nil
		}
		return GetNetworksResponse{networks, ""}, nil
	}
}
