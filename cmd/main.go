package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/kallepan/hexarch-golang/config"

	"github.com/kallepan/hexarch-golang/internal/adapters/app/api"
	"github.com/kallepan/hexarch-golang/internal/adapters/core/arithmetic"
	rpc "github.com/kallepan/hexarch-golang/internal/adapters/framework/left/grpc"
	"github.com/kallepan/hexarch-golang/internal/adapters/framework/right/db"
)

func main() {
	ctx, close := context.WithCancel(context.Background())
	defer close()

	config.InitLogger()

	// right side adapters
	dbaseAdapter, err := db.NewAdapter(ctx)
	if err != nil {
		slog.Error("db adapter init failed", "error", err)
		os.Exit(1)
	}
	defer dbaseAdapter.CloseDbConnection(ctx)

	// left side adapters
	arithmeticAdapter := arithmetic.NewAdapter()
	apiAdapter := api.NewAdapter(arithmeticAdapter, dbaseAdapter)
	grpcAdapter := rpc.NewAdapter(apiAdapter)

	slog.Info("Starting server...")
	grpcAdapter.Run()
}
