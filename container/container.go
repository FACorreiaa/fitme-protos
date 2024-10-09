package container

import (
	apb "github.com/FACorreiaa/fitme-protos/modules/activity/generated"
	calculator "github.com/FACorreiaa/fitme-protos/modules/calculator/generated"

	customer "github.com/FACorreiaa/fitme-protos/modules/customer/generated"
	user "github.com/FACorreiaa/fitme-protos/modules/user/generated"

	"github.com/FACorreiaa/fitme-protos/utils"
)

// Brokers is the encapsulation for all grpc clients. We use it this way so
// that we can:
//
// 1. Mock remotes in tests
// 2. Have a single point of change when contracts change
// 3. Use a common logic route for onboarding
//
// Ensure that you're using the interface type here and not the implementation
type Brokers struct {
	Customer       customer.CustomerClient
	Auth           user.AuthClient
	Calculator     calculator.CalculatorClient
	Activity       apb.ActivityClient
	TransportUtils *utils.TransportUtils
}

//func NewBrokers(transportUtils *utils.TransportUtils, auth *user.AuthClient, customer *customer.CustomerClient,
//	calc *calculator.CalculatorClient, activity *apb.ActivityClient) *Brokers {
//	if transportUtils == nil {
//		return nil
//	}
//
//	brokers := new(Brokers)
//	brokers.TransportUtils = transportUtils
//	brokers.Customer = customer
//	brokers.Auth = auth
//	brokers.Calculator = calc
//	brokers.Activity = activity
//	utils.Transport = transportUtils
//
//	return brokers
//}

// NewBrokers creates a common container instance for use with all cluster sevices
func NewBrokers(transportUtils *utils.TransportUtils) *Brokers {
	if transportUtils == nil {
		return nil
	}

	brokers := new(Brokers)
	brokers.TransportUtils = transportUtils
	//brokers.Customer = customer
	// brokers.Auth = auth
	// brokers.Calculator = calc
	utils.Transport = transportUtils

	return brokers
}
