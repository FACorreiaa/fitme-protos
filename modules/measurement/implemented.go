package measurement

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
	client     generated.UserMeasurementsClient
}

var (
	_ generated.UserMeasurementsClient = (*Broker)(nil)
	_ core.Broker                      = (*Broker)(nil)
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
	b.client = generated.NewUserMeasurementsClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

func (b *Broker) CreateWeight(ctx context.Context, in *generated.CreateWeightReq, opts ...grpc.CallOption) (*generated.CreateWeightRes, error) {
	return b.client.CreateWeight(ctx, in, opts...)
}

func (b *Broker) GetWeights(ctx context.Context, in *generated.GetWeightsReq, opts ...grpc.CallOption) (*generated.GetWeightsRes, error) {
	return b.client.GetWeights(ctx, in, opts...)
}

func (b *Broker) GetWeight(ctx context.Context, in *generated.GetWeightReq, opts ...grpc.CallOption) (*generated.GetWeightRes, error) {
	return b.client.GetWeight(ctx, in, opts...)
}

func (b *Broker) DeleteWeight(ctx context.Context, in *generated.DeleteWeightReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.DeleteWeight(ctx, in, opts...)
}

func (b *Broker) UpdateWeight(ctx context.Context, in *generated.UpdateWeightReq, opts ...grpc.CallOption) (*generated.UpdateWeightRes, error) {
	return b.client.UpdateWeight(ctx, in, opts...)
}

func (b *Broker) CreateWaterMeasurement(ctx context.Context, in *generated.CreateWaterIntakeReq, opts ...grpc.CallOption) (*generated.CreateWaterIntakeRes, error) {
	return b.client.CreateWaterMeasurement(ctx, in, opts...)
}

func (b *Broker) GetWaterMeasurements(ctx context.Context, in *generated.GetWaterIntakesReq, opts ...grpc.CallOption) (*generated.GetWaterIntakesRes, error) {
	return b.client.GetWaterMeasurements(ctx, in, opts...)
}

func (b *Broker) GetWaterMeasurement(ctx context.Context, in *generated.GetWaterIntakeReq, opts ...grpc.CallOption) (*generated.GetWaterIntakeRes, error) {
	return b.client.GetWaterMeasurement(ctx, in, opts...)
}

func (b *Broker) DeleteWaterMeasurement(ctx context.Context, in *generated.DeleteWaterIntakeReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.DeleteWaterMeasurement(ctx, in, opts...)
}

func (b *Broker) UpdateWaterMeasurement(ctx context.Context, in *generated.UpdateWaterIntakeReq, opts ...grpc.CallOption) (*generated.UpdateWaterIntakeRes, error) {
	return b.client.UpdateWaterMeasurement(ctx, in, opts...)
}

func (b *Broker) CreateWasteLineMeasurement(ctx context.Context, in *generated.CreateWasteLineReq, opts ...grpc.CallOption) (*generated.CreateWasteLineRes, error) {
	return b.client.CreateWasteLineMeasurement(ctx, in, opts...)
}

func (b *Broker) GetWasteLineMeasurements(ctx context.Context, in *generated.GetWasteLinesReq, opts ...grpc.CallOption) (*generated.GetWasteLinesRes, error) {
	return b.client.GetWasteLineMeasurements(ctx, in, opts...)
}

func (b *Broker) GetWasteLineMeasurement(ctx context.Context, in *generated.GetWasteLineReq, opts ...grpc.CallOption) (*generated.GetWasteLineRes, error) {
	return b.client.GetWasteLineMeasurement(ctx, in, opts...)
}

func (b *Broker) DeleteWasteLineMeasurement(ctx context.Context, in *generated.DeleteWasteLineReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.DeleteWasteLineMeasurement(ctx, in, opts...)
}

func (b *Broker) UpdateWasteLineMeasurement(ctx context.Context, in *generated.UpdateWasteLineReq, opts ...grpc.CallOption) (*generated.UpdateWasteLineRes, error) {
	return b.client.UpdateWasteLineMeasurement(ctx, in, opts...)
}
