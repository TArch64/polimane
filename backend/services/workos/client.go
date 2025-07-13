package workos

import (
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/env"
)

var UserManagement *usermanagement.Client

func Init() {
	UserManagement = usermanagement.NewClient(env.Instance.WorkOS.ApiKey)
}
