package roles

import (
	"github.com/yogaagungk/assets-management/common"
)

type Service struct {
	repo *Repository
}

// InjectDep is a function for inject db to Repository object
func ProvideService(repo *Repository) *Service {
	return &Service{repo}
}

func (service *Service) Find(param Role, offset string, limit string) ([]Role, string) {
	roles, isNotFound := service.repo.Find(param, offset, limit)

	if isNotFound {
		return nil, common.DATA_NOT_FOUND
	}

	return roles, common.DATA_FOUND
}

func (service *Service) FindByID(id uint64) (Role, string) {
	role, isNotFound := service.repo.FindByID(id)

	if isNotFound {
		return Role{}, common.DATA_NOT_FOUND
	}

	return role, common.DATA_FOUND
}

func (service *Service) FindByName(name string) (Role, string) {
	role, isNotFound := service.repo.FindByName(name)

	if isNotFound {
		return Role{}, common.DATA_NOT_FOUND
	}

	return role, common.DATA_FOUND
}
