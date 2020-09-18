//单元测试

package test

import (
	"Goweb/initRouter"
	"bytes"
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

var router = initRouter.SetRouter()

func TestUserSave(t *testing.T)  {
	username := "jian"
	//router := initRouter.SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet,"/user/" + username,nil)
	router.ServeHTTP(w,req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户" + username + "已经保存", w.Body.String())
}

func TestUserSaveQuery(t *testing.T)  {
	username := "jian"
	age := 24
	//router := initRouter.SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/?name=" + username + "&age=" + strconv.Itoa(age), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户:" + username + ",年龄:" + strconv.Itoa(age) + "已经保存", w.Body.String())
}

func TestIndexHtml(t *testing.T)  {
	//router := initRouter.SetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserPostForm(t *testing.T)  {
	value := url.Values{}
	value.Add("email", "jian@gmail.com")
	value.Add("pwd", "123")
	value.Add("pwdAgain", "123")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type","application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusMovedPermanently, w.Code)
}

func TestUserPostFormError(t *testing.T)  {
	value := url.Values{}
	value.Add("email", "jian")
	value.Add("pwd", "123")
	value.Add("pwdAgain", "1234")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type","application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

//登陆测试
func TestUserLogin(t *testing.T)  {
	email := "jian@gmail.com"
	value := url.Values{}
	value.Add("email",email)
	value.Add("pwd","12345")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type","application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, strings.Contains(w.Body.String(),email), true)
}