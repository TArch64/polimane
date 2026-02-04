package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"polimane/backend/model"
)

type rowBuilder[K comparable, V any] func(key K, value V) (string, []any)

func NamedUpdateValues[K comparable, V any](set map[K]V) (string, []any) {
	idx := 0

	return buildUpdateValues(set, func(key K, value V) (string, []any) {
		idx++
		keyVar := sql.Named(fmt.Sprintf("uk_%d", idx), key)
		valueVar := sql.Named(fmt.Sprintf("uv_%d", idx), value)
		row := fmt.Sprintf("(%s, %s)", buildUpdateVar(keyVar), buildUpdateVar(valueVar))
		return row, []any{keyVar, valueVar}
	})
}

func buildUpdateValues[K comparable, V any](set map[K]V, row rowBuilder[K, V]) (string, []any) {
	builder := strings.Builder{}
	var args []any
	first := true

	for key, value := range set {
		if first {
			first = false
		} else {
			builder.WriteString(", ")
		}

		rowQuery, rowArgs := row(key, value)
		builder.WriteString(rowQuery)
		args = append(args, rowArgs...)
	}

	return builder.String(), args
}

func buildUpdateVar(arg sql.NamedArg) string {
	switch arg.Value.(type) {
	case model.ID:
		return fmt.Sprintf("(@%s)::uuid", arg.Name)
	case uint16, int16:
		return fmt.Sprintf("(@%s)::smallint", arg.Name)
	default:
		return fmt.Sprintf("@%s", arg.Name)
	}
}
