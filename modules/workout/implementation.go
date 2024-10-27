package workout

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/fitme-protos/core"
	"github.com/FACorreiaa/fitme-protos/modules/workout/generated"
	"github.com/FACorreiaa/fitme-protos/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     generated.WorkoutClient
}

var (
	_ generated.WorkoutClient = (*Broker)(nil)
	_ core.Broker             = (*Broker)(nil)
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
	b.client = generated.NewWorkoutClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetExercises(ctx context.Context, in *generated.GetExercisesReq, opts ...grpc.CallOption) (*generated.GetExercisesRes, error) {
	return b.client.GetExercises(ctx, in, opts...)
}

func (b *Broker) GetExerciseID(ctx context.Context, in *generated.GetExerciseIDReq, opts ...grpc.CallOption) (*generated.GetExerciseIDRes, error) {
	return b.client.GetExerciseID(ctx, in, opts...)
}

func (b *Broker) CreateExercise(ctx context.Context, in *generated.CreateExerciseReq, opts ...grpc.CallOption) (*generated.CreateExerciseRes, error) {
	return b.client.CreateExercise(ctx, in, opts...)
}

func (b *Broker) UpdateExercice(ctx context.Context, in *generated.UpdateExerciseReq, opts ...grpc.CallOption) (*generated.UpdateExerciseRes, error) {
	return b.client.UpdateExercice(ctx, in, opts...)
}

func (b *Broker) DeleteExercise(ctx context.Context, in *generated.DeleteExerciseReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.DeleteExercise(ctx, in, opts...)
}

func (b *Broker) GetWorkoutPlanExercises(ctx context.Context, in *generated.GetWorkoutPlanExercisesReq, opts ...grpc.CallOption) (*generated.GetWorkoutPlanExercisesRes, error) {
	return b.client.GetWorkoutPlanExercises(ctx, in, opts...)
}

func (b *Broker) GetExerciseByIdWorkoutPlan(ctx context.Context, in *generated.GetExerciseByIdWorkoutPlanReq, opts ...grpc.CallOption) (*generated.GetExerciseByIdWorkoutPlanRes, error) {
	return b.client.GetExerciseByIdWorkoutPlan(ctx, in, opts...)
}

func (b *Broker) DeleteExerciseByIdWorkoutPlan(ctx context.Context, in *generated.DeleteExerciseByIdWorkoutPlanReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.DeleteExerciseByIdWorkoutPlan(ctx, in, opts...)
}

func (b *Broker) UpdateExerciseByIdWorkoutPlan(ctx context.Context, in *generated.UpdateExerciseByIdWorkoutPlanReq, opts ...grpc.CallOption) (*generated.UpdateExerciseByIdWorkoutPlanRes, error) {
	return b.client.UpdateExerciseByIdWorkoutPlan(ctx, in, opts...)
}

func (b *Broker) InsertExerciseWorkoutPlan(ctx context.Context, in *generated.InsertExerciseWorkoutPlanReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.InsertExerciseWorkoutPlan(ctx, in, opts...)
}

func (b *Broker) GetWorkoutPlans(ctx context.Context, in *generated.GetWorkoutPlansReq, opts ...grpc.CallOption) (*generated.GetWorkoutPlansRes, error) {
	plans, err := b.client.GetWorkoutPlans(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return plans, nil
}

func (b *Broker) GetWorkoutPlan(ctx context.Context, in *generated.GetWorkoutPlanReq, opts ...grpc.CallOption) (*generated.GetWorkoutPlanRes, error) {
	return b.client.GetWorkoutPlan(ctx, in, opts...)
}

func (b *Broker) DeleteWorkoutPlan(ctx context.Context, in *generated.DeleteWorkoutPlanReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.DeleteWorkoutPlan(ctx, in, opts...)
}

func (b *Broker) UpdateWorkoutPlan(ctx context.Context, in *generated.UpdateWorkoutPlanReq, opts ...grpc.CallOption) (*generated.UpdateWorkoutPlanRes, error) {
	return b.client.UpdateWorkoutPlan(ctx, in, opts...)
}

func (b *Broker) InsertWorkoutPlan(ctx context.Context, in *generated.InsertWorkoutPlanReq, opts ...grpc.CallOption) (*generated.InsertWorkoutPlanRes, error) {
	return b.client.InsertWorkoutPlan(ctx, in, opts...)
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}
