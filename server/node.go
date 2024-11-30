package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"pbft/constz"
	"pbft/library/parallel"
	"pbft/library/slice"
	"pbft/message"
	"sync"
)

func StartNode(port int32, isByzantine bool) error {
	var (
		ctx            = context.Background()
		address        = fmt.Sprintf(constz.BaseEndpointFormat, port)
		nodeId         int32
		isPrimaryNode  = port == constz.PrimaryNodePort
		otherNodePorts = make([]int32, 0)
		nodeClients    = make(map[int32]message.NodeServiceClient)
	)
	node := &Node{}
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	message.RegisterNodeServiceServer(grpcServer, node)

	reflection.Register(grpcServer)

	done := make(chan error)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			done <- err
		}
	}()
	if isPrimaryNode {
		isByzantine = false
		nodeId = 0
	} else {
		conn, err := grpc.DialContext(ctx, fmt.Sprintf(constz.BaseEndpointFormat, constz.PrimaryNodePort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		nodeClients[constz.PrimaryNodePort] = message.NewNodeServiceClient(conn)
		result, err := nodeClients[constz.PrimaryNodePort].AddNode(ctx, &message.AddNodeRequest{Port: port})
		if err != nil {
			return err
		}
		fmt.Printf("Node Id is %v\n", result.GetNodeId())

		otherNodePorts = append(otherNodePorts, constz.PrimaryNodePort)
		nodeId = result.NodeId
		for _, p := range result.OtherNodePorts {
			otherNodePorts = append(otherNodePorts, p)
			conn, _ = grpc.DialContext(ctx, fmt.Sprintf(constz.BaseEndpointFormat, p), grpc.WithInsecure())
			nodeClients[p] = message.NewNodeServiceClient(conn)
		}
	}

	node.nodeId = nodeId
	node.port = port
	node.isByzantine = isByzantine
	node.otherNodePorts = otherNodePorts
	node.nodeClients = nodeClients

	err = <-done
	if err != nil {
		return err
	}
	return nil
}

type Node struct {
	message.UnimplementedNodeServiceServer
	nodeId         int32
	port           int32
	isByzantine    bool
	otherNodePorts []int32
	nodeClients    map[int32]message.NodeServiceClient
	mutex          sync.Mutex
}

func (n *Node) PrePrepare(ctx context.Context, req *message.PrePrepareRequest) (*message.PrePrepareResponse, error) {
	return &message.PrePrepareResponse{}, nil
}

func (n *Node) Prepare(ctx context.Context, req *message.PrepareRequest) (*message.PrepareResponse, error) {
	return &message.PrepareResponse{}, nil
}

func (n *Node) Commit(ctx context.Context, req *message.CommitRequest) (*message.CommitResponse, error) {
	return &message.CommitResponse{}, nil
}

func (n *Node) Request(ctx context.Context, req *message.RequestRequest) (*message.RequestResponse, error) {
	return &message.RequestResponse{}, nil
}

func (n *Node) AddNode(ctx context.Context, req *message.AddNodeRequest) (*message.AddNodeResponse, error) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	clones := slice.Make(n.otherNodePorts...)
	n.otherNodePorts = append(n.otherNodePorts, req.Port)
	conn, err := grpc.DialContext(context.Background(), fmt.Sprintf(constz.BaseEndpointFormat, req.Port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	n.nodeClients[req.Port] = message.NewNodeServiceClient(conn)

	if n.nodeId == constz.PrimaryNodeId {
		p := parallel.New(ctx)
		for _, port := range clones {
			func(port int32) {
				_ = p.Register(parallel.NewFunc(
					&message.AddNodeRequest{Port: req.Port},
					func(ctx context.Context, request *message.AddNodeRequest) (*message.AddNodeResponse, error) {
						return n.nodeClients[port].AddNode(ctx, request)
					},
				))
			}(port)
		}
		if err = p.Wait(); err != nil {
			return nil, err
		}

	} else {
		clones = make([]int32, 0)
	}

	fmt.Printf("Add node: %v successfully\n", req.Port)
	return &message.AddNodeResponse{
		NodeId:         int32(len(clones) + 1),
		OtherNodePorts: clones,
	}, nil
}
