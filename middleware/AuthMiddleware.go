package middleware 

import (
	"strings"
	"github.com/alaaProg/restful"
	"github.com/gofiber/fiber"
	"github.com/alaaProg/postapi/models"
)



func Auth(ctx *fiber.Ctx){
	token := ctx.Get("Authorization")

	if token == ""{
		ctx.Status(401).JSON(fiber.Map{
			"msg": "Unauthorized",
		})
		return 
	}

	token = strings.Split(token, " ")[1]

	tkn, user, err := restful.VerifyToken(token)
	if err != nil {
		ctx.Status(401).JSON(fiber.Map{
			"msg": err.Error(),
		})
		return 
	}

	

	if tkn.Valid{

		models.CurrentUser, err = new(models.UserModel).Get(user.Id)
		if err != nil{
			ctx.Status(401).JSON(fiber.Map{
				"msg": "Unauthorized",
			})
			return 
		}

		ctx.Next()
		return 
	}

	ctx.Status(401).JSON(fiber.Map{
		"msg": "Unauthorized",
	})
}