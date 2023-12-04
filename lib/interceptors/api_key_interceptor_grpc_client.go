package interceptors

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type APIKeyClientInterceptor struct {
	apiKey string
}

func NewAPIKeyClientInterceptor(
	apiKey string,
) *APIKeyClientInterceptor {
	return &APIKeyClientInterceptor{
		apiKey: apiKey,
	}
}

func (interceptor *APIKeyClientInterceptor) clientUnaryInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		start := time.Now()
		log.Printf("Invoked RPC method=%s; Duration=%s", method,
			time.Since(start))

		return invoker(interceptor.attachAPIKey(ctx), method, req, reply, cc, opts...)
	}
}

func (interceptor *APIKeyClientInterceptor) attachAPIKey(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "x-api-key", interceptor.apiKey)
}

func (interceptor *APIKeyClientInterceptor) WithAPIKeyClientUnaryInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(interceptor.clientUnaryInterceptor())
}
