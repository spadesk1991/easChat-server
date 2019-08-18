package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spadesk1991/easChat-server/modules"
)

func Register(ctx *gin.Context) {
	type body struct {
		UserName string `json:"user_name,omitempty"`
		Password string `json:"password,omitempty"`
	}
	req := body{}
	ctx.ShouldBind(req)
	user := modules.User{
		UserName: req.UserName,
		Password: req.Password,
	}
	err := user.Create()
	if err != nil {
		ctx.JSON(http.StatusGone, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusGone, gin.H{"msg": "注册成功"})
}
