package main

import (
	"context"
	"fmt"

	filehandler "github.com/SoumyadipPayra/CPP-compiler-service/src/fileHandler"
	"github.com/SoumyadipPayra/CPP-compiler-service/src/validate"
	cpb "github.com/SoumyadipPayra/protobufs/go-protos/cpp_compiler"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CPPCompilerServer struct {
	cpb.UnimplementedCPPCompilerServer
}

func (c *CPPCompilerServer) PingPong(ctx context.Context, req *cpb.PingRequest) (*cpb.PingResponse, error) {
	if err := validate.PingPong(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &cpb.PingResponse{Msg: fmt.Sprintf("ping pong reply : %s", req.GetMsg())}, nil
}

func (c *CPPCompilerServer) CompileAndRun(ctx context.Context, req *cpb.CompileAndRunRequest) (*cpb.CompileAndRunResponse, error) {
	if err := validate.CompileAndRun(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logBytes, err := filehandler.Handle(ctx, req.GetUserName(), req.GetFilePath(), req.GetFileData())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &cpb.CompileAndRunResponse{
		Logs: logBytes,
	}, nil
}
