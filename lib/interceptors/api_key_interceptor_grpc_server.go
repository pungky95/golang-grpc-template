package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type APIKeyServerInterceptor struct {
	apiKey string
}

func NewAPIKeyServerInterceptor(
	apiKey string,
) *APIKeyServerInterceptor {
	return &APIKeyServerInterceptor{
		apiKey: apiKey,
	}
}

func (interceptor *APIKeyServerInterceptor) serverUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "invalid metadata")
	}

	var xAPIKey []string
	if xAPIKey = md["x-api-key"]; len(xAPIKey) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "missing credentials")
	}

	if xAPIKey[0] != interceptor.apiKey {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	return handler(ctx, req)
}

func (interceptor *APIKeyServerInterceptor) WithAPIKeyServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(interceptor.serverUnaryInterceptor)
}
