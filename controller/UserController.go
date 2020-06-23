package controller

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/alaaProg/restful"
	"github.com/alaaProg/postapi/models"
)

type UserCtrl struct{

	restful.Controller
}


func (ctrl *UserCtrl) Show(ctx *fiber.Ctx) {

	id := ctx.Params("id")
	user, err := new(models.UserModel).Get(id)
	
	if err != nil {
		ctx.Status(404).JSON( Map{ 

			"msg" : fmt.Sprintf("NotFound User %s", id), 
		})
		return 
	}

	ctx.JSON(Map{ "data": user, })
}

func (ctrl *UserCtrl) Get(ctx *fiber.Ctx) {

	ctx.JSON(Map{
		"data": new(models.UserModel).All(),
	})
}

func (ctrl *UserCtrl) Update(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	user, err := new(models.UserModel).Get(id)
	
	if err != nil {
		ctx.Status(404).JSON( Map{ 

			"msg" : fmt.Sprintf("NotFound User %s", id), 
		})
		return 
	}

	ctx.BodyParser(&user)
	if err := user.Save(); err != nil {
		ctx.Status(401).JSON(Map{ "msg":err.Error(), })
		return 
	}

	ctx.JSON(Map{ 
		"msg" : "Successfuly Update User ",
		"data": user, 
	})
}

func (ctrl *UserCtrl) Delete(ctx *fiber.Ctx) {

	id := ctx.Params("id")
	user, err := new(models.UserModel).Get(id)
	
	if err != nil {
		ctx.Status(404).JSON( Map{ 
			"msg" : fmt.Sprintf("NotFound User %s", id), 
		})
		return 
	}

	if err := user.Delete(); err != nil{
		ctx.Status(401).JSON(Map{ "msg" : err.Error(),  })
		return 
	}

	ctx.JSON(Map{
		"msg": fmt.Sprintf("Successfuly Delete User %s", id),
	})
}

func (ctrl *UserCtrl) Post(ctx *fiber.Ctx) {

	user := new(models.UserModel)
	ctx.BodyParser(user)

	if messages, err := ctrl.Valid(user); err{
		
		ctx.Status(401).JSON(messages)
		return
	}

	if err := user.Create(); err != nil {
		ctx.Status(401).JSON(Map{
			"msg": err.Error(), 
		})
		return 
	}

	ctx.JSON(Map{
		"msg" : "Sucessfuly Create User" ,
		"data" : user,
	})
}
