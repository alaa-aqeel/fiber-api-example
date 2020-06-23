package controller 


import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/alaaProg/postapi/models"
	"github.com/alaaProg/restful"
)

type PermissionCtrl struct{
	restful.Controller
}

func (ctrl *PermissionCtrl) Get(ctx *fiber.Ctx){

	ctx.JSON(new(models.PermissionModel).All())
} 

func  (ctrl *PermissionCtrl) Update(ctx *fiber.Ctx){

	prem, err := new(models.PermissionModel).Get(ctx.Params("id"))

	if err != nil {
		ctx.Status(404).JSON( Map{ 
			"msg" : fmt.Sprintf("NotFound Permission %s", ctx.Params("id")), 
		})
		return 
	}

	role, err := new(models.RoleModel).Get(ctx.FormValue("role"))
	if err != nil {
		ctx.Status(404).JSON( Map{ 
			"msg" : fmt.Sprintf("NotFound Role %s", ctx.FormValue("role")), 
		})
		return 
	}


	prem.AppendRole(role)

	ctx.JSON(Map{
		"msg": fmt.Sprintf("Successfuly Append Permission:%s to Role:%s", prem.Name, role.Name),
	})
}

func (ctrl *PermissionCtrl)  Delete(ctx *fiber.Ctx){

	prem, err := new(models.PermissionModel).Get(ctx.Params("id"))
	if err != nil {
		ctx.Status(404).JSON( Map{ 
			"msg" : fmt.Sprintf("NotFound Permission %s", ctx.Params("id")), 
		})
		return 
	}

	role, err := prem.GetRole(ctx.FormValue("role"))
	// ctx.JSON(role)
	// return 
	if err != nil {
		ctx.Status(404).JSON( Map{ 
			"msg" : fmt.Sprintf("Role:%s have not permission:%s ", 
						ctx.FormValue("role"), prem.Name), 
		})
		return 
	}


	prem.DeleteRole(role)

	ctx.JSON(Map{
		"msg": fmt.Sprintf("Successfuly Delete Permission:%s from Role:%s", prem.Name, role.Name),
	})
}