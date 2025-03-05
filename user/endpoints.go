package main

import (
	"context"
	"github.com/savageking-io/noerrorcode/user/pb"
)

type endpoints struct {
	pb.UnimplementedUserServer
}

func (e *endpoints) AuthenticateCredentials(ctx context.Context, req *pb.CredentialsAuthRequest) (*pb.AuthResponse, error) {
	return nil, nil
}

func (e *endpoints) AuthenticatePlatform(ctx context.Context, req *pb.PlatformAuthRequest) (*pb.AuthResponse, error) {
	return nil, nil
}
