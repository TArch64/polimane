package appcontext

import "context"

type Ctx struct {
	context.Context
}

func Provider() *Ctx {
	return &Ctx{
		Context: context.Background(),
	}
}
