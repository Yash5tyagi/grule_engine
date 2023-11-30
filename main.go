package main

import (
	"fmt"
	"rule_engine/rule_engine"

	"github.com/hyperjumptech/grule-rule-engine/logger"
)

type User struct {
	Name              string  `json:"name"`
	Username          string  `json:"username"`
	Email             string  `json:"email"`
	Age               int     `json:"age"`
	Gender            string  `json:"gender"`
	TotalOrders       int     `json:"total_orders"`
	AverageOrderValue float64 `json:"average_order_value"`
}

type OfferService interface {
	CheckOfferForUser(user User) bool
}

type OfferServiceClient struct {
	ruleEngineSvc *rule_engine.RuleEngineSvc
}

func NewOfferService(ruleEngineSvc *rule_engine.RuleEngineSvc) OfferService {
	return &OfferServiceClient{
		ruleEngineSvc: ruleEngineSvc,
	}
}

func (svc OfferServiceClient) CheckOfferForUser(user User) bool {
	offerCard := rule_engine.NewUserOfferContext()
	offerCard.UserOfferInput = &rule_engine.UserOfferInput{
		Name:              user.Name,
		Username:          user.Username,
		Email:             user.Email,
		Gender:            user.Gender,
		Age:               user.Age,
		TotalOrders:       user.TotalOrders,
		AverageOrderValue: user.AverageOrderValue,
	}

	err := svc.ruleEngineSvc.Execute(offerCard)
	if err != nil {
		logger.Log.Error("get user offer rule engine failed", err)
	}

	return offerCard.UserOfferOutput.IsOfferApplicable
}

func main() {
	fmt.Println("TODO:implementing grule engine")
	ruleEngineSvc := rule_engine.NewRuleEngineSvc()
	offerSvc := NewOfferService(ruleEngineSvc)
	userA := User{
		Name:              "yash tyagi",
		Username:          "ytyagi",
		Email:             "yash@tyagi.com",
		Gender:            "Male",
		Age:               20,
		TotalOrders:       50,
		AverageOrderValue: 225,
	}

	fmt.Println("offer validity for user A: ", offerSvc.CheckOfferForUser(userA))

	userB := User{
		Name:              "name2",
		Username:          "user2",
		Email:             "user@abc.com",
		Gender:            "Female",
		Age:               25,
		TotalOrders:       10,
		AverageOrderValue: 80,
	}

	fmt.Println("offer validity for user B: ", offerSvc.CheckOfferForUser(userB))
}
