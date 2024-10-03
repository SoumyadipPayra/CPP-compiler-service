package main

import (
	"context"
	"fmt"

	cpb "github.com/SoumyadipPayra/protobufs/go-protos/cpp_compiler"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CPPCompilerServer struct {
	cpb.UnimplementedCPPCompilerServer
}

func (c *CPPCompilerServer) PingPong(_ context.Context, req *cpb.PingRequest) (*cpb.PingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "nil request")
	}
	return &cpb.PingResponse{Msg: fmt.Sprintf("ping pong reply : %s", req.GetMsg())}, nil
}

func (c *CPPCompilerServer) CompileAndRun(_ context.Context, req *cpb.CompileAndRunRequest) (*cpb.CompileAndRunResponse, error) {

}
