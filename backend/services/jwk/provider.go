package jwk

import (
	"context"

	"github.com/lestrrat-go/jwx/v3/jwk"
)

type Interface interface {
	Fetch(ctx context.Context, u string) (jwk.Set, error)
}

type Client struct{}

func (j *Client) Fetch(ctx context.Context, u string) (jwk.Set, error) {
	return jwk.Fetch(ctx, u)
}

func Provider() Interface {
	return &Client{}
}
