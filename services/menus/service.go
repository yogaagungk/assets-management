package menus

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

func (service *Service) Save(entity Menu) string {
	tx, _ := service.repo.db.Begin()

	_, err := service.repo.Save(tx, entity)

	if err == nil {
		tx.Commit()

		return common.SAVE_SUCCESS
	} else {
		tx.Rollback()

		return common.SAVE_FAILED
	}
}

func (service *Service) Update(entity Menu) string {
	_, isFound := service.repo.FindByID(entity.ID)

	if !isFound {
		return common.DATA_NOT_FOUND
	} else {
		tx, _ := service.repo.db.Begin()

		_, rowAffected := service.repo.Update(tx, entity)

		if rowAffected == 1 {
			tx.Commit()

			return common.UPDATE_SUCCESS
		} else {
			tx.Rollback()

			return common.UPDATE_FAILED
		}
	}
}

func (service *Service) Delete(id uint64) string {
	menu, isFound := service.repo.FindByID(id)

	if !isFound {
		return common.DATA_NOT_FOUND
	} else {
		tx, _ := service.repo.db.Begin()

		_, rowAffected := service.repo.Delete(tx, menu)

		if rowAffected == 1 {
			tx.Commit()

			return common.DELETE_SUCCESS
		} else {
			tx.Rollback()

			return common.DELETE_FAILED
		}
	}
}

func (service *Service) Find(param Menu, offset string, limit string) ([]Menu, string) {
	menus, isFound := service.repo.Find(param, offset, limit)

	if !isFound {
		return nil, common.DATA_NOT_FOUND
	}

	return menus, common.DATA_FOUND
}

func (service *Service) Count(param Menu) uint {
	return service.repo.Count(param)
}

func (service *Service) FindByID(id uint64) (Menu, string) {
	menu, isFound := service.repo.FindByID(id)

	if !isFound {
		return Menu{}, common.DATA_NOT_FOUND
	}

	return menu, common.DATA_FOUND
}
