package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"math/rand"
	"net"
	"pbft/constz"
	"pbft/interceptor"
	"pbft/library/hash"
	"pbft/library/math"
	"pbft/library/parallel"
	"pbft/library/random"
	"pbft/library/slice"
	"pbft/message"
	"sync"
)

func StartNode(port int32, isByzantine bool) error {
	var (
		ctx            = context.Background()
		address        = fmt.Sprintf(constz.BaseEndpointFormat, port)
		nodeId         int32
		otherNodePorts = make([]int32, 0)
		nodeClients    = make(map[int32]message.NodeServiceClient)
		prePrepareLog  = make(map[int32]*PrePrepareLog)
		prepareLog     = make(map[int32]*LogEntry)
		commitLog      = make(map[int32]*LogEntry)
	)
	node := &Node{port: port}
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.UnaryLogErrorInterceptor()),
	)
	message.RegisterNodeServiceServer(grpcServer, node)

	reflection.Register(grpcServer)

	done := make(chan error)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			done <- err
		}
	}()
	if node.isPrimaryNode() {
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
	node.prePrepareLog = prePrepareLog
	node.prepareLog = prepareLog
	node.commitLog = commitLog

	if err = <-done; err != nil {
		return err
	}
	return nil
}

type PrePrepareLog struct {
	digest string
	data   *message.Data
}
type LogEntry struct {
	count     int32
	processed bool
}
type Node struct {
	message.UnimplementedNodeServiceServer
	nodeId         int32
	port           int32
	isByzantine    bool
	otherNodePorts []int32
	sequenceId     int32
	nodeClients    map[int32]message.NodeServiceClient
	prePrepareLog  map[int32]*PrePrepareLog
	prepareLog     map[int32]*LogEntry
	commitLog      map[int32]*LogEntry
	mutex          sync.Mutex
}

func (n *Node) Request(ctx context.Context, req *message.RequestRequest) (*message.RequestResponse, error) {
	fmt.Printf("Node %v received request from client with content: \n", n.nodeId)
	fmt.Println(req)
	if n.isPrimaryNode() {
		digest, err := hash.Hash(req.Data)
		if err != nil {
			return nil, err
		}
		prePrepareMessage := &message.PrePrepareRequest{
			ViewId:     0,
			Digest:     digest,
			SequenceId: n.increaseSequenceId(),
			Data:       req.Data,
		}

		p := parallel.New(ctx)
		for _, client := range n.nodeClients {
			func(client message.NodeServiceClient) {
				_ = p.Register(parallel.NewFunc(
					prePrepareMessage,
					func(ctx context.Context, request *message.PrePrepareRequest) (*message.PrePrepareResponse, error) {
						return client.PrePrepare(ctx, request)
					},
				))
			}(client)
		}
	}
	return &message.RequestResponse{}, nil
}

func (n *Node) PrePrepare(ctx context.Context, req *message.PrePrepareRequest) (*message.PrePrepareResponse, error) {
	fmt.Printf("Node %v received preprepare request with content: \n", n.nodeId)
	fmt.Println(req)
	if digest, err := hash.Hash(req.Data); err != nil || digest != req.Digest {
		return nil, fmt.Errorf("Invalid Data")
	}

	if _, ok := n.prePrepareLog[req.SequenceId]; ok {
		return nil, fmt.Errorf("Duplicate sequenceId")
	}

	n.prePrepareLog[req.SequenceId] = &PrePrepareLog{
		digest: req.Digest,
		data:   req.Data,
	}

	p := parallel.New(ctx)
	for _, client := range n.nodeClients {
		func(client message.NodeServiceClient) {
			_ = p.Register(parallel.NewFunc(
				&message.PrepareRequest{
					ViewId: 0,
					Digest: math.TernaryOp(n.isByzantineNode() && rand.Intn(2) == 0,
						random.RandomString(1024),
						req.Digest,
					),
					SequenceId: req.SequenceId,
					NodeId:     n.nodeId,
				},
				func(ctx context.Context, request *message.PrepareRequest) (*message.PrepareResponse, error) {
					return client.Prepare(ctx, request)
				},
			))
		}(client)
	}
	return &message.PrePrepareResponse{}, nil
}

func (n *Node) Prepare(ctx context.Context, req *message.PrepareRequest) (*message.PrepareResponse, error) {
	fmt.Printf("Node %v received prepare request with content: \n", n.nodeId)
	fmt.Println(req)
	log, ok := n.prePrepareLog[req.SequenceId]
	if !ok || log.digest != req.Digest {
		return nil, fmt.Errorf("Invalid Prepare Request: digest mismatch or sequence ID not found")
	}

	entry, exists := n.prepareLog[req.SequenceId]
	if !exists {
		entry = &LogEntry{count: 1, processed: false}
		n.prepareLog[req.SequenceId] = entry
	}
	entry.count++

	if entry.count >= n.quorumSize() && !entry.processed {
		entry.processed = true

		p := parallel.New(ctx)
		for _, client := range n.nodeClients {
			func(client message.NodeServiceClient) {
				_ = p.Register(parallel.NewFunc(
					&message.CommitRequest{
						ViewId: req.ViewId,
						Digest: math.TernaryOp(n.isByzantineNode() && rand.Intn(2) == 0,
							random.RandomString(1024),
							req.Digest,
						),
						SequenceId: req.SequenceId,
						NodeId:     n.nodeId,
					},
					func(ctx context.Context, request *message.CommitRequest) (*message.CommitResponse, error) {
						return client.Commit(ctx, request)
					},
				))
			}(client)
		}
	}
	return &message.PrepareResponse{}, nil
}

func (n *Node) Commit(ctx context.Context, req *message.CommitRequest) (*message.CommitResponse, error) {
	fmt.Printf("Node %v received commit request with content: \n", n.nodeId)
	fmt.Println(req)
	log, ok := n.prePrepareLog[req.SequenceId]
	if !ok || log.digest != req.Digest {
		return nil, fmt.Errorf("Invalid Commit Request: digest mismatch or sequence ID not found")
	}

	entry, exists := n.commitLog[req.SequenceId]
	if !exists {
		entry = &LogEntry{count: 1, processed: false}
		n.commitLog[req.SequenceId] = entry
	}
	entry.count++

	if entry.count >= n.quorumSize() && !entry.processed {
		entry.processed = true
		// do something here in reality
		fmt.Printf("Transaction committed: sequence ID %d\n", req.SequenceId)

		delete(n.prePrepareLog, req.SequenceId)
		delete(n.prepareLog, req.SequenceId)
		delete(n.commitLog, req.SequenceId)
	}
	return &message.CommitResponse{}, nil
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

	if n.isPrimaryNode() {
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

	fmt.Printf("Node %v connect node: %v successfully\n", n.port, req.Port)
	return &message.AddNodeResponse{
		NodeId:         int32(len(clones) + 1),
		OtherNodePorts: clones,
	}, nil
}

func (n *Node) GetListOtherPort(ctx context.Context, req *message.GetListOtherPortRequest) (*message.GetListOtherPortResponse, error) {
	return &message.GetListOtherPortResponse{
		Ports: slice.Make(n.otherNodePorts...),
	}, nil
}

func (n *Node) isByzantineNode() bool {
	return n.isByzantine == true
}
func (n *Node) isPrimaryNode() bool {
	return n.nodeId == constz.PrimaryNodeId
}

func (n *Node) increaseSequenceId() int32 {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.sequenceId++
	return n.sequenceId
}
func (n *Node) quorumSize() int32 {
	return (int32(len(n.otherNodePorts))+1)*2/3 + 1
}
