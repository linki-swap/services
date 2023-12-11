package apigateway

import (
	"context"

	"github.com/linki-swap/services/internal"
)

type apiGateway struct {
}

func NewService() Service {
	return &apiGateway{}
}

func (api *apiGateway) GetOrder(ctx context.Context, id string) {
}

func (api *apiGateway) OrderStatus(ctx context.Context, id string) {
}

func (api *apiGateway) CreateOrder(ctx context.Context, order internal.Order) {
}

func (api *apiGateway) CancelOrder(ctx context.Context, id string) {
}

func (api *apiGateway) OrderRoute(ctx context.Context) {
}

// Networks recieves a context and network type {mainnet or testnet} and returns a list of supported
// chains of the network type.
func (api *apiGateway) Networks(ctx context.Context, t internal.NetworkType) ([]internal.Network, error) {
	/* return []internal.Network{
		{
			Name:   "Ethereum",
			Symbol: "ETH",
			Logo:   "https://get.celer.app/cbridge-icons/ETH.png",
		},
		{
			Name:   "Ethereum",
			Symbol: "ETH",
			Logo:   "https://get.celer.app/cbridge-icons/ETH.png",
		},
		{
			Name:   "Ethereum",
			Symbol: "ETH",
			Logo:   "https://get.celer.app/cbridge-icons/ETH.png",
		},
	}, nil */
	return []internal.Network{}, nil
}

func (api *apiGateway) Assets(context.Context) {
}
