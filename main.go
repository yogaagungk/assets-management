package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yogaagungk/assets-management/config"
	"github.com/yogaagungk/assets-management/di"
	"github.com/yogaagungk/assets-management/handler"
	"github.com/yogaagungk/assets-management/middleware"
	"github.com/yogaagungk/assets-management/services/menus"
	"github.com/yogaagungk/assets-management/services/roles"
	"github.com/yogaagungk/assets-management/services/users"
)

func main() {
	db := config.OpenDatabaseConnection()
	redis := config.OpenRedisPool()

	defer func() {
		db.Close()
		redis.Close()
	}()

	menusHandler := initializeMenuHandler()
	rolesHandler := initializeRoleHandler()
	usersHandler := initializeUserHandler()

	authorizationService := InitializeAuthorizationService()

	r := gin.Default()
	r.Use(gin.Recovery())

	authorizedMapping := r.Group("/")
	authorizedMapping.Use(authorizationService.Authorization())
	{
		// mapping MENU
		authorizedMapping.POST("/menus", menusHandler.Post)
		authorizedMapping.PUT("/menus", menusHandler.Put)
		authorizedMapping.GET("/menus", menusHandler.Get)
		authorizedMapping.DELETE("/menus/:id", menusHandler.Delete)

		// mapping Logout
		authorizedMapping.GET("/logout", usersHandler.Logout)
	}

	// mapping ROLE
	r.GET("/roles", rolesHandler.Get)

	// mapping Register and Login
	r.POST("/register", usersHandler.Register)
	r.POST("/login", usersHandler.Login)

	r.Run(":8081")
}

func initializeMenuHandler() handler.Menu {
	wire.Build(di.MenuRepositoryInjectionSet, menus.ProvideService, handler.ProvideMenu)

	return handler.Menu{}
}

func initializeRoleHandler() handler.Role {
	wire.Build(di.RoleRepositoryInjectionSet, roles.ProvideService, handler.ProvideRole)

	return handler.Role{}
}

func initializeUserHandler() handler.User {
	wire.Build(di.UserRepositoryInjectionSet, config.ProvideRedisPool, roles.ProvideRepo, users.ProvideService, handler.ProvideUser)

	return handler.User{}
}

func initializeAuthorizationService() middleware.AuthService {
	wire.Build(config.ProvideRedisPool, middleware.ProvideAuthService)

	return middleware.AuthService{}
}
