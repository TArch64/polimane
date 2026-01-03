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

type UserSubscription struct {
	*Identifiable
	*Timestamps
	Plan           SubscriptionPlan   `json:"plan"`
	Status         SubscriptionStatus `json:"status" gorm:"default:'active'"`
	BillingTry     uint8              `json:"billingTry" gorm:"default:0"`
	TrialStartedAt time.Time          `json:"trialStartedAt"`
	TrialEndsAt    time.Time          `json:"trialEndsAt"`
	CanceledAt     *time.Time         `json:"canceledAt"`
	LastBilledAt   *time.Time         `json:"lastBilledAt"`

	// Relations
	User *User `json:"-" gorm:"foreignKey:ID"`
}
