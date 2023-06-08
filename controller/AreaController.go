package controller

import (
	"fitness/dao"
	"fitness/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArea(c *gin.Context) {
	var area dao.Area
	if err := c.ShouldBind(&area); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "should bind error",
		})
		return
	}
	if service.IsExistsArea(area.AreaId) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1001,
			"msg":  "场地编码不可以重复",
		})
		return
	} else {
		area.State = "空闲中"
		if err := service.CreateArea(&area); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "service error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":     "create area success",
			"state":   "0",
			"code":    "200",
			"local":   area.AreaLocal,
			"area_id": area.AreaId,
		})
	}

}

func DeleteArea(c *gin.Context) {
	var area dao.Area
	if err := c.BindJSON(&area); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "should bind error",
		})
		return
	}
	println("area_id :", area.AreaId)
	if err := service.DeleteArea(area.AreaId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "del area service error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "del area success",
		"id":  area.AreaName,
	})
}

func UpdateArea(c *gin.Context) {
	var area dao.Area
	if err := c.ShouldBindJSON(&area); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "should bind error",
		})
		return
	}
	if err := service.UpdateArea(&area); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1011,
			"msg":  "更新场地信息失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "更新场地信息成功",
			"code": 200,
		})
		return
	}
}

func GetAreaList(c *gin.Context) {

	total, arealist := service.GetAreaList()
	if total == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1010,
			"msg":  "未查询到数据",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": arealist,
		})
	}
}

func GetAreaById(c *gin.Context) {
	var area dao.Area
	if err := c.BindJSON(&area); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if todoList, err := service.GetAreaById(strconv.FormatUint(uint64(area.ID), 10)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "service error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "get area success",
			"code": 200,
			"data": todoList,
		})
		return
	}
}
func GetAreaByAreaId(c *gin.Context) {
	var area dao.Area
	if err := c.BindJSON(&area); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if todoList, err := service.GetAreaByAreaId(area.AreaId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "查询失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "get area success",
			"code": 200,
			"data": todoList,
		})
		return
	}
}

func GetArea(c *gin.Context) {
	var area dao.Area
	if err := c.BindJSON(&area); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	println("name:", area.AreaName)
	println("name:", area.AreaId)
	println("name:", area.Type)
	if todoList, err := service.GetArea(area.Type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "service error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "get area success",
			"code": 200,
			"data": todoList,
		})
		return
	}
}
