package model

type SubscriptionPlanID string

const (
	SubscriptionBeta  SubscriptionPlanID = "beta"
	SubscriptionBasic SubscriptionPlanID = "basic"
	SubscriptionPro   SubscriptionPlanID = "pro"
)

var (
	BasicPlan = &SubscriptionPlan{
		ID:           SubscriptionBasic,
		MonthlyPrice: 100,
		Tier:         1,

		Limits: &SubscriptionLimits{
			SchemasCreated: ptr[uint16](20),
			SchemaBeads:    ptr[uint16](1500),
			SharedAccess:   ptr[uint8](1),
		},
	}

	ProPlan = &SubscriptionPlan{
		ID:           SubscriptionPro,
		MonthlyPrice: 300,
		Tier:         2,

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
	Tier         uint8               `json:"tier"`
	MonthlyPrice float32             `json:"monthlyPrice"`
	Limits       *SubscriptionLimits `json:"limits"`
}

func (s *SubscriptionPlan) IsBeta() bool {
	return s.ID == SubscriptionBeta
}

type SubscriptionLimits struct {
	SchemasCreated *uint16 `json:"schemasCreated,omitempty"`
	SchemaBeads    *uint16 `json:"schemaBeads,omitempty"`
	SharedAccess   *uint8  `json:"sharedAccess,omitempty"`
}
