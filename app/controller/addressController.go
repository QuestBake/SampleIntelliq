package controller

import (
	"net/http"
	"pracSpace/restHandler_Gin/app/model"
	"pracSpace/restHandler_Gin/app/service"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

//AddAddress all
func AddAddress(ctx *gin.Context) {
	var address model.Address
	ctx.BindJSON(&address)
	res, err := service.AddAddress(&address)
	if err != nil {
		ctx.JSON(http.StatusOK, &model.AppResponse{Status: 400, Body: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &model.AppResponse{Status: 200, Body: res, Msg: ""})
}

//UpdateAddress all
func UpdateAddress(ctx *gin.Context) {
	var address model.Address
	ctx.BindJSON(&address)
	res, err := service.UpdateAddress(&address)
	if err != nil {
		ctx.JSON(http.StatusOK, &model.AppResponse{Status: 400, Body: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &model.AppResponse{Status: 200, Body: res, Msg: ""})
}

//RemoveAddress all
func RemoveAddress(ctx *gin.Context) {
	addressID := ctx.Param("address_id")
	res, err := service.RemoveAddress(bson.ObjectIdHex(addressID))
	if err != nil {
		ctx.JSON(http.StatusOK, &model.AppResponse{Status: 400, Body: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &model.AppResponse{Status: 200, Body: res, Msg: ""})
}

//FindAddressByCity all
func FindAddressByCity(ctx *gin.Context) {
	city := ctx.Param("city")
	address, err := service.FindAddressByCity(city)
	if err != nil {
		ctx.JSON(http.StatusOK, &model.AppResponse{Status: 400, Body: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &model.AppResponse{Status: 200, Body: address, Msg: ""})
}

//AllAddresses all
func AllAddresses(ctx *gin.Context) {
	addresses, err := service.FindAllAddresses()
	if err != nil {
		ctx.JSON(http.StatusOK, &model.AppResponse{Status: 400, Body: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &model.AppResponse{Status: 200, Body: addresses, Msg: ""})
}

//FindAddressesByState allfind by state
func FindAddressesByState(ctx *gin.Context) {
	city := ctx.Param("state")
	addresses, err := service.FindAddressByState(city)
	if err != nil {
		ctx.JSON(http.StatusOK, &model.AppResponse{Status: 400, Body: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &model.AppResponse{Status: 200, Body: addresses, Msg: ""})
}
