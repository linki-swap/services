package lane

import (
	"context"

	"github.com/linki-swap/services/internal"
)

type Service interface {
	Networks(context.Context, internal.NetworkType) ([]internal.Network, error)
	Assets(context.Context)
}

type ServiceMiddleware func(Service) Service
