package jwk

import (
	"context"

	"github.com/lestrrat-go/jwx/v3/jwk"
)

type Client interface {
	Fetch(ctx context.Context, u string) (jwk.Set, error)
}

type Impl struct{}

func (j *Impl) Fetch(ctx context.Context, u string) (jwk.Set, error) {
	return jwk.Fetch(ctx, u)
}

func Provider() Client {
	return &Impl{}
}
