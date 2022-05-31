package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID 					primitive.ObjectID `bson:"_id"`
	Name 				string `bson:"name"`
	SearchString 		string `bson:"searchString"`
	Phone 				string	`bosn:"phone"`
	Email 				string `bson:"email"`
	HashedPassword 		string `bson:"hashedPassWord"`
	Banned 				bool `bson:"banned"`
	Avatar 				string `bson:"avatar"`
	Statistic 			UserStatistic `bson:"statistic"`
	CreatedAt 			time.Time `bson:"createdAt"`
	UpdatedAt 			time.Time `bson:"updatedAt"`
}

type UserStatistic struct {
	TotalOrder 				  int `bson:"totalOrder" json:"totalOrder"`
	TotalOrderPending         int `bson:"totalOrderPending" json:"totalOrderPending"`
	TotalOrderWaitingApproved int `bson:"totalOrderWaitingApproved" json:"totalOrderWaitingApproved"`
	TotalOrderDelivering      int `bson:"totalOrderDelivering" json:"totalOrderDelivering"`
	TotalOrderDelivered       int `bson:"totalOrderDelivered" json:"totalOrderDelivered"`
	TotalOrderCompleted       int `bson:"totalOrderCompleted" json:"totalOrderCompleted"`
	TotalOrderCanceled        int `bson:"totalOrderCanceled" json:"totalOrderCanceled"`
	TotalCustomer             int `bson:"totalCustomer" json:"totalCustomer"`

}