package rpc

import (
	"context"
	"net"
	"testing"

	"github.com/kallepan/hexarch-golang/config"
	"github.com/kallepan/hexarch-golang/internal/adapters/app/api"
	"github.com/kallepan/hexarch-golang/internal/adapters/core/arithmetic"
	"github.com/kallepan/hexarch-golang/internal/adapters/framework/left/grpc/pb"
	"github.com/kallepan/hexarch-golang/internal/adapters/framework/right/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	config.InitLogger()
	lis = bufconn.Listen(bufSize)

	grpcServer := grpc.NewServer()

	// context
	ctx, close := context.WithCancel(context.Background())
	defer close()

	// ports
	dbaseAdapter, err := db.NewAdapter(ctx)
	if err != nil {
		panic(err)
	}

	// left side adapters
	arithmeticAdapter := arithmetic.NewAdapter()
	apiAdapter := api.NewAdapter(arithmeticAdapter, dbaseAdapter)
	grpcAdapter := NewAdapter(apiAdapter)

	pb.RegisterArithmeticServiceServer(grpcServer, grpcAdapter)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	t.Helper()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	return conn
}

func TestAddition(t *testing.T) {
	ctx, close := context.WithCancel(context.Background())
	defer close()

	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	req := &pb.OperationParameters{X: 1, Y: 2}
	resp, err := client.GetAddition(ctx, req)
	if err != nil {
		t.Fatalf("Addition() error = %v", err)
		return
	}
	if resp.Value != 3 {
		t.Fatalf("Addition() result = %v, want 3", resp.Value)
	}
}

func TestSubtraction(t *testing.T) {
	ctx, close := context.WithCancel(context.Background())
	defer close()

	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	req := &pb.OperationParameters{X: 1, Y: 2}
	resp, err := client.GetSubtraction(ctx, req)
	if err != nil {
		t.Fatalf("Subtraction() error = %v", err)
		return
	}
	if resp.Value != -1 {
		t.Fatalf("Subtraction() result = %v, want -1", resp.Value)
	}
}

func TestMultiplication(t *testing.T) {
	ctx, close := context.WithCancel(context.Background())
	defer close()

	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	req := &pb.OperationParameters{X: 1, Y: 2}
	resp, err := client.GetMultiplication(ctx, req)
	if err != nil {
		t.Fatalf("Multiplication() error = %v", err)
		return
	}
	if resp.Value != 2 {
		t.Fatalf("Multiplication() result = %v, want 2", resp.Value)
	}
}

func TestDivision(t *testing.T) {
	ctx, close := context.WithCancel(context.Background())
	defer close()

	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	req := &pb.OperationParameters{X: 1, Y: 2}
	resp, err := client.GetDivision(ctx, req)
	if err != nil {
		t.Fatalf("Division() error = %v", err)
		return
	}
	if resp.Value != 0 {
		t.Fatalf("Division() result = %v, want 0", resp.Value)
	}
}
