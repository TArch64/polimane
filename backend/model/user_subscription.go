package model

import (
	"time"

	t "gorm.io/datatypes"
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

	BetaLimits = SubscriptionLimits{}

	BasicLimits = SubscriptionLimits{
		SchemasCreated: ptr[uint16](20),
		SharedAccess:   ptr[uint8](1),
	}

	ProLimits = SubscriptionLimits{
		SchemasCreated: ptr[uint16](100),
		SharedAccess:   ptr[uint8](5),
	}
)

type UserSubscription struct {
	UserID         ID                                `gorm:"primaryKey" json:"-"`
	Plan           SubscriptionPlan                  `json:"plan"`
	Status         SubscriptionStatus                `json:"status" gorm:"default:active"`
	Counters       t.JSONType[*SubscriptionCounters] `json:"counters" gorm:"default:'{}'::json"`
	BillingTry     uint8                             `json:"-" gorm:"default:0"`
	TrialStartedAt time.Time                         `json:"-"`
	TrialEndsAt    time.Time                         `json:"-"`
	CanceledAt     *time.Time                        `json:"-"`
	LastBilledAt   *time.Time                        `json:"-"`

	// Relations
	User *User `json:"-"`
}

type SubscriptionCounters struct {
	SchemasCreated uint16 `json:"schemasCreated"`
}

type SubscriptionLimits struct {
	SchemasCreated *uint16 `json:"schemasCreated,omitempty"`
	SharedAccess   *uint8  `json:"sharedAccess,omitempty"`
}

func (u *UserSubscription) Limits() *SubscriptionLimits {
	switch u.Plan {
	case SubscriptionBeta:
		return &BetaLimits
	case SubscriptionBasic:
		return &BasicLimits
	case SubscriptionPro:
		return &ProLimits
	default:
		panic("unknown subscription plan")
	}
}
