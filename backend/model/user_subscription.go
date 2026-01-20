package model

import (
	"time"

	"gorm.io/datatypes"
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

	BasicLimits = SubscriptionLimits{
		SchemasCreated: ptr[uint16](20),
		SchemaBeads:    ptr[uint16](1500),
		SharedAccess:   ptr[uint8](1),
	}

	ProLimits = SubscriptionLimits{
		SchemasCreated: ptr[uint16](100),
		SchemaBeads:    nil,
		SharedAccess:   ptr[uint8](5),
	}
)

type UserSubscription struct {
	UserID         ID                       `gorm:"primaryKey" json:"-"`
	Plan           SubscriptionPlan         `json:"plan"`
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

type SubscriptionLimits struct {
	SchemasCreated *uint16 `json:"schemasCreated,omitempty"`
	SchemaBeads    *uint16 `json:"schemaBeads,omitempty"`
	SharedAccess   *uint8  `json:"sharedAccess,omitempty"`
}

func (u *UserSubscription) Limits() *SubscriptionLimits {
	switch u.Plan {
	case SubscriptionBeta:
		return &ProLimits
	case SubscriptionBasic:
		return &BasicLimits
	case SubscriptionPro:
		return &ProLimits
	default:
		panic("unknown subscription plan")
	}
}
