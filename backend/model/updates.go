package model

import "github.com/guregu/dynamo/v2"

type UpdateValue interface{}
type UpdateOperation func(update *dynamo.Update) *dynamo.Update
type Updates []UpdateOperation

func NewUpdates() Updates {
	return make(Updates, 0)
}

func (u Updates) Set(key string, value UpdateValue) Updates {
	return append(u, func(update *dynamo.Update) *dynamo.Update {
		return update.Set(key, value)
	})
}

func (u Updates) Add(key string, value string) Updates {
	return append(u, func(update *dynamo.Update) *dynamo.Update {
		return update.AddStringsToSet(key, value)
	})
}

func (u Updates) Delete(key string, value string) Updates {
	return append(u, func(update *dynamo.Update) *dynamo.Update {
		return update.DeleteStringsFromSet(key, value)
	})
}

func (u Updates) Apply(update *dynamo.Update) *dynamo.Update {
	for _, apply := range u {
		update = apply(update)
	}
	return update
}
