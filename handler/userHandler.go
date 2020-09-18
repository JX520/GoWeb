package handler

import (
	"Goweb/model"
	"Goweb/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func UserSave(ctx *gin.Context) {
	username := ctx.Param("name")
	ctx.String(http.StatusOK,"用户" + username + "已经保存")
}

func UserSaveByQuery(ctx *gin.Context)  {
	username := ctx.Query("name")
	age := ctx.Query("age")
	ctx.String(http.StatusOK,"用户:" + username + ",年龄:" + age + "已经保存")
}

//用户注册到Mysql
func UserRegister(ctx *gin.Context)  {
	var user model.UserModel
	if err := ctx.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		ctx.String(http.StatusBadRequest,"输入的数据不合法")
	}else {
		id := user.Save()
		log.Println("新增用户id为：", id)
		//ctx.Redirect(http.StatusMovedPermanently, "/")
		//ctx.String(http.StatusOK, "email：" + user.Email + "，密码：" + user.Pwd + "，二次密码：" + user.PwdAgain)
	}
}

//用户登陆
func UserLogin(ctx *gin.Context)  {
	var user model.UserModel
	if e := ctx.Bind(&user); e != nil{
		log.Panicln("邮箱不正确！", e.Error())
	}
	u := user.QueryUser()
	if u.Pwd == user.Pwd{
		log.Println("登陆成功：", u.Email)
		// 设置cookie
		ctx.SetCookie("user_cookie", strconv.Itoa(u.Id), 1000, "/", "localhost", false, true)
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email":u.Email,
			"id":u.Id,
		})
		//ctx.Redirect(http.StatusMovedPermanently, "/")
	}else {
		log.Println("密码不正确！")
	}
}

//用户查找成功后跳转个人信息页
func UserProfile(ctx *gin.Context)  {
	id := ctx.Query("id")
	var user model.UserModel
	i, err := strconv.Atoi(id) // 将String转化为int
	u, e := user.QueryById(i)
	if e != nil || err != nil{
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error":e,
		})
	}
	ctx.HTML(http.StatusOK, "user_profile.tmpl", gin.H{
		"user":u,
	})
}

//更新用户头像
func UpdateUserProfile(ctx *gin.Context)  {
	var user model.UserModel
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error":err.Error(),
		})
		log.Panicln("绑定用户错误",err.Error())
	}
	file, e := ctx.FormFile("avatar-file")
	if e != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error":e,
		})
		log.Panicln("文件上传错误", e.Error())
	}
	path := utils.RootPath()
	path = filepath.Join(path, "avatar")
	fmt.Println("path =>", path)
	e = os.MkdirAll(path, os.ModePerm)
	if e != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error":e,
		})
		log.Panicln("无法创建文件夹", e.Error())
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	e = ctx.SaveUploadedFile(file, filepath.Join(path, fileName))
	if e != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法保存文件", e.Error())
	}
	avatarUrl := "/avatar/" + fileName
	user.Avatar = sql.NullString{String: avatarUrl}
	e = user.Update(user.Id + 1)
	if e != nil {
		ctx.HTML(http.StatusOK,"error.tmpl", gin.H{
			"error":e,
		})
		log.Panicln("数据无法更新", e.Error())
	}

	ctx.Redirect(http.StatusMovedPermanently, "/user/profile?id=" + strconv.Itoa(user.Id+1) )
}