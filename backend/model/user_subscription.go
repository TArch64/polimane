package model

import (
	"time"
)

type SubscriptionPlan string
type SubscriptionStatus string

const (
	SubscriptionBeta  SubscriptionPlan = "beta"
	SubscriptionBasic SubscriptionPlan = "basic"
	SubscriptionPro   SubscriptionPlan = "pro"

	SubscriptionActive   SubscriptionStatus = "active"
	SubscriptionCanceled SubscriptionStatus = "canceled"
	SubscriptionUnpaid   SubscriptionStatus = "unpaid"
)

var (
	SubscriptionTrialDuration = 14 * 24 * time.Hour
)

type UserSubscription struct {
	*Identifiable
	*Timestamps
	Plan           SubscriptionPlan   `json:"plan"`
	Status         SubscriptionStatus `json:"status" gorm:"default:active"`
	BillingTry     uint8              `json:"-" gorm:"default:0"`
	TrialStartedAt time.Time          `json:"-"`
	TrialEndsAt    time.Time          `json:"-"`
	CanceledAt     *time.Time         `json:"-"`
	LastBilledAt   *time.Time         `json:"-"`

	// Relations
	User *User `json:"-" gorm:"foreignKey:ID"`
}
