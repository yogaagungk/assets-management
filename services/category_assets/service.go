package category_assets

import "github.com/yogaagungk/assets-management/common"

type Service struct {
	repo *Repository
}

// InjectDep is a function for inject db to Repository object
func ProvideService(repo *Repository) *Service {
	return &Service{repo}
}

func (service *Service) Save(entity CategoryAsset) string {
	tx, _ := service.repo.db.Begin()

	_, err := service.repo.Save(tx, entity)

	if err != nil {
		tx.Rollback()

		return common.SAVE_FAILED
	} else {
		tx.Commit()

		return common.SAVE_SUCCESS
	}
}

func (service *Service) Update(entity CategoryAsset) string {
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
	categoryAsset, isFound := service.repo.FindByID(id)

	if !isFound {
		return common.DATA_NOT_FOUND
	} else {
		tx, _ := service.repo.db.Begin()

		_, rowAffected := service.repo.Delete(tx, categoryAsset)

		if rowAffected == 1 {
			tx.Commit()

			return common.DELETE_SUCCESS
		} else {
			tx.Rollback()

			return common.DELETE_FAILED
		}
	}
}

func (service *Service) Find(param CategoryAsset, offset string, limit string) ([]CategoryAsset, string) {
	categoryAssets, isFound := service.repo.Find(param, offset, limit)

	if !isFound {
		return nil, common.DATA_NOT_FOUND
	}

	return categoryAssets, common.DATA_FOUND
}

func (service *Service) Count(param CategoryAsset) uint {
	return service.repo.Count(param)
}

func (service *Service) FindByID(id uint64) (CategoryAsset, string) {
	categoryAsset, isFound := service.repo.FindByID(id)

	if !isFound {
		return CategoryAsset{}, common.DATA_NOT_FOUND
	}

	return categoryAsset, common.DATA_FOUND
}
