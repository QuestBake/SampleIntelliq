package controller

import (
	"net/http"
	"pracSpace/restHandler_Gin/app/model"
	"pracSpace/restHandler_Gin/app/service"

	"github.com/gin-gonic/gin"
)

//TestHandler test
func TestHandler(ctx *gin.Context) {
	ctx.JSON(200, "Success !!")
}

//FindAllUsers find all
func FindAllUsers(ctx *gin.Context) {
	users, err := service.FindAllUsers()
	if err != nil {
		ctx.JSON(http.StatusOK, &model.AppResponse{Status: 400, Body: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &model.AppResponse{Status: 200, Body: users, Msg: ""})
}

//AuthenticateUser login user
func AuthenticateUser(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, &model.AppResponse{Status: 400, Body: nil, Msg: "Invalid Argument"})
		return
	}
	loggedUser, err := service.AuthenticateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, &model.AppResponse{Status: 400, Body: nil, Msg: err.Error()})
		return
	}
	address, err := service.FindAddressByID(loggedUser.Address.ID)
	if err != nil {
		address = nil
	}
	loggedUser.Address = *address
	ctx.JSON(http.StatusOK, &model.AppResponse{Status: 200, Body: loggedUser, Msg: ""})
}

//SearchUsers index based search
func SearchUsers(ctx *gin.Context) {
	term := ctx.Param("term")
	users, err := service.SearchUsersWithSearchIndex(term)
	if err != nil {
		ctx.JSON(http.StatusOK, &model.AppResponse{Status: 400, Body: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &model.AppResponse{Status: 200, Body: users, Msg: ""})
}
