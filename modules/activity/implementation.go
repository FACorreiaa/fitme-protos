package activity

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/fitme-protos/core"
	"github.com/FACorreiaa/fitme-protos/modules/activity/generated"
	"github.com/FACorreiaa/fitme-protos/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     generated.ActivityClient
}

var (
	_ generated.ActivityClient = (*Broker)(nil)
	_ core.Broker              = (*Broker)(nil)
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
	b.client = generated.NewActivityClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

func (b *Broker) GetActivity(ctx context.Context, in *generated.GetActivityReq, opts ...grpc.CallOption) (*generated.GetActivityRes, error) {
	//TODO implement me
	return b.client.GetActivity(ctx, in, opts...)
}

func (b *Broker) GetActivitiesByID(ctx context.Context, in *generated.GetActivityIDReq, opts ...grpc.CallOption) (*generated.GetActivityIDRes, error) {
	//TODO implement me
	return b.client.GetActivitiesByID(ctx, in, opts...)
}

func (b *Broker) GetActivitiesByName(ctx context.Context, in *generated.GetActivityNameReq, opts ...grpc.CallOption) (*generated.GetActivityNameRes, error) {
	//TODO implement me
	return b.client.GetActivitiesByName(ctx, in, opts...)
}

func (b *Broker) GetUserExerciseSession(ctx context.Context, in *generated.GetUserExerciseSessionReq, opts ...grpc.CallOption) (*generated.GetUserExerciseSessionRes, error) {
	//TODO implement me
	return b.client.GetUserExerciseSession(ctx, in, opts...)
}

func (b *Broker) GetUserExerciseTotalData(ctx context.Context, in *generated.GetUserExerciseTotalDataReq, opts ...grpc.CallOption) (*generated.GetUserExerciseTotalDataRes, error) {
	//TODO implement me
	return b.client.GetUserExerciseTotalData(ctx, in, opts...)
}

func (b *Broker) GetUserExerciseSessionStats(ctx context.Context, in *generated.GetUserExerciseSessionStatsReq, opts ...grpc.CallOption) (*generated.GetUserExerciseSessionStatsRes, error) {
	//TODO implement me
	return b.client.GetUserExerciseSessionStats(ctx, in, opts...)
}

func (b *Broker) GetExerciseSessionStats(ctx context.Context, in *generated.GetExerciseSessionStatsOccurrenceReq, opts ...grpc.CallOption) (*generated.GetExerciseSessionStatsOccurrenceRes, error) {
	//TODO implement me
	return b.client.GetExerciseSessionStats(ctx, in, opts...)
}

func (b *Broker) StartActivityTracker(ctx context.Context, in *generated.StartActivityTrackerReq, opts ...grpc.CallOption) (*generated.StartActivityTrackerRes, error) {
	//TODO implement me
	return b.client.StartActivityTracker(ctx, in, opts...)
}

func (b *Broker) PauseActivityTracker(ctx context.Context, in *generated.PauseActivityTrackerReq, opts ...grpc.CallOption) (*generated.PauseActivityTrackerRes, error) {
	//TODO implement me
	return b.client.PauseActivityTracker(ctx, in, opts...)
}

func (b *Broker) ResumeActivityTracker(ctx context.Context, in *generated.ResumeActivityTrackerReq, opts ...grpc.CallOption) (*generated.ResumeActivityTrackerRes, error) {
	//TODO implement me
	return b.client.ResumeActivityTracker(ctx, in, opts...)
}

func (b *Broker) StopActivityTracker(ctx context.Context, in *generated.StopActivityTrackerReq, opts ...grpc.CallOption) (*generated.StopActivityTrackerRes, error) {
	//TODO implement me
	return b.client.StopActivityTracker(ctx, in, opts...)
}

func (b *Broker) DeleteExerciseSession(ctx context.Context, in *generated.DeleteExerciseSessionReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	//TODO implement me
	return b.client.DeleteExerciseSession(ctx, in, opts...)
}

func (b *Broker) DeleteAllExercisesSession(ctx context.Context, in *generated.DeleteAllExercisesSessionReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	//TODO implement me
	return b.client.DeleteAllExercisesSession(ctx, in, opts...)
}
