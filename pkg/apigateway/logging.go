package apigateway

import (
	"context"
	"time"

	"github.com/go-kit/log"
	"github.com/linki-swap/services/internal"
)

func LoggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next Service) Service {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	Service
}

func (mw logmw) Networks(ctx context.Context, t internal.NetworkType) (output []internal.Network, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "networks",
			"input", t,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Service.Networks(ctx, t)
	return
}
