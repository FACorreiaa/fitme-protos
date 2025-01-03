package meal

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/FACorreiaa/fitme-protos/core"
	"github.com/FACorreiaa/fitme-protos/modules/meal/generated"
	"github.com/FACorreiaa/fitme-protos/utils"
)

type Broker struct {
	serverAddr string
	conn       *grpc.ClientConn
	client     generated.MealPlanClient
}

var (
	_ generated.MealPlanClient = (*Broker)(nil)
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
	b.client = generated.NewMealPlanClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

func (b *Broker) GetMealPlan(ctx context.Context, in *generated.GetMealPlanReq, opts ...grpc.CallOption) (*generated.GetMealPlanRes, error) {
	return b.client.GetMealPlan(ctx, in, opts...)
}

func (b *Broker) GetMealPlans(ctx context.Context, in *generated.GetMealPlansReq, opts ...grpc.CallOption) (*generated.GetMealPlansRes, error) {
	return b.client.GetMealPlans(ctx, in, opts...)
}

func (b *Broker) CreateMealPlan(ctx context.Context, in *generated.CreateMealPlanReq, opts ...grpc.CallOption) (*generated.CreateMealPlanRes, error) {
	return b.client.CreateMealPlan(ctx, in, opts...)
}

func (b *Broker) UpdateMealPlan(ctx context.Context, in *generated.UpdateMealPlanReq, opts ...grpc.CallOption) (*generated.UpdateMealPlanRes, error) {
	return b.client.UpdateMealPlan(ctx, in, opts...)
}

func (b *Broker) DeleteMealPlan(ctx context.Context, in *generated.DeleteMealPlanReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.DeleteMealPlan(ctx, in, opts...)
}

func (b *Broker) AddIngredientToMealPlan(ctx context.Context, in *generated.AddIngredientReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.AddIngredientToMealPlan(ctx, in, opts...)
}

func (b *Broker) DeleteIngredientFromMealPlan(ctx context.Context, in *generated.DeleteIngredientReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.DeleteIngredientFromMealPlan(ctx, in, opts...)
}

func (b *Broker) CreateCalorieIntakeObjective(ctx context.Context, in *generated.CreateCalorieIntakeObjectiveReq, opts ...grpc.CallOption) (*generated.CreateCalorieIntakeObjectiveRes, error) {
	return b.client.CreateCalorieIntakeObjective(ctx, in, opts...)
}

func (b *Broker) UpdateCalorieIntakeObjective(ctx context.Context, in *generated.UpdateCalorieIntakeObjectiveReq, opts ...grpc.CallOption) (*generated.UpdateCalorieIntakeObjectiveRes, error) {
	return b.client.UpdateCalorieIntakeObjective(ctx, in, opts...)
}

func (b *Broker) DeleteCalorieIntakeObjective(ctx context.Context, in *generated.DeleteCalorieIntakeObjectiveReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.client.DeleteCalorieIntakeObjective(ctx, in, opts...)
}
