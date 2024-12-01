package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func UnaryLogErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {
		res, err = handler(ctx, req)
		if err != nil {
			fmt.Print(err)
		}
		return res, err
	}
}
