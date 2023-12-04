package internal

type Network struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Logo   string `json:"logourl"`
}

type Lane struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type NetworkType string

const (
	Testnet NetworkType = "testnet"
	Mainnet NetworkType = "mainnet"
)
