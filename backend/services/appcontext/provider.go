package appcontext

import "context"

func Provider() context.Context {
	return context.Background()
}
