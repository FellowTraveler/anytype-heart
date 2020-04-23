package cafe

import (
	"context"
	"crypto/x509"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/anytypeio/go-anytype-library/cafe/pb"
	"github.com/anytypeio/go-anytype-library/wallet"
	"github.com/mr-tron/base58"
	"github.com/textileio/go-threads/core/thread"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var _ pb.APIClient = (*Online)(nil)

const simultaneousRequests = 4

type Client interface {
	pb.APIClient
	Shutdown() error
}

type Token struct {
	Token     string
	ExpiresAt time.Time
}

type Online struct {
	client        pb.APIClient
	token         *Token
	getTokenMutex sync.Mutex

	throttle *int32

	device  wallet.Keypair
	account wallet.Keypair

	conn *grpc.ClientConn
}

func (c *Online) getSignature(payload string) (*pb.WithSignature, error) {
	as, err := c.account.Sign([]byte(payload))
	if err != nil {
		return nil, fmt.Errorf("can't create account signature")
	}

	asB58 := base58.Encode(as)
	ds, err := c.device.Sign([]byte(payload + asB58))
	if err != nil {
		return nil, fmt.Errorf("can't create device signature")
	}

	return &pb.WithSignature{
		AccountAddress:   c.account.Address(),
		DeviceAddress:    c.device.Address(),
		AccountSignature: asB58,
		DeviceSignature:  base58.Encode(ds),
	}, nil
}

func (c *Online) withToken(ctx context.Context) (context.Context, error) {
	token, err := c.requestToken(ctx)
	fmt.Printf("setToken 1 %p", ctx)

	if err != nil {
		return nil, err
	}

	ctx = thread.NewTokenContext(ctx, thread.Token(token.Token))
	fmt.Printf("setToken 2 %p", ctx)

	return ctx, nil
}

func (c *Online) requestToken(ctx context.Context) (*Token, error) {
	c.getTokenMutex.Lock()
	defer c.getTokenMutex.Unlock()
	if c.token != nil && c.token.ExpiresAt.After(time.Now().Add(time.Second*30)) {
		return c.token, nil
	}

	server, err := c.client.AuthGetToken(ctx)
	if err != nil {
		return nil, err
	}

	sig, err := c.getSignature("")
	if err != nil {
		return nil, err
	}

	err = server.Send(&pb.AuthGetTokenRequest{Signature: sig})
	if err != nil {
		return nil, fmt.Errorf("failed to send auth code request: %w", err)
	}

	resp, err := server.Recv()
	if err != nil {
		return nil, fmt.Errorf("failed to get auth code %w", err)
	}

	authCode := resp.GetAuthCode()
	sig, err = c.getSignature(authCode)
	if err != nil {
		return nil, err
	}

	err = server.Send(&pb.AuthGetTokenRequest{AuthCode: authCode, Signature: sig})
	if err != nil {
		return nil, fmt.Errorf("failed to send auth code request: %w", err)
	}

	resp, err = server.Recv()
	if err != nil {
		return nil, fmt.Errorf("failed to get token %w", err)
	}

	if resp.GetToken() == nil {
		return nil, fmt.Errorf("failed to get token: token is nil")
	}

	expiresAt := time.Unix(resp.GetToken().ExpiresAt, 0)
	c.token = &Token{Token: resp.GetToken().Token, ExpiresAt: expiresAt}

	return c.token, nil
}

func (c *Online) AuthGetToken(ctx context.Context, opts ...grpc.CallOption) (pb.API_AuthGetTokenClient, error) {
	return c.client.AuthGetToken(ctx, opts...)
}

func (c *Online) ThreadLogFollow(ctx context.Context, in *pb.ThreadLogFollowRequest, opts ...grpc.CallOption) (*pb.ThreadLogFollowResponse, error) {
	ctx, err := c.withToken(ctx)
	if err != nil {
		return nil, err
	}
	return c.client.ThreadLogFollow(ctx, in, opts...)
}

func (c *Online) FilePin(ctx context.Context, in *pb.FilePinRequest, opts ...grpc.CallOption) (*pb.FilePinResponse, error) {
	for {
		v := atomic.LoadInt32(c.throttle)
		if v >= simultaneousRequests {
			time.Sleep(time.Second)
			continue
		}

		if atomic.CompareAndSwapInt32(c.throttle, v, v+1) {
			break
		}
	}
	defer atomic.AddInt32(c.throttle, -1)

	ctx, err := c.withToken(ctx)
	if err != nil {
		return nil, err
	}

	return c.client.FilePin(ctx, in, opts...)
}

func (c *Online) ProfileFind(ctx context.Context, in *pb.ProfileFindRequest, opts ...grpc.CallOption) (pb.API_ProfileFindClient, error) {
	ctx, err := c.withToken(ctx)
	if err != nil {
		return nil, err
	}

	return c.client.ProfileFind(ctx, in, opts...)
}

func NewClient(url string, device wallet.Keypair, account wallet.Keypair) (Client, error) {
	certpool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(url, grpc.WithUserAgent("<todo>"), grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(certpool, "")), grpc.WithPerRPCCredentials(thread.Credentials{}))
	if err != nil {
		return nil, err
	}

	throttle := int32(0)
	return &Online{
		pb.NewAPIClient(conn),
		nil,
		sync.Mutex{},
		&throttle,
		device,
		account,
		conn,
	}, nil
}

func (c *Online) Shutdown() error {
	return c.conn.Close()
}
