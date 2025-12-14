package users

import (
	"polimane/backend/repository"
)

func WorkosIDEq(id string) repository.Scope {
	return repository.Where("workos_id = ?", id)
}
