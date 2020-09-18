package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		fmt.Printf("%s::%s \t %s \t %s", time.Now().Format("2020-09-18 16:16:16"), host, url, method)
		context.Next()
		fmt.Println(context.Writer.Status())
	}
}

// 是否登录的验证
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookie, e := context.Request.Cookie("user_cookie") // 获取cookie
		if e == nil {
			println("已经登录")
			// 每次请求都刷新cookie
			println(cookie.Name)
			context.SetCookie(cookie.Name, cookie.Value, 1000, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
			context.Next()
		}else {
			println("未登录，请先登录")
			context.Abort() // 表示对当前的请求进行终止
			context.HTML(http.StatusUnauthorized, "401.tmpl", nil)
		}

	}
}

