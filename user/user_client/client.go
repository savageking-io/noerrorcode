package user_client

import (
	"context"
	"fmt"
	"github.com/savageking-io/noerrorcode/user/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type UserClient struct {
	conn   *grpc.ClientConn
	client pb.UserClient
}

func (d *UserClient) Connect(grpcServerHostname string, grpcServerPort uint32) error {
	log.Traceln("UserClient::Connect")
	var err error
	d.conn, err = grpc.NewClient(fmt.Sprintf("%s:%d", grpcServerHostname, grpcServerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("grpc connect failed: %v", err)
	}

	d.client = pb.NewUserClient(d.conn)
	return nil
}

func (d *UserClient) Disconnect() error {
	log.Traceln("UserClient::Disconnect")
	if d.conn == nil {
		return fmt.Errorf("attempt to disconnect from a nil connection")
	}
	return d.conn.Close()
}

func (d *UserClient) AuthenticateCredentials(username, password string) (*pb.AuthResponse, error) {
	log.Traceln("UserClient::AuthenticateCredentials")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := d.client.AuthenticateCredentials(ctx, &pb.CredentialsAuthRequest{Username: username, Password: password})
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate credentials: %v", err)
	}
	log.Traceln("UserClient::AuthenticateCredentials: response: %+v", r)

	return r, nil
}

func (d *UserClient) AuthenticatePlatform(platformId, platformToken string) (*pb.AuthResponse, error) {
	log.Traceln("UserClient::AuthenticatePlatform")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := d.client.AuthenticatePlatform(ctx, &pb.PlatformAuthRequest{Platform: platformId, Token: platformToken})
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate platform: %v", err)
	}
	log.Traceln("UserClient::AuthenticatePlatform: response: %+v", r)
	return r, nil
}
