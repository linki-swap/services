package lane

import (
	"context"

	"github.com/linki-swap/services/internal"
)

type laneService struct{}

func NewService() Service {
	return &laneService{}
}

func (ls *laneService) Networks(context.Context, internal.NetworkType) ([]internal.Network, error) {
	return []internal.Network{
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
	}, nil
}

func (ls *laneService) Assets(context.Context) {

}
