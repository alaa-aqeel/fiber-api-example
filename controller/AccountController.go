package controller 


import (
	
	"github.com/gofiber/fiber"
	"github.com/alaaProg/restful"
	"github.com/alaaProg/postapi/models"
)

type AccountCtrl struct{
	restful.Controller

}


func (ctrl *AccountCtrl) Get(ctx *fiber.Ctx) {

	ctx.JSON(Map{

		"msg" : models.CurrentUser,
	})
}


func (ctrl *AccountCtrl) Update(ctx *fiber.Ctx) {

	ctx.BodyParser(&models.CurrentUser)

	if err := models.CurrentUser.Save(); err != nil {
		ctx.Status(401).JSON(Map{ "msg": err.Error(), })
		return 
	}

	ctx.JSON(Map{

		"data" : models.CurrentUser,
		"msg":"Successfuly update",
	})
}

func (ctrl *AccountCtrl) Delete(ctx *fiber.Ctx) {

	if err := models.CurrentUser.Delete(); err != nil{
		ctx.Status(401).JSON(Map{ "msg" : err.Error(),  })
		return 
	}

	ctx.JSON(Map{
		"data" : models.CurrentUser,
		"msg"  : "Successfuly Delete Account",

	})
}