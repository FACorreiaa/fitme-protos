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
	serverAddr                string
	conn                      *grpc.ClientConn
	mealPlanClient            generated.MealPlanClient
	mealTrackClient           generated.TrackMealProgressClient
	mealClient                generated.MealClient
	ingredientsClient         generated.IngredientsClient
	mealReminderClient        generated.MealReminderClient
	goalsRecommendationClient generated.GoalRecommendationClient
}

var (
	_ generated.MealPlanClient           = (*Broker)(nil)
	_ generated.TrackMealProgressClient  = (*Broker)(nil)
	_ generated.MealClient               = (*Broker)(nil)
	_ generated.IngredientsClient        = (*Broker)(nil)
	_ generated.MealReminderClient       = (*Broker)(nil)
	_ generated.GoalRecommendationClient = (*Broker)(nil)
	_ core.Broker                        = (*Broker)(nil)
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
	b.mealPlanClient = generated.NewMealPlanClient(b.conn)

	return b.conn, nil
}

func (b *Broker) GetAddress() string {
	return b.serverAddr
}

func (b *Broker) GetMealPlan(ctx context.Context, in *generated.GetMealPlanReq, opts ...grpc.CallOption) (*generated.GetMealPlanRes, error) {
	return b.mealPlanClient.GetMealPlan(ctx, in, opts...)
}

func (b *Broker) GetMealPlans(ctx context.Context, in *generated.GetMealPlansReq, opts ...grpc.CallOption) (*generated.GetMealPlansRes, error) {
	return b.mealPlanClient.GetMealPlans(ctx, in, opts...)
}

func (b *Broker) CreateMealPlan(ctx context.Context, in *generated.CreateMealPlanReq, opts ...grpc.CallOption) (*generated.CreateMealPlanRes, error) {
	return b.mealPlanClient.CreateMealPlan(ctx, in, opts...)
}

func (b *Broker) UpdateMealPlan(ctx context.Context, in *generated.UpdateMealPlanReq, opts ...grpc.CallOption) (*generated.UpdateMealPlanRes, error) {
	return b.mealPlanClient.UpdateMealPlan(ctx, in, opts...)
}

func (b *Broker) DeleteMealPlan(ctx context.Context, in *generated.DeleteMealPlanReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.mealPlanClient.DeleteMealPlan(ctx, in, opts...)
}

func (b *Broker) AddIngredientToMealPlan(ctx context.Context, in *generated.AddIngredientReq, opts ...grpc.CallOption) (*generated.AddIngredientRes, error) {
	return b.mealPlanClient.AddIngredientToMealPlan(ctx, in, opts...)
}

func (b *Broker) DeleteIngredientFromMealPlan(ctx context.Context, in *generated.DeleteIngredientReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.mealPlanClient.DeleteIngredientFromMealPlan(ctx, in, opts...)
}

func (b *Broker) CreateCalorieIntakeObjective(ctx context.Context, in *generated.CreateCalorieIntakeObjectiveReq, opts ...grpc.CallOption) (*generated.CreateCalorieIntakeObjectiveRes, error) {
	return b.mealPlanClient.CreateCalorieIntakeObjective(ctx, in, opts...)
}

func (b *Broker) UpdateCalorieIntakeObjective(ctx context.Context, in *generated.UpdateCalorieIntakeObjectiveReq, opts ...grpc.CallOption) (*generated.UpdateCalorieIntakeObjectiveRes, error) {
	return b.mealPlanClient.UpdateCalorieIntakeObjective(ctx, in, opts...)
}

func (b *Broker) DeleteCalorieIntakeObjective(ctx context.Context, in *generated.DeleteCalorieIntakeObjectiveReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.mealPlanClient.DeleteCalorieIntakeObjective(ctx, in, opts...)
}

func (b *Broker) GetUserProgress(ctx context.Context, in *generated.GetUserProgressReq, opts ...grpc.CallOption) (*generated.GetUserProgressRes, error) {
	return b.mealTrackClient.GetUserProgress(ctx, in, opts...)
}

func (b *Broker) GetAllProgress(ctx context.Context, in *generated.GetAllProgressReq, opts ...grpc.CallOption) (*generated.GetAllProgressRes, error) {
	return b.mealTrackClient.GetAllProgress(ctx, in, opts...)
}

func (b *Broker) GetAllStatistics(ctx context.Context, in *generated.GetAllStatisticsReq, opts ...grpc.CallOption) (*generated.GetAllStatisticsRes, error) {
	return b.mealTrackClient.GetAllStatistics(ctx, in, opts...)
}

func (b *Broker) GetMeal(ctx context.Context, in *generated.GetMealReq, opts ...grpc.CallOption) (*generated.GetMealRes, error) {
	return b.mealClient.GetMeal(ctx, in, opts...)
}

func (b *Broker) GetMeals(ctx context.Context, in *generated.GetMealsReq, opts ...grpc.CallOption) (*generated.GetMealsRes, error) {
	return b.mealClient.GetMeals(ctx, in, opts...)
}

