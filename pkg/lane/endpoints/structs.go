package endpoints

import "github.com/linki-swap/services/internal"

type GetNetworksRequest struct {
	NetworkType internal.NetworkType `json:"networktype"`
}

type GetNetworksResponse struct {
	Networks []internal.Network `json:"networks"`
	Err      string             `json:"err,omitempty"`
}
