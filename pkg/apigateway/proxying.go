package apigateway

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	lanepb "github.com/linki-swap/services/api/v1/pb/lane"
	"github.com/linki-swap/services/internal"
	"github.com/linki-swap/services/pkg/lane/endpoints"
	"google.golang.org/grpc"
)

type proxymw struct {
	Service
	networksEndpoint endpoint.Endpoint
}

func ProxyingMiddleware(ctx context.Context, instance string, logger log.Logger) ServiceMiddleware {
	// If instances is empty, don't proxy.
	if instance == "" {
		logger.Log("proxy_to", "none")
		return func(next Service) Service { return next }
	}

	conn, err := grpc.Dial(instance, grpc.WithInsecure())
	if err != nil {
		logger.Log("GRPC connection error", err)
		return func(next Service) Service { return next }
	}

	// Each individual endpoint is an grpc/transport.Client (which implements
	// endpoint.Endpoint) that gets wrapped with various middlewares. If you
	// made your own client library, you'd do this work there, so your server
	// could rely on a consistent set of client behavior.
	var networksEndpoint endpoint.Endpoint
	{
		networksEndpoint = grpctransport.NewClient(
			conn,
			"pb.Lane",
			"Networks",
			encodeGRPCNetworkRequest,
			decodeGRPCNetworkResponse,
			lanepb.NetworksReply{},
		).Endpoint()
	}

	// And finally, return the ServiceMiddleware, implemented by proxymw.
	return func(next Service) Service {
		return proxymw{next, networksEndpoint}
	}
}

func encodeGRPCNetworkRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(internal.NetworkType)
	return &lanepb.NetworksRequest{Networktype: string(req)}, nil
}

func decodeGRPCNetworkResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*lanepb.NetworksReply)
	nets := []internal.Network{}
	for _, v := range reply.Networks {
		nets = append(nets, internal.Network{
			Name:   v.Name,
			Logo:   v.Logo,
			Symbol: v.Symbol,
		})
	}
	return endpoints.GetNetworksResponse{Networks: nets, Err: reply.Err}, nil
}

func (pmw proxymw) Networks(ctx context.Context, netType internal.NetworkType) ([]internal.Network, error) {
	resp, err := pmw.networksEndpoint(ctx, netType)
	if err != nil {
		return nil, err
	}
	response := resp.(endpoints.GetNetworksResponse)
	return response.Networks, errors.New(response.Err)
}
