package authz

import (
	"log"

	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/storage"
	"github.com/casbin/casbin/v2"
)

type authzStorage struct {
	enforcer *casbin.Enforcer
}

func InitAuthzStorage(cfg dto.Config) storage.AuthzStorage {

	// default adapter is used
	e, err := casbin.NewEnforcer(cfg.Authz.Model, cfg.Authz.Adapter)

	if err != nil {
		log.Fatal("failed to create a casbin enforcer", err.Error())
	}

	return &authzStorage{
		enforcer: e,
	}

}

func (athz *authzStorage) CheckPermision(sub, resource, act string) (bool, error) {

	return athz.enforcer.Enforce(sub, resource, act)
}

func (athz *authzStorage) AddRole(userId string, role string) error {

	_, err := athz.enforcer.AddRoleForUser(userId, role)

	if err != nil {
		return err
	}

	return athz.enforcer.SavePolicy()
}

func (athz *authzStorage) AddPolicy(role, obj, act string) error {

	_, err := athz.enforcer.AddPolicy(role, obj, act)

	if err != nil {
		return err
	}

	return athz.enforcer.SavePolicy()
}

func (athz *authzStorage) CreateRole(userId string, role string) error {

	_, err := athz.enforcer.AddGroupingPolicy(userId, role)

	if err != nil {
		return err
	}

	return nil
}
