package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yogaagungk/assets-management/services/roles"
)

type Role struct {
	service *roles.Service
}

func ProvideRole(service *roles.Service) Role {
	return Role{service}
}

func (handler *Role) Get(c *gin.Context) {
	offset := c.DefaultQuery("offset", "0")
	limit := c.DefaultQuery("limit", "10")

	id := c.Query("id")
	name := c.Query("name")

	if id != "" {
		ID, _ := strconv.ParseUint(id, 10, 64)

		data, status := handler.service.FindByID(ID)

		c.JSON(http.StatusOK, gin.H{"data": data, "status": status})

		return
	}

	var param roles.Role

	if name != "" {
		param.Name = name
	}

	datas, status := handler.service.Find(param, offset, limit)

	c.JSON(http.StatusOK, gin.H{"data": datas, "status": status})
}
