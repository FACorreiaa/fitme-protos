package measurementpackage

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/fitme-protos/core"
	"github.com/FACorreiaa/fitme-protos/modules/measurement/generated"
	"github.com/FACorreiaa/fitme-protos/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     generated.MeasurementsClient
}

var (
	_ generated.MeasurementsClient = (*Broker)(nil)
	_ core.Broker                  = (*Broker)(nil)
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
	b.client = generated.NewMeasurementsClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

func (b *Broker) GetWeights(ctx context.Context, in *generated.NilRes, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.GetWeights(ctx, in, opts...)
}