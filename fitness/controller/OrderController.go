package controller

import (
	"fitness/dao"
	"fitness/service"
	"fitness/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 修改订单信息
func UpdateOrder(c *gin.Context) {
	var order dao.Order
	err := c.ShouldBind(&order)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "bind json error",
		})
		return
	}
	tx := service.UpdateOrder(&order)
	if tx.RowsAffected != 0 {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "更新订单信息成功",
		})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "更新订单信息失败",
		})
		return
	}
}

// 获取订单列表
func GetOrderList(c *gin.Context) {

	total, order := service.GetOrderList()
	if total == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1010,
			"msg":  "未查询到数据",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": order,
		})
	}
}

// 修改订单状态
func UpdateOrderState(c *gin.Context) {
	total, order := service.GetOrderList()
	println("total", total)
	for i := 0; i < total; i++ {
		if order[i].BookingStatus == "预约成功" {
			if utils.ComPareTime(order[i].EndTime) {
				println(order[i].BookingId)
				service.UpdateOrderStatus(order[i].BookingId, "预约结束")
				service.UpdateAreaState(order[i].AreaId, "空闲中")
				service.UpdateOrderState(order[i].BookingId, "空闲中")
			} else {
				println("nihao ", order[i].BookingId)
			}
		}
	}
}
func UpdateOrderStatus(c *gin.Context) {
	total, order := service.GetOrderList()
	println("total", total)
	for i := 0; i < total; i++ {
		if order[i].BookingStatus == "预约成功" {
			if utils.ComPareTime(order[i].StartTime) && !utils.ComPareTime(order[i].EndTime) {
				println("开始时间", order[i].StartTime)
				service.UpdateAreaState(order[i].AreaId, "使用中")
				service.UpdateOrderState(order[i].BookingId, "使用中")
			} else {
				println("cuowu", order[i].BookingId)
			}
		}
	}
}

// 确认订单状态
func OkOrderState(c *gin.Context) {
	var order dao.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "bind json error",
		})
		return
	}

	getOrder := service.GetOrder(order.BookingId)
	if getOrder.BookingStatus == "待确定" {
		service.UpdateOrderStatus(order.BookingId, "预约成功")
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "确认成功",
		})
		return
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "订单已过期或订单已经预约成功",
		})
		return
	}

}

// 添加订单
func AddOrder(c *gin.Context) {
	var order dao.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(400, gin.H{"msg": "bind json error"})
		return
	}
	mobileAny := c.MustGet("mobile")
	mobile := mobileAny.(string)
	user := service.GetUser(mobile)

	//判断场地状态 是否为空闲中
	if checkOrder, _ := service.CheckOrder(&order); !checkOrder {
		c.JSON(400, gin.H{"msg": "该场地此时间段已被预约"})
	} else {
		//service添加预约
		a, _ := service.GetAreaByAreaId(order.AreaId)
		if utils.ComPareTime(order.StartTime) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 1018,
				"msg":  "预约时间错误~",
			})
			return
		} else if utils.ComPareTimeWith(order.StartTime, order.EndTime) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 1015,
				"msg":  "预约时间错误",
			})
			return
		} else {
			order.BookingId = utils.RandInt(12)
			order.AreaId = a.AreaId
			order.AreaName = a.AreaName
			order.AreaLocal = a.AreaLocal
			order.AreaDesc = a.AreaDesc
			order.AreaFee = a.AreaFee
			order.State = a.State
			order.UserId = user.UserId
			if err := service.AddOrder(&order); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"code": 1019,
					"msg":  "预约失败",
				})
				return
			} else {
				service.UpdateOrderStatus(order.BookingId, "待确定")
				service.UpdateAreaState(order.AreaId, "预约中")
				service.UpdateOrderState(order.BookingId, "预约中")
				c.JSON(http.StatusOK, gin.H{
					"code": 200,
					"msg":  "添加预约成功",
				})
				return
			}
		}

	}
}

// 删除订单
func DeleteOrder(c *gin.Context) {
	var order dao.Order
	if err := c.ShouldBind(&order); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "bind json error",
		})
		return
	}
	//查询该订单id 获取订单信息 area_id booking_id
	getOrder := service.GetOrder(order.BookingId)
	if getOrder.BookingId == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "该订单不存在",
		})
		return
	}
	if err := service.DeleteOrder(order.BookingId); err != nil {
		c.JSON(400, gin.H{
			"code": 1020,
			"msg":  "删除订单失败",
		})
		return
	} else {
		//删除成功 将场地状态改为空闲
		if getOrder.BookingStatus != "预约结束" {
			service.UpdateAreaState(getOrder.AreaId, "空闲中")
			service.UpdateOrderState(getOrder.BookingId, "空闲中")
		}
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "删除订单成功",
		})
		return
	}

}

// 取消订单
func CancelOrder(c *gin.Context) {
	var order dao.Order
	if err := c.ShouldBind(&order); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "bind json error",
		})
		return
	}
	//查询该订单id 获取订单信息 area_id booking_id
	getOrder := service.GetOrder(order.BookingId)
	if getOrder.BookingId == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "该订单不存在",
		})
		return
	}
	if !utils.ComPareTime(order.StartTime) {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "您的预约已经开始，该订单不能取消",
		})
		return
	} else {
		if err := service.DeleteOrder(order.BookingId); err != nil {
			c.JSON(400, gin.H{
				"code": 1020,
				"msg":  "取消订单失败",
			})
			return
		} else {
			//操作成功 将场地状态改为空闲
			service.UpdateAreaState(getOrder.AreaId, "空闲中")
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "取消订单成功",
			})
			return
		}
	}
}

// 根据订单号查询
func GetOrder(c *gin.Context) {
	var order dao.Order
	err := c.ShouldBind(&order)
	if err != nil {
		return
	}

	println("订单号", order.BookingId, "你好")
	getOrder := service.GetOrder(order.BookingId)
	c.JSON(200, gin.H{
		"code": 200,
		"data": getOrder,
	})
	return
}

// 根据场地id查询
func GetOrderByAreaId(c *gin.Context) {
	var order dao.Order
	err := c.ShouldBind(&order)
	if err != nil {
		return
	}

	println(order.AreaId, "你好")
	getOrder := service.GetOrderByAreaId(order.AreaId)
	c.JSON(200, gin.H{
		"code": 200,
		"data": getOrder,
		"id":   order.AreaId,
	})
	return
}

// 查询订单
func GetOrderByMobile(c *gin.Context) {
	mobileAny := c.MustGet("mobile")
	mobile := mobileAny.(string)
	user := service.GetUser(mobile)
	total, order := service.GetOrderByUserId(user.UserId)
	c.JSON(200, gin.H{
		"code":  200,
		"data":  order,
		"total": total,
	})
	return
}
