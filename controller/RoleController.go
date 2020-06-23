package controller


import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/alaaProg/restful"
	"github.com/alaaProg/postapi/models"
)

type RoleCtrl struct{
	restful.Controller
}


func (ctrl *RoleCtrl) DeleteRoleFromUser(ctx *fiber.Ctx){


	role, err := new(models.RoleModel).Get(ctx.Params("roleid"))
	if err != nil {
		ctx.Status(404).JSON( Map{ 
			"msg" : fmt.Sprintf("NotFound Role %s", ctx.Params("roleid")), 
		})
		return 
	}


	user, err := role.GetUser(ctx.FormValue("user"))
	if err != nil {
		ctx.Status(401).JSON( Map{ 
			"msg" : fmt.Sprintf("User %s haven't role:%s", ctx.FormValue("user"), role.Name), 
		})
		return 
	}

	role.DeleteUser(user)
	ctx.JSON(Map{
		"msg": fmt.Sprintf("Successfuly Delete role:%s from user:%s", role.Name, user.Username),
	})
}

func (ctrl *RoleCtrl) AppendRoleToUser(ctx *fiber.Ctx){

	user, err := new(models.UserModel).Get(ctx.FormValue("user")) 
	if err != nil {
		ctx.Status(404).JSON( Map{ 
			"msg" : fmt.Sprintf("NotFound User %s", ctx.FormValue("user")), 
		})
		return 
	}

	role, err := new(models.RoleModel).Get(ctx.Params("roleid"))
	if err != nil {
		ctx.Status(404).JSON( Map{ 
			"msg" : fmt.Sprintf("NotFound Role %s", ctx.Params("roleid")), 
		})
		return 
	}

	role.AppendUser(user)

	ctx.JSON(Map{
		"msg": fmt.Sprintf("Successfuly Set role:%s to user:%s", role.Name, user.Username),
	})
}

func (ctrl *RoleCtrl) Show(ctx *fiber.Ctx) {

	id := ctx.Params("id")
	role, err := new(models.RoleModel).Get(id)
	
	if err != nil {
		ctx.Status(404).JSON( Map{ 
			"msg" : fmt.Sprintf("NotFound Role %s", id), 
		})
		return 
	}


	ctx.JSON(Map{ 
		"data": role,
	})
}

func (ctrl *RoleCtrl) Get(ctx *fiber.Ctx) {

	ctx.JSON(Map{

		"data": new(models.RoleModel).All(),
	})
}

func (ctrl *RoleCtrl) Post(ctx *fiber.Ctx) {

	role := new(models.RoleModel)
	ctx.BodyParser(role)

	if message, err := ctrl.Valid(role); err {
		ctx.Status(401).JSON(message)
		return
	}

	if err := role.Create(); err != nil {
		ctx.Status(401).JSON(Map{
			"msg": err.Error(), 
		})
		return 
	}

	ctx.JSON(Map{
		"msg" : "Sucessfuly Create User" ,
		"data" : role,
	})
}

func (ctrl *RoleCtrl) Delete(ctx *fiber.Ctx) {

	id := ctx.Params("id")
	role, err := new(models.RoleModel).Get(id)
	
	if err != nil {
		ctx.Status(404).JSON( Map{ 
			"msg" : fmt.Sprintf("NotFound Role %s", id), 
		})
		return 
	}

	if err := role.Delete(); err != nil{
		ctx.Status(401).JSON(Map{ "msg" : err.Error(),  })
		return 
	}

	ctx.JSON(Map{
		"msg": fmt.Sprintf("Successfuly Delete Role %s", id),
	})
}

func (ctrl *RoleCtrl) Update(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	role, err := new(models.RoleModel).Get(id)
	
	if err != nil {
		ctx.Status(404).JSON( Map{ 

			"msg" : fmt.Sprintf("NotFound User %s", id), 
		})
		return 
	}

	ctx.BodyParser(&role)
	if err := role.Save(); err != nil {
		ctx.Status(401).JSON(Map{ "msg":err.Error(), })
		return 
	}

	ctx.JSON(Map{ 
		"msg" : "Successfuly Update User ",
		"data": role, 
	})
}