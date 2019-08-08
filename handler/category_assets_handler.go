package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yogaagungk/assets-management/services/category_assets"
)

type CategoryAsset struct {
	service *category_assets.Service
}

func ProvideCategoryAsset(service *category_assets.Service) CategoryAsset {
	return CategoryAsset{service}
}

func (handler *CategoryAsset) Post(c *gin.Context) {
	var entity category_assets.CategoryAsset

	c.BindJSON(&entity)

	status := handler.service.Save(entity)

	c.JSON(http.StatusOK, gin.H{"status": status})
}

func (handler *CategoryAsset) Put(c *gin.Context) {
	var entity category_assets.CategoryAsset

	c.Bind(&entity)

	status := handler.service.Update(entity)

	c.JSON(http.StatusOK, gin.H{"status": status})
}

func (handler *CategoryAsset) Get(c *gin.Context) {
	offset := c.Query("offset")
	limit := c.Query("limit")

	id := c.Query("id")

	if id != "" {
		ID, _ := strconv.ParseUint(id, 10, 64)

		data, status := handler.service.FindByID(ID)

		c.JSON(http.StatusOK, gin.H{"data": data, "status": status})

		return
	}

	var param category_assets.CategoryAsset

	name := c.Query("name")

	if name != "" {
		param.Name = name
	}

	datas, status := handler.service.Find(param, offset, limit)

	c.JSON(http.StatusOK, gin.H{"data": datas, "status": status})
}

func (handler *CategoryAsset) Delete(c *gin.Context) {
	paramId := c.Param("id")

	id, _ := strconv.ParseUint(paramId, 10, 64) // convert from string to uint64

	status := handler.service.Delete(id)

	c.JSON(http.StatusOK, gin.H{"status": status})
}
