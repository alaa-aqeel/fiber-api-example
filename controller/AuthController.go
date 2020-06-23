package controller 



import (
	"strconv"
	"github.com/gofiber/fiber"
	"github.com/alaaProg/restful"
	"github.com/alaaProg/postapi/models"
)

type AuthCtrl struct{
	restful.Controller

}

type FormValidation struct{

	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true" length:"6,20"`
}


func (ctrl *AuthCtrl) Login(ctx *fiber.Ctx){
	form := new(FormValidation)
	ctx.BodyParser(form)

	if message, err := ctrl.Valid(form); err {

		ctx.Status(401).JSON(message)
		return
	}

	user, err := new(models.UserModel).GetBy("username", form.Username)
	if err != nil || user.Password != form.Password {
		ctx.JSON(Map{
			"msg" : "Username or Password is incorrect",
		})
		return 
	}

	token, err := ctrl.CreateToken(strconv.Itoa(int(user.ID)), 10)
	if err != nil{
		panic(err)
		ctx.Status(401).Send()
		return 
	}

	ctx.JSON(Map{

		"token": token,
	})
}



func (ctrl *AuthCtrl) Register(ctx *fiber.Ctx){
	form := new(FormValidation)
	ctx.BodyParser(form)

	if message, err := ctrl.Valid(form); err {

		ctx.Status(401).JSON(message)
		return
	}

	
	user := new(models.UserModel)
	ctx.BodyParser(user)

	if err := user.Create(); err != nil {
		ctx.Status(401).JSON(Map{
			"msg": err.Error(), 
		})
		return 
	}

	ctx.JSON(Map{
		"data": user,
		"msg" : "Successfuly create new user",
	})
}