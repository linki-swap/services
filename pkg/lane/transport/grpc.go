package transport

import (
	"context"

	lanepb "github.com/linki-swap/services/api/v1/pb/lane"

	"github.com/go-kit/kit/transport/grpc"
	"github.com/linki-swap/services/internal"
	"github.com/linki-swap/services/pkg/lane/endpoints"
)

type grpcServer struct {
	networks grpc.Handler
	lanepb.UnimplementedLaneServer
}

func NewGRPCServer(ep endpoints.Set) lanepb.LaneServer {
	return &grpcServer{
		networks: grpc.NewServer(ep.NetworksEndpoint,
			decodeGRPCNetworksRequest,
			encodeGRPCNetworksResponse,
		),
	}
}

func (g *grpcServer) Networks(ctx context.Context, r *lanepb.NetworksRequest) (*lanepb.NetworksReply, error) {
	_, rep, err := g.networks.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*lanepb.NetworksReply), nil
}

func decodeGRPCNetworksRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*lanepb.NetworksRequest)
	// perform network activities.
	return endpoints.GetNetworksRequest{NetworkType: internal.NetworkType(req.Networktype)}, nil
}

func encodeGRPCNetworksResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.GetNetworksResponse)
	nets := []*lanepb.Network{}
	for _, v := range reply.Networks {
		nets = append(nets, &lanepb.Network{
			Name:   v.Name,
			Logo:   v.Logo,
			Symbol: v.Symbol,
		})
	}
	return &lanepb.NetworksReply{Networks: nets, Err: reply.Err}, nil
}
