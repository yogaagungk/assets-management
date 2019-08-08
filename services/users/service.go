package users

import (
	"encoding/json"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/yogaagungk/assets-management/common"
	"github.com/yogaagungk/assets-management/services/roles"
	"github.com/yogaagungk/assets-management/util/auth"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo     *Repository
	roleRepo *roles.Repository
	redis    redis.Conn
}

// InjectDep is a function for inject db to Repository object
func ProvideService(repo *Repository, roleRepo *roles.Repository, redis redis.Conn) *Service {
	return &Service{repo, roleRepo, redis}
}

func (service *Service) Register(param User) string {
	role, isNotFound := service.roleRepo.FindByName(common.ROLE_ADMIN) // find ID of ROLE_ADMIN

	if isNotFound {
		return common.SAVE_FAILED
	}

	var entity User
	entity.ID = param.ID
	entity.Name = param.Name
	entity.Username = param.Username
	entity.Password = hashPassword(param.Password) // hashing password
	entity.RoleId = role.ID                        // default role user is ROLE_ADMIN

	_, err := service.repo.Save(entity)

	if err != nil {
		return common.SAVE_FAILED
	} else {
		return common.SAVE_SUCCESS
	}
}

func (service *Service) Login(param User) (auth.UserAuth, string) {
	user, isNotFound := service.repo.FindByUsername(param.Username)

	if isNotFound {
		return auth.UserAuth{}, common.LOGIN_FAILED
	}

	if !checkPasswordHash(param.Password, user.Password) {
		return auth.UserAuth{}, common.LOGIN_FAILED
	}

	var currentUser auth.UserAuth
	currentUser.Name = user.Name
	currentUser.Username = user.Username
	currentUser.RoleName = user.Role.Name
	currentUser.Token = auth.GenerateToken(currentUser)

	currentUserString, err := json.Marshal(currentUser)

	if err != nil {
		log.Println(err.Error())
	}

	_, errR := service.redis.Do("SET", currentUser.Username, currentUserString)

	if errR != nil {
		log.Println(errR.Error())
	}

	return currentUser, common.LOGIN_SUCCESS
}

func (service *Service) Logout(username string) string {
	service.redis.Do("DEL", username)

	return common.LOGOUT_SUCCESS
}

func (service *Service) Update(entity User) string {
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
	user, isNotFound := service.repo.FindByID(id)

	if isNotFound {
		return common.DATA_NOT_FOUND
	} else {
		_, rowAffected := service.repo.Delete(user)

		if rowAffected == 1 {
			return common.DELETE_SUCCESS
		} else {
			return common.DELETE_FAILED
		}
	}
}

func (service *Service) Find(param User, offset string, limit string) ([]User, string) {
	users, isNotFound := service.repo.Find(param, offset, limit)

	if isNotFound {
		return nil, common.DATA_NOT_FOUND
	}

	return users, common.DATA_FOUND
}

func (service *Service) Count(param User) uint {
	return service.repo.Count(param)
}

func (service *Service) FindByID(id uint64) (User, string) {
	user, isNotFound := service.repo.FindByID(id)

	if isNotFound {
		return User{}, common.DATA_NOT_FOUND
	}

	return user, common.DATA_FOUND
}

// function to hash password using bcrypt
func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes)
}

// comapare hashed password and plain text using bcrypt
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
