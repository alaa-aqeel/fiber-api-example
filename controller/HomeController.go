package controller 



import (

	"github.com/gofiber/fiber"
	"github.com/alaaProg/restful"
	"github.com/alaaProg/postapi/models"
)


type HomeCtrl struct{
	restful.Controller
}


func (ctrl *HomeCtrl) Get(ctx *fiber.Ctx) {

	ctx.JSON(Map{
		
		"msg" : models.CurrentUser,
	})
}