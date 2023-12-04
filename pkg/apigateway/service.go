package apigateway

import (
	"context"

	"github.com/linki-swap/services/internal"
)

type Service interface {
	GetOrder(context.Context, string)
	OrderStatus(context.Context, string)
	CreateOrder(context.Context, internal.Order)
	CancelOrder(context.Context, string)
	OrderRoute(context.Context)

	Networks(context.Context, internal.NetworkType) ([]internal.Network, error)
	Assets(context.Context)
}

// ServiceMiddleware is a chainable behavior modifier for the apigateway service.
type ServiceMiddleware func(Service) Service
