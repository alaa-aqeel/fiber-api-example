package middleware 




import (
	"github.com/gofiber/fiber"
	"github.com/alaaProg/postapi/models"
)

func HashRolePermission(ctx *fiber.Ctx){

	if role, err :=  models.CurrentUser.Role(); err == nil{
		switch method := ctx.Method();{
			case method == "GET" && role.HashPermission("read"):
				ctx.Next()
				return

			case method == "POST" && role.HashPermission("write"):
				ctx.Next()
				return

			case method == "PUT" && role.HashPermission("update"):
				ctx.Next()
				return

			case method == "DELETE" && role.HashPermission("delete"):
				ctx.Next()
				return

		}
		
	}

	ctx.Status(401).JSON(fiber.Map{
		"msg": "Access Denied",
	})
}