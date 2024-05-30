package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 1477
// Hash 9501
// Hash 9896
// Hash 5499
// Hash 2655
// Hash 6765
// Hash 6237
// Hash 2459
// Hash 3237
// Hash 3678
// Hash 5528
// Hash 6241
// Hash 2257
// Hash 1752
// Hash 9897
// Hash 8151
// Hash 4892
// Hash 1520
// Hash 7447
// Hash 4192
// Hash 4184
// Hash 3755
// Hash 7988
// Hash 8704
// Hash 1204
// Hash 3326
// Hash 2532
// Hash 6707
// Hash 1472
// Hash 4179
// Hash 8043
// Hash 3788
// Hash 1820
// Hash 6192
// Hash 7156