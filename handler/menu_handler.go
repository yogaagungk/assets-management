package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yogaagungk/assets-management/services/menus"
)

type Menu struct {
	service *menus.Service
}

func ProvideMenu(service *menus.Service) Menu {
	return Menu{service}
}

func (handler *Menu) Post(c *gin.Context) {
	var entity menus.Menu

	c.BindJSON(&entity)

	status := handler.service.Save(entity)

	c.JSON(http.StatusOK, gin.H{"status": status})
}

func (handler *Menu) Put(c *gin.Context) {
	var entity menus.Menu

	c.Bind(&entity)

	status := handler.service.Update(entity)

	c.JSON(http.StatusOK, gin.H{"status": status})
}

func (handler *Menu) Get(c *gin.Context) {
	offset := c.Query("offset")
	limit := c.Query("limit")

	id := c.Query("id")

	if id != "" {
		ID, _ := strconv.ParseUint(id, 10, 64)

		data, status := handler.service.FindByID(ID)

		c.JSON(http.StatusOK, gin.H{"data": data, "status": status})

		return
	}

	var param menus.Menu

	name := c.Query("name")

	if name != "" {
		param.Name = name
	}

	datas, status := handler.service.Find(param, offset, limit)

	c.JSON(http.StatusOK, gin.H{"data": datas, "status": status})
}

func (handler *Menu) Delete(c *gin.Context) {
	paramId := c.Param("id")

	id, _ := strconv.ParseUint(paramId, 10, 64) // convert from string to uint64

	status := handler.service.Delete(id)

	c.JSON(http.StatusOK, gin.H{"status": status})
}
