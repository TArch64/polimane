package workos

import (
	"context"

	"github.com/workos/workos-go/v4/pkg/mfa"
)

type MFA interface {
	DeleteFactor(ctx context.Context, opts mfa.DeleteFactorOpts) error
	VerifyChallenge(ctx context.Context, opts mfa.VerifyChallengeOpts) (mfa.VerifyChallengeResponse, error)
	GetFactor(ctx context.Context, opts mfa.GetFactorOpts) (mfa.Factor, error)
}

func (i *Impl) MFA() MFA {
	return i.mfa
}
