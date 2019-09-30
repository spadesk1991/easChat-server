package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spadesk1991/easChat-server/modules"
	"gopkg.in/mgo.v2/bson"
)

// Register Register
// @Summary 注册用户
// @Description 注册用户
// @Produce json
// @Param body body controllers.LoginParams true "body参数"
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "retcode：10002 参数错误； retcode：10003 校验错误"
// @Failure 401 {string} string "retcode：10001 登录失败"
// @Failure 500 {string} string "retcode：20001 服务错误；retcode：20002 接口错误；retcode：20003 无数据错误；retcode：20004 数据库异常；retcode：20005 缓存异常"
// @Router /user/person/register [post]
func Register(ctx *gin.Context) {
	req := LoginParams{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user := modules.User{
		UserName: req.UserName,
		Password: req.Password,
	}
	err := user.Create()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

// GetUser 查询用户信息
// @Summary 查询用户信息
// @Description 查询数据库用户信息
// @Produce json
// @Param id query string true "mongodb文档id"
// @Success 200 {object} modules.User "返回用户信息"
// @Failure 400  "retcode：10002 参数错误； retcode：10003 校验错误"
// @Failure 401  "retcode：10001 登录失败"
// @Failure 500  "retcode：20001 服务错误；retcode：20002 接口错误；retcode：20003 无数据错误；retcode：20004 数据库异常；retcode：20005 缓存异常"
// @Router /user/person/{id} [get]
func GetUser(ctx *gin.Context) {
	userName, _ := ctx.GetQuery("username")
	id, _ := ctx.GetQuery("id")
	query := bson.M{}
	if userName != "" {
		query["user_name"] = userName
	}
	if id != "" {
		query["_id"] = bson.ObjectIdHex(id)
	}
	user, err := modules.GetUser(query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// GetUsers 查询用户信息列表
// @Summary 查询用户信息列表
// @Description 查询数据库用户信息
// @Produce json
// @Success 200 {array} modules.User "返回用户信息"
// @Failure 400 {string} string "retcode：10002 参数错误； retcode：10003 校验错误"
// @Failure 401 {string} string "retcode：10001 登录失败"
// @Failure 500 {string} string "retcode：20001 服务错误；retcode：20002 接口错误；retcode：20003 无数据错误；retcode：20004 数据库异常；retcode：20005 缓存异常"
// @Router /user/person [get]
func GetUsers(ctx *gin.Context) {
	users, err := modules.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

type LoginParams struct {
	UserName string `json:"user_name,omitempty" binding:"required" example:"张三"`
	Password string `json:"password,omitempty" binding:"required" example:"123456"`
}

// Login 登录
// @Summary 登录
// @Description 登录
// @Produce json
// @Param body body controllers.LoginParams true "body参数"e
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "retcode：10002 参数错误； retcode：10003 校验错误"
// @Failure 401 {string} string "retcode：10001 登录失败"
// @Failure 500 {string} string "retcode：20001 服务错误；retcode：20002 接口错误；retcode：20003 无数据错误；retcode：20004 数据库异常；retcode：20005 缓存异常"
// @Router /user/person/login [post]
func Login(ctx *gin.Context) {
	req := LoginParams{}
	ctx.ShouldBind(&req)
	user := modules.User{
		UserName: req.UserName,
		Password: req.Password,
	}
	user, err := modules.GetUser(bson.M{"user_name": req.UserName, "password": req.Password})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if user.ID.Hex() != "" {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusBadRequest, "ok")
	}
}
