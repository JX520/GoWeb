package article

import (
	"Goweb/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)


// @Summary 提交新的文章内容
// @Id 1
// @Tags 文章
// @version 1.0
// @Accept application/x-json-stream
// @Param article body model.Article true "文章"
// @Success 200 object model.Result 成功后返回值
// @Failure 409 object model.Result 添加失败
// @Router /article [post]
func Insert(ctx *gin.Context)  {
	article := model.Article{}
	var id = -1
	var msg = "添加数据失败"
	if e := ctx.ShouldBindJSON(&article); e == nil{
		id = article.Insert() // 去数据库执行插入方法，并返回id
		msg = "添加数据成功"
	}
	result := model.Result{
		Code:    http.StatusOK,
		Message: msg,
		Data:    gin.H{
			"id":id,
		},
	}
	ctx.JSON(http.StatusOK, result)
}


// @Summary 通过文章 id 获取单个文章内容
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200 object model.Result 成功后返回值
// @Router /article/{id} [get]
func GetOne(ctx *gin.Context)  {
	id := ctx.Param("id")
	i, e := strconv.Atoi(id) //将字符串转换为int
	if e != nil {
		log.Panicln("转换int失败", e.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result":model.Result{
				Code:    http.StatusBadRequest,
				Message: "转换int失败",
				Data:    e.Error(),
			},
		})
	}
	article := model.Article{Id: i}
	art := article.FindById() // 通过id查找数据
	ctx.JSON(http.StatusOK, gin.H{
		"result":model.Result{
			Code:    http.StatusOK,
			Message: "单个数据查询成功",
			Data:    art,
		},
	})
}

// @Summary 获取所有的文章
// @version 1.0
// @Accept application/x-json-stream
// @Success 200 object model.Result 成功后返回值
// @Router /articles [get]
func GetAll(ctx *gin.Context)  {
	article := model.Article{}
	articles := article.FindAll()
	ctx.JSON(http.StatusOK, gin.H{
		"result": model.Result{
			Code:    http.StatusOK,
			Message: "查询所有数据成功",
			Data:    articles,
		},
	})
}

// @Summary 通过id删除指定文章类型
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200  object model.Result 成功后返回值
// @Router /article/{id} [delete]
func DeleteOne(ctx *gin.Context)  {
	id := ctx.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("转换id失败", e.Error())
	}
	article := model.Article{Id: i}
	article.DeleteOne()
}