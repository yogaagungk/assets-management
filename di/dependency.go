package di

import (
	"github.com/google/wire"
	"github.com/yogaagungk/assets-management/config"
	"github.com/yogaagungk/assets-management/services/menus"
	"github.com/yogaagungk/assets-management/services/roles"
	"github.com/yogaagungk/assets-management/services/users"
)

var MenuRepositoryInjectionSet = wire.NewSet(config.ProvideDatabase, menus.ProvideRepo)

var RoleRepositoryInjectionSet = wire.NewSet(config.ProvideDatabase, roles.ProvideRepo)

var UserRepositoryInjectionSet = wire.NewSet(config.ProvideDatabase, users.ProvideRepo)
