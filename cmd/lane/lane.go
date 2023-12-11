package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	pb "github.com/linki-swap/services/api/v1/pb/lane"
	laneService "github.com/linki-swap/services/pkg/lane"
	"github.com/linki-swap/services/pkg/lane/endpoints"
	"github.com/linki-swap/services/pkg/lane/transport"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
)

const defaultGRPCPort = "8082"

func main() {
	var (
		logger   log.Logger
		grpcAddr = net.JoinHostPort("0.0.0.0", envString("HTTP_PORT", defaultGRPCPort))
	)

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	var service = laneService.NewService()
	service = laneService.LoggingMiddleware(logger)(service)
	eps := endpoints.NewEndpointSet(service)
	grpcServer := transport.NewGRPCServer(eps)

	var g group.Group
	{
		// The gRPC listener mounts the Go kit gRPC server we created.
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", grpcAddr)
			// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
			// the here demonstrated zipkin tracing middleware.
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			pb.RegisterLaneServer(baseServer, grpcServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
