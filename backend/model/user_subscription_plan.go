package model

type SubscriptionPlanID string

const (
	SubscriptionBeta  SubscriptionPlanID = "beta"
	SubscriptionBasic SubscriptionPlanID = "basic"
	SubscriptionPro   SubscriptionPlanID = "pro"

	YearlyDiscountModifier = 0.8
)

var (
	BasicPlan = &SubscriptionPlan{
		ID:           SubscriptionBasic,
		MonthlyPrice: 100,
		YearlyPrice:  (100 * 12) * YearlyDiscountModifier,

		Limits: &SubscriptionLimits{
			SchemasCreated: ptr[uint16](20),
			SchemaBeads:    ptr[uint16](1500),
			SharedAccess:   ptr[uint8](1),
		},
	}

	ProPlan = &SubscriptionPlan{
		ID:           SubscriptionPro,
		MonthlyPrice: 300,
		YearlyPrice:  (300 * 12) * YearlyDiscountModifier,

		Limits: &SubscriptionLimits{
			SchemasCreated: ptr[uint16](100),
			SchemaBeads:    nil,
			SharedAccess:   ptr[uint8](5),
		},
	}

	BetaPlan = &SubscriptionPlan{
		ID:     SubscriptionBeta,
		Limits: ProPlan.Limits,
	}

	Plans = map[SubscriptionPlanID]*SubscriptionPlan{
		SubscriptionBeta:  BetaPlan,
		SubscriptionBasic: BasicPlan,
		SubscriptionPro:   ProPlan,
	}
)

type SubscriptionPlan struct {
	ID           SubscriptionPlanID  `json:"id"`
	MonthlyPrice float32             `json:"monthlyPrice"`
	YearlyPrice  float32             `json:"yearlyPrice"`
	Limits       *SubscriptionLimits `json:"limits"`
}

type SubscriptionLimits struct {
	SchemasCreated *uint16 `json:"schemasCreated,omitempty"`
	SchemaBeads    *uint16 `json:"schemaBeads,omitempty"`
	SharedAccess   *uint8  `json:"sharedAccess,omitempty"`
}
