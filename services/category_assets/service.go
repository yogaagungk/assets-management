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
	_, err := service.repo.Save(entity)

	if err != nil {
		return common.SAVE_FAILED
	} else {
		return common.SAVE_SUCCESS
	}
}

func (service *Service) Update(entity CategoryAsset) string {
	_, isNotFound := service.repo.FindByID(entity.ID)

	if isNotFound {
		return common.DATA_NOT_FOUND
	} else {
		_, rowAffected := service.repo.Update(entity)

		if rowAffected == 1 {
			return common.UPDATE_SUCCESS
		} else {
			return common.UPDATE_FAILED
		}
	}
}

func (service *Service) Delete(id uint64) string {
	categoryAsset, isNotFound := service.repo.FindByID(id)

	if isNotFound {
		return common.DATA_NOT_FOUND
	} else {
		_, rowAffected := service.repo.Delete(categoryAsset)

		if rowAffected == 1 {
			return common.DELETE_SUCCESS
		} else {
			return common.DELETE_FAILED
		}
	}
}

func (service *Service) Find(param CategoryAsset, offset string, limit string) ([]CategoryAsset, string) {
	categoryAssets, isNotFound := service.repo.Find(param, offset, limit)

	if isNotFound {
		return nil, common.DATA_NOT_FOUND
	}

	return categoryAssets, common.DATA_FOUND
}

func (service *Service) Count(param CategoryAsset) uint {
	return service.repo.Count(param)
}

func (service *Service) FindByID(id uint64) (CategoryAsset, string) {
	categoryAsset, isNotFound := service.repo.FindByID(id)

	if isNotFound {
		return CategoryAsset{}, common.DATA_NOT_FOUND
	}

	return categoryAsset, common.DATA_FOUND
}
