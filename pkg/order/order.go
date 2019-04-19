package order

import (
	"time"

	"github.com/ashkarin/ashkarin-pizza-shop/pkg/pizza"
)

// Status is a type to represent order status
type Status int

// Order stages
const (
	StatusUnknown Status = iota + 1
	StatusAccepted
	StatusInProcess
	StatusCompleted
	StatusServed
)

// Order is an order in the system
type Order struct {
	ID          interface{}    `json:"_id,omitempty" bson:"_id,omitempty"`
	AcceptedAt  time.Time      `json:"acceptedAt" bson:"acceptedAt"`
	CompletedAt time.Time      `json:"completedAt" bson:"completedAt"`
	Content     []*pizza.Pizza `json:"content" bson:"content"`
	Status      Status         `json:"status" bson:"status"`
}