func (b *Broker) CreateMeal(ctx context.Context, in *generated.CreateMealReq, opts ...grpc.CallOption) (*generated.CreateMealRes, error) {
	return b.mealClient.CreateMeal(ctx, in, opts...)
}

func (b *Broker) UpdateMeal(ctx context.Context, in *generated.UpdateMealReq, opts ...grpc.CallOption) (*generated.UpdateMealRes, error) {
	return b.mealClient.UpdateMeal(ctx, in, opts...)
}

func (b *Broker) DeleteMeal(ctx context.Context, in *generated.DeleteMealReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.mealClient.DeleteMeal(ctx, in, opts...)
}

func (b *Broker) AddIngredientToMeal(ctx context.Context, in *generated.AddIngredientReq, opts ...grpc.CallOption) (*generated.AddIngredientRes, error) {
	return b.mealClient.AddIngredientToMeal(ctx, in, opts...)
}

func (b *Broker) RemoveIngredientFromMeal(ctx context.Context, in *generated.DeleteIngredientReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.mealClient.RemoveIngredientFromMeal(ctx, in, opts...)
}

func (b *Broker) UpdateIngredientInMeal(ctx context.Context, in *generated.UpdateMealIngredientReq, opts ...grpc.CallOption) (*generated.UpdateIngredientRes, error) {
	return b.mealClient.UpdateIngredientInMeal(ctx, in, opts...)
}

func (b *Broker) GetMealIngredients(ctx context.Context, in *generated.GetMealIngredientsReq, opts ...grpc.CallOption) (*generated.GetMealIngredientsRes, error) {
	return b.mealClient.GetMealIngredients(ctx, in, opts...)
}

func (b *Broker) GetMealIngredient(ctx context.Context, in *generated.GetMealIngredientReq, opts ...grpc.CallOption) (*generated.GetMealIngredientRes, error) {
	return b.mealClient.GetMealIngredient(ctx, in, opts...)
}

func (b *Broker) GetIngredients(ctx context.Context, in *generated.GetIngredientsReq, opts ...grpc.CallOption) (*generated.GetIngredientsRes, error) {
	return b.ingredientsClient.GetIngredients(ctx, in, opts...)
}

func (b *Broker) GetIngredient(ctx context.Context, in *generated.GetIngredientReq, opts ...grpc.CallOption) (*generated.GetIngredientRes, error) {
	return b.ingredientsClient.GetIngredient(ctx, in, opts...)
}

func (b *Broker) CreateIngredient(ctx context.Context, in *generated.CreateIngredientReq, opts ...grpc.CallOption) (*generated.CreateIngredientRes, error) {
	return b.ingredientsClient.CreateIngredient(ctx, in, opts...)
}

func (b *Broker) UpdateIngredient(ctx context.Context, in *generated.UpdateIngredientReq, opts ...grpc.CallOption) (*generated.UpdateIngredientRes, error) {
	return b.ingredientsClient.UpdateIngredient(ctx, in, opts...)
}

func (b *Broker) DeleteIngredient(ctx context.Context, in *generated.DeleteIngredientReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.ingredientsClient.DeleteIngredient(ctx, in, opts...)
}

func (b *Broker) CreateReminder(ctx context.Context, in *generated.CreateReminderReq, opts ...grpc.CallOption) (*generated.CreateReminderRes, error) {
	return b.mealReminderClient.CreateReminder(ctx, in, opts...)
}

func (b *Broker) GetReminders(ctx context.Context, in *generated.GetRemindersReq, opts ...grpc.CallOption) (*generated.GetRemindersRes, error) {
	return b.mealReminderClient.GetReminders(ctx, in, opts...)
}

func (b *Broker) UpdateReminder(ctx context.Context, in *generated.UpdateReminderReq, opts ...grpc.CallOption) (*generated.UpdateReminderRes, error) {
	return b.mealReminderClient.UpdateReminder(ctx, in, opts...)
}

func (b *Broker) DeleteReminder(ctx context.Context, in *generated.DeleteReminderReq, opts ...grpc.CallOption) (*generated.NilRes, error) {
	return b.mealReminderClient.DeleteReminder(ctx, in, opts...)
}

func (b *Broker) RecommendCalorieObjective(ctx context.Context, in *generated.RecommendCalorieObjectiveReq, opts ...grpc.CallOption) (*generated.RecommendCalorieObjectiveRes, error) {
	return b.goalsRecommendationClient.RecommendCalorieObjective(ctx, in, opts...)
}

func (b *Broker) AdjustGoals(ctx context.Context, in *generated.AdjustGoalsReq, opts ...grpc.CallOption) (*generated.AdjustGoalsRes, error) {
	return b.goalsRecommendationClient.AdjustGoals(ctx, in, opts...)
}

func (b *Broker) GetGoalSuggestions(ctx context.Context, in *generated.GetGoalSuggestionsReq, opts ...grpc.CallOption) (*generated.GetGoalSuggestionsRes, error) {
	return b.goalsRecommendationClient.GetGoalSuggestions(ctx, in, opts...)
}
