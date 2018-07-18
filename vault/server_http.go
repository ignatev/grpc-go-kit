package vault

import (
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	m := http.NewServerMux()
	m.Handle("/hash", httptransport.Newserver(
		ctx,
		endpoints.HashEndpoint,
		decodeHashRequest,
		encodeResponse,
	))
	m.Handle("/validate", httptransport.NewServer(
		ctx,
		endpoints.ValidateEndpoint,
		decodeValidateRequest,
		encodeResponse,
	))
	return m
}