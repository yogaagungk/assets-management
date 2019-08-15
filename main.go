package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yogaagungk/assets-management/config"
	"github.com/yogaagungk/assets-management/handler"
	"github.com/yogaagungk/assets-management/middleware"
	"github.com/yogaagungk/assets-management/services/menus"
	"github.com/yogaagungk/assets-management/services/roles"
	"github.com/yogaagungk/assets-management/services/users"

	_ "github.com/go-sql-driver/mysql"
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

	authorizationService := initializeAuthorizationService()

	r := gin.Default()
	r.Use(gin.Recovery())

	authorizedMapping := r.Group("/")
	authorizedMapping.Use(authorizationService.Authorization())
	{
		authorizedMapping.POST("/menus", menusHandler.Post)
		authorizedMapping.PUT("/menus", menusHandler.Put)
		authorizedMapping.GET("/menus", menusHandler.Get)
		authorizedMapping.DELETE("/menus/:id", menusHandler.Delete)

		// mapping Logout
		authorizedMapping.GET("/logout", usersHandler.Logout)
	}

	// mapping ROLE
	r.GET("/roles", rolesHandler.Get)

	// mapping USER
	r.GET("users", usersHandler.Get)

	// mapping Register and Login
	r.POST("/register", usersHandler.Register)
	r.POST("/login", usersHandler.Login)

	r.Run(":8080")
}

func initializeMenuHandler() handler.Menu {
	db := config.ProvideDatabase()
	repository := menus.ProvideRepo(db)
	service := menus.ProvideService(repository)
	menu := handler.ProvideMenu(service)
	return menu
}

func initializeRoleHandler() handler.Role {
	db := config.ProvideDatabase()
	repository := roles.ProvideRepo(db)
	service := roles.ProvideService(repository)
	role := handler.ProvideRole(service)
	return role
}

func initializeAuthorizationService() middleware.AuthService {
	conn := config.ProvideRedisPool()
	authService := middleware.ProvideAuthService(conn)
	return authService
}

func initializeUserHandler() handler.User {
	db := config.ProvideDatabase()
	repository := users.ProvideRepo(db)
	rolesRepository := roles.ProvideRepo(db)
	conn := config.ProvideRedisPool()
	service := users.ProvideService(repository, rolesRepository, conn)
	user := handler.ProvideUser(service)
	return user
}
