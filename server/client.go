package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"pbft/constz"
	"pbft/library/parallel"
	"pbft/message"
)

func StartClient() error {
	var (
		ctx         = context.Background()
		nodeClients []message.NodeServiceClient
	)

	conn, err := grpc.DialContext(ctx, fmt.Sprintf(constz.BaseEndpointFormat, constz.PrimaryNodePort), grpc.WithInsecure())
	if err != nil {
		return err
	}
	nodeClients = append(nodeClients, message.NewNodeServiceClient(conn))
	result, err := nodeClients[0].GetListOtherPort(ctx, &message.GetListOtherPortRequest{})
	if err != nil {
		return err
	}
	for _, port := range result.Ports {
		conn, err := grpc.DialContext(ctx, fmt.Sprintf(constz.BaseEndpointFormat, port), grpc.WithInsecure())
		if err != nil {
			return err
		}
		nodeClients = append(nodeClients, message.NewNodeServiceClient(conn))
	}

	p := parallel.New(ctx)
	for _, client := range nodeClients {
		func(client message.NodeServiceClient) {
			_ = p.Register(parallel.NewFunc(
				&message.RequestRequest{
					Data: &message.Data{
						PreviousBlockHash: "test",
						BlockHash:         "test",
						BlockHeight:       "test",
					},
				},
				func(ctx context.Context, request *message.RequestRequest) (*message.RequestResponse, error) {
					return client.Request(ctx, request)
				},
			))
		}(client)
	}
	return nil
}

type Client struct {
	nodeClients []message.NodeServiceClient
}
