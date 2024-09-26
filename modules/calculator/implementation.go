package calculator

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/fitme-protos/core"
	"github.com/FACorreiaa/fitme-protos/modules/calculator/generated"
	"github.com/FACorreiaa/fitme-protos/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     generated.CalculatorServiceClient
}

var (
	_ generated.CalculatorServiceClient = (*Broker)(nil)
	_ core.Broker                       = (*Broker)(nil)
)

func NewBroker(serverAddr string) (*Broker, error) {
	b := new(Broker)
	b.serverAddr = serverAddr

	if b.serverAddr == "" {
		return nil, errors.New("null routed upstream host")
	}

	return b, nil
}

func (b *Broker) NewConnection() (*grpc.ClientConn, error) {
	conn, err := utils.NewConnection(b.serverAddr)
	if err != nil {
		return nil, errors.New("could not open connection")
	}

	b.conn = conn
	b.client = generated.NewCalculatorServiceClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

func (b *Broker) CreateUserMacro(ctx context.Context, in *generated.CreateUserMacroRequest, opts ...grpc.CallOption) (*generated.CreateUserMacroResponse, error) {
	return b.client.CreateUserMacro(ctx, in, opts...)
}

func (b *Broker) GetUserMacros(ctx context.Context, in *generated.GetUserMacroRequest, opts ...grpc.CallOption) (*generated.GetUserMacroResponse, error) {
	return b.client.GetUserMacros(ctx, in, opts...)
}

func (b *Broker) GetUsersMacros(ctx context.Context, in *generated.GetAllUserMacrosRequest, opts ...grpc.CallOption) (*generated.GetAllUserMacrosResponse, error) {
	return b.client.GetUsersMacros(ctx, in, opts...)
}
