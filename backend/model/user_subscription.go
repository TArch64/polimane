package model

import (
	"time"

	"gorm.io/datatypes"
)

type SubscriptionStatus string

const (
	SubscriptionActive   SubscriptionStatus = "active"
	SubscriptionCanceled SubscriptionStatus = "canceled"
	SubscriptionUnpaid   SubscriptionStatus = "unpaid"
)

var (
	SubscriptionTrialDuration = 14 * 24 * time.Hour
)

type UserSubscription struct {
	UserID         ID                       `gorm:"primaryKey" json:"-"`
	PlanID         SubscriptionPlanID       `json:"plan" gorm:"plan"`
	Status         SubscriptionStatus       `json:"status" gorm:"default:active"`
	Counters       SubscriptionCountersJSON `json:"counters" gorm:"default:'{}'::json"`
	BillingTry     uint8                    `json:"-" gorm:"default:0"`
	TrialStartedAt time.Time                `json:"-"`
	TrialEndsAt    time.Time                `json:"-"`
	CanceledAt     *time.Time               `json:"-"`
	LastBilledAt   *time.Time               `json:"-"`

	// Relations
	User *User `json:"-"`
}

type SubscriptionCounters struct {
	SchemasCreated uint16 `json:"schemasCreated"`
}

type SubscriptionCountersJSON = datatypes.JSONType[*SubscriptionCounters]

func (u *UserSubscription) Plan() *SubscriptionPlan {
	return Plans[u.PlanID]
}
