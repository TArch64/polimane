package main

import (
	"fmt"
	"io"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"

	"polimane/backend/model"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(
		&model.User{},
		&model.Schema{},
		&model.UserSchema{},
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}

	io.WriteString(os.Stdout, stmts)
}
