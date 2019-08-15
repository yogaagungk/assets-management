package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yogaagungk/assets-management/services/users"
	"github.com/yogaagungk/assets-management/util/auth"
)

type User struct {
	service *users.Service
}

func ProvideUser(service *users.Service) User {
	return User{service}
}

func (handler *User) Register(c *gin.Context) {
	var entity users.User

	c.BindJSON(&entity)

	status := handler.service.Register(entity)

	c.JSON(http.StatusOK, gin.H{"status": status})
}

func (handler *User) Login(c *gin.Context) {
	var entity users.User

	c.BindJSON(&entity)

	currentUser, status := handler.service.Login(entity)

	c.JSON(http.StatusOK, gin.H{"status": status, "data": currentUser})
}

func (handler *User) Logout(c *gin.Context) {
	tokenHeader := c.GetHeader("Authorization")

	tokenString := tokenHeader[7:len(tokenHeader)]

	status := handler.service.Logout(auth.ParseToken(tokenString))

	c.JSON(http.StatusOK, gin.H{"status": status})
}

func (handler *User) Get(c *gin.Context) {
	offset := c.Query("offset")
	limit := c.Query("limit")

	id := c.Query("id")

	if id != "" {
		ID, _ := strconv.ParseUint(id, 10, 64)

		data, status := handler.service.FindByID(ID)

		c.JSON(http.StatusOK, gin.H{"data": data, "status": status})

		return
	}

	var param users.User

	name := c.Query("name")

	if name != "" {
		param.Name = name
	}

	datas, status := handler.service.Find(param, offset, limit)

	c.JSON(http.StatusOK, gin.H{"data": datas, "status": status})
}
