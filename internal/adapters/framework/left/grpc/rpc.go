// Contains grpc endpoints
package rpc

import (
	"context"
	"errors"
	"log/slog"

	"github.com/kallepan/hexarch-golang/internal/adapters/framework/left/grpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca *Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	slog.Info("Got request for GetAddition")

	x := req.GetX()
	y := req.GetY()
	res := &pb.Answer{}

	result, err := grpca.api.GetAddition(x, y)
	if err != nil {
		return res, status.Error(codes.Internal, err.Error())
	}

	res.Value = result

	return res, nil
}

func (grpca *Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	slog.Info("Got request for GetSubtraction")

	x := req.GetX()
	y := req.GetY()
	res := &pb.Answer{}

	result, err := grpca.api.GetSubtraction(x, y)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res.Value = result

	return res, nil
}

func (grpca *Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	slog.Info("Got request for GetMultiplication")

	x := req.GetX()
	y := req.GetY()
	answer := &pb.Answer{}

	result, err := grpca.api.GetMultiplication(x, y)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	answer.Value = result

	return answer, nil
}

func (grpca *Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	slog.Info("Got request for GetDivision")

	x := req.GetX()
	y := req.GetY()
	answer := &pb.Answer{}

	if y == 0 {
		return nil, status.Error(codes.InvalidArgument, errors.New("division by zero").Error())
	}

	result, err := grpca.api.GetDivision(x, y)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	answer.Value = result

	return answer, nil
}
