package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spadesk1991/easChat-server/modules"
	"gopkg.in/mgo.v2/bson"
)

func Register(ctx *gin.Context) {
	type body struct {
		UserName string `json:"user_name,omitempty"`
		Password string `json:"password,omitempty"`
	}
	req := body{}
	ctx.ShouldBind(&req)
	user := modules.User{
		UserName: req.UserName,
		Password: req.Password,
	}
	err := user.Create()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
}

func GetUser(ctx *gin.Context) {
	userName, _ := ctx.GetQuery("username")
	user, err := modules.GetUser(bson.M{"user_name": userName})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func GetUsers(ctx *gin.Context) {
	users, err := modules.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func Login(ctx *gin.Context) {
	type body struct {
		UserName string `json:"user_name,omitempty"`
		Password string `json:"password,omitempty"`
	}
	req := body{}
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
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "登录失败"})
	}
}
