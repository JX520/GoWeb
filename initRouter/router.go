package initRouter

import (
	"Goweb/handler"
	"Goweb/middleware"
	"Goweb/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter() *gin.Engine {
	router := gin.Default()

	//全局中间件的使用
	//router := gin.New()
	//router.Use(middleware.Logger(), gin.Recovery())

	//加载templates的html模板
	//if mode := gin.Mode(); mode == gin.TestMode {
	//	router.LoadHTMLGlob("./../templates/*")
	//}else {
		router.LoadHTMLGlob("templates/*")
	//}
	router.Static("/statics","./statics")
	router.StaticFS("/avatar", http.Dir(utils.RootPath() + "avatar/"))
	// 请求"/"
	index := router.Group("/")
	{
		//请求返回单个字符串
		//index.Any("",commonMethod)

		//请求返回整个html
		index.Any("",handler.Index)
	}

	// 请求"/user"
	userRouter := router.Group("/user")
	{
		//userRouter.GET("/:name",handler.UserSave)
		userRouter.GET("",handler.UserSaveByQuery)
		userRouter.POST("/register",handler.UserRegister)
		userRouter.POST("/login",handler.UserLogin)
		userRouter.GET("/profile/",middleware.Auth(), handler.UserProfile)
		userRouter.POST("/update",middleware.Auth(), handler.UpdateUserProfile)
	}

	//请求"/register"
	//registerRouter := router.Group("/register")
	//{
	//	registerRouter.POST("/", handler.UserRegister)
	//}
	return router
}

//func commonMethod(ctx *gin.Context)  {
//	ctx.String(http.StatusOK,"hello jian" + strings.ToLower(ctx.Request.Method) + "method")
//}

