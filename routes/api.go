package routes 


import (
	"github.com/gofiber/fiber"
	"github.com/alaaProg/restful"
	ctrl "github.com/alaaProg/postapi/controller"
	middl "github.com/alaaProg/postapi/middleware"
)



func Api(app *fiber.App){


	api   := app.Group("/api/")
	admin := api.Group("/admin/")

	api.Post("/login/", new(ctrl.AuthCtrl).Login)
	api.Post("/register/", new(ctrl.AuthCtrl).Register)

	api.Use(middl.Auth)

	// 
	api.Get("/user/", new(ctrl.AccountCtrl).Get)
	api.Put("/user/", new(ctrl.AccountCtrl).Update)
	api.Delete("/user/", new(ctrl.AccountCtrl).Delete)


	admin.Use(middl.HashRolePermission)
	// 
	restful.Resource("/user/", admin, new(ctrl.UserCtrl))
	restful.Resource("/permission/", admin, new(ctrl.PermissionCtrl))

	//
	role := restful.Resource("/role/", admin, new(ctrl.RoleCtrl))
	role.Post("/:roleid/user", new(ctrl.RoleCtrl).AppendRoleToUser)
	role.Delete("/:roleid/user", new(ctrl.RoleCtrl).DeleteRoleFromUser)


}