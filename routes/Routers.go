package routes

import (
	"fitness/controller"
	"fitness/middleware"
	"fitness/middleware/errorHandle"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(errorHandle.AppErrorHandleMiddleware())

	//管理员 登录 注册 退出登录
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/logout", controller.Logout)

	//用户组
	UserGroup := r.Group("user")
	{
		//用户登录
		UserGroup.POST("/login", controller.UserLogin)
		//用户创建（admin）
		UserGroup.POST("/create", controller.CreateUser)
		//用户注册
		UserGroup.POST("/register", controller.RegisterUser)
		//获取用户列表
		UserGroup.POST("/getuserinfo", controller.UserInfo)
		//查询用户
		UserGroup.POST("/getuser", controller.GetUser)
		//删除用户
		UserGroup.POST("/deluserbymobile", controller.DeleteUserByMobile)
		//修改用户头像
		UserGroup.POST("/editavatar", controller.EditUserAvatar)
		//退出登录
		UserGroup.POST("/user/logout", controller.UserLogOut)
		//修改用户信息
		UserGroup.POST("/updateuser", controller.Update)
	}

	UserGroup = r.Group("user")
	UserGroup.Use(middleware.JWT())
	{
		//获取用户登录信息
		UserGroup.GET("/info", controller.User)
		//用户更新
		UserGroup.POST("/update", controller.Update)
		//修改密码
		UserGroup.POST("/repassword", controller.UpdatePassword)
		//更换头像
		UserGroup.POST("/upavatar", controller.UploadUserAvatar)
		//上传到七牛云
		r.POST("/upload", controller.UpLoad)
	}

	//教练组
	CoachGroup := r.Group("coach")
	{
		CoachGroup.DELETE("/delcoachbyid", controller.DeleteCoachById)
		CoachGroup.POST("/delcoachbymobile", controller.DeleteCoachByMobile)
		CoachGroup.POST("/update", controller.UpdateCoach)
		CoachGroup.GET("/getcoachbyid", controller.GetCoachById)
		CoachGroup.POST("/getcoachbymobile", controller.GetCoachByMobile)
		CoachGroup.GET("/getallcoachs", controller.GetAllCoach)
		CoachGroup.POST("/getcoachinfo", controller.CoachInfo)
		CoachGroup.POST("/create", controller.CreateCoach)
		CoachGroup.POST("/login", controller.CoachLogin)
	}

	//场地组
	AreaGroup := r.Group("area")
	{
		AreaGroup.POST("/create", controller.CreateArea)
		AreaGroup.POST("/update", controller.UpdateArea)

		//删除场地
		AreaGroup.POST("/delarea", controller.DeleteArea)
		//获取场地列表
		AreaGroup.POST("/list", controller.GetAreaList)
		//查询场地
		AreaGroup.POST("/get_area", controller.GetAreaByAreaId)
		//查询场地 根据场地类型
		AreaGroup.POST("/getarea", controller.GetArea)

	}

	OrderGroup := r.Group("order")
	OrderGroup.Use(middleware.JWT())
	{
		//订单列表
		OrderGroup.POST("/list", controller.GetOrderList)
		//更新订单信息
		OrderGroup.POST("/update", controller.UpdateOrder)
		//更新订单状态
		OrderGroup.POST("/updateorderstate", controller.UpdateOrderState)
		OrderGroup.POST("/updateorderstatus", controller.UpdateOrderStatus)
		//确认订单状态
		OrderGroup.POST("/ok", controller.OkOrderState)
		//添加订单
		OrderGroup.POST("/add", controller.AddOrder)
		//删除订单
		OrderGroup.POST("/delete", controller.DeleteOrder)
		//取消订单
		OrderGroup.POST("/cancel", controller.CancelOrder)
		//根据订单号查询
		OrderGroup.POST("/get", controller.GetOrder)
		//根据areaid查询
		OrderGroup.POST("/getbyareaid", controller.GetOrderByAreaId)
		//根据登录用户查询

		OrderGroup.POST("/getorder", controller.GetOrderByMobile)
	}
	return r
}
