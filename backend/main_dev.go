//go:build dev

package main

import "polimane/backend/app"

func main() {
	err := app.New().Listen(":3000")
	if err != nil {
		panic(err)
	}
}
