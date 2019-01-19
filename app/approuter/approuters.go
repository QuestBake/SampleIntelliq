package approuter

import (
	"pracSpace/restHandler_Gin/app/controller"

	"github.com/gin-gonic/gin"
)

var mrouter *gin.Engine

//AddRouters adding routes
func AddRouters(router *gin.Engine) {
	mrouter = router
	mrouter.GET("/", controller.TestHandler)
	addUserRouters()
	addAddressRouters()
}

func addUserRouters() {
	userRoutes := mrouter.Group("/user")
	{
		userRoutes.GET("/all", controller.FindAllUsers)
		userRoutes.POST("/login", controller.AuthenticateUser)
		userRoutes.GET("/search/:term", controller.SearchUsers)
	}
}

func addAddressRouters() {
	addrRoutes := mrouter.Group("/address")
	{
		addrRoutes.GET("/all", controller.AllAddresses)
		addrRoutes.POST("/add", controller.AddAddress)
		addrRoutes.PUT("/update", controller.UpdateAddress)
		addrRoutes.DELETE("/remove/:address_id", controller.RemoveAddress)
		addrRoutes.GET("/find/city/:city", controller.FindAddressByCity)
		addrRoutes.GET("/find/state/:state", controller.FindAddressesByState)
	}
}
