package transport

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/linki-swap/services/internal/util"
	"github.com/linki-swap/services/pkg/apigateway/endpoints"
)

func NewHTTPHandler(ep endpoints.Set) *http.ServeMux {
	m := http.NewServeMux()
	m.Handle("/getnetworks", httptransport.NewServer(
		ep.NetworksEndpoint,
		decodeNetworksRequest,
		encodeResponse,
	))

	return m
}

func decodeNetworksRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetNetworksRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case util.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case util.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
