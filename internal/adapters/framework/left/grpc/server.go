// Contains server startup code
package rpc

import (
	"log/slog"
	"net"

	"github.com/kallepan/hexarch-golang/internal/adapters/framework/left/grpc/pb"
	"github.com/kallepan/hexarch-golang/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca *Adapter) Run() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		slog.Error("failed to listen", "error", err)
		panic(err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()

	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		slog.Error("failed to serve", "error", err)
		panic(err)
	}
}
