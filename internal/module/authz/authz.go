package authz

import (
	"context"

	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/BoruTamena/jobboard/internal/storage"
)

type authzModule struct {
	athzStorage storage.AuthzStorage
}

func InitAuthzModule(athz storage.AuthzStorage) module.AuthzModule {
	return &authzModule{
		athzStorage: athz,
	}
}

// check if user has a permission
func (athz *authzModule) CheckUserPermission(user, resource, act string) bool {

	permission, _ := athz.athzStorage.CheckPermision(user, resource, act)

	return permission
}

func (athz *authzModule) AddRoleForUser(ctx context.Context, user, role string) error {

	err := athz.athzStorage.AddRole(user, role)

	if err != nil {
		return err
	}

	return nil
}

func (athz *authzModule) AddPolicy(ctx context.Context, role, obj, act string) error {

	err := athz.athzStorage.AddPolicy(role, obj, act)

	if err != nil {
		// log error here
		return err
	}

	return nil
}
