package controller

import (
	"github.com/kataras/iris/v12/_examples/mvc/login/services"
	"time"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/consts"
	"github.com/liguoqinjim/iris_template/logger"
	"github.com/liguoqinjim/iris_template/service"
	"github.com/liguoqinjim/iris_template/validator"
	"github.com/liguoqinjim/iris_template/web/core"
	"github.com/liguoqinjim/iris_template/web/param"
	"github.com/liguoqinjim/iris_template/web/viewmodel"
)

type UserController struct{}

// @Summary 用户登录
// @Description 用户登录
// @Tags user
// @Accept  json
// @Produce  json
// @Param body body params.LoginParam true "login"
// @Success 200 {object} viewmodel.User
// @Failure 400 {object} viewmodel.Response
// @Router /user/login [post]
func (c *UserController) PostLogin(ctx iris.Context) error {
	param := &params.LoginParam{}
	if err := ctx.ReadJSON(param); err != nil {
		return consts.ErrorParam
	}

	reqId := core.GetReqID(ctx)
	logger.Infow("reqId", reqId)

	if err := validator.ValidateF(param); err != nil {
		return err
	}

	if user, err := service.UserService.Login(param); err != nil {
		return err
	} else {
		//jwt
		token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": time.Now().Add(time.Hour).Unix(),
			"iat": time.Now().Unix(),
			"u":   user.Id,
		})
		tokenString, _ := token.SignedString([]byte(config.Conf.Secret.Jwt))

		//更新jwt到redis或mysql

		ctx.JSON(viewmodel.NewUser(user, tokenString))
	}

	return nil
}

// PostRegister godoc
// @Summary 用户注册
// @Description 用户注册
// @Tags user
// @Accept  json
// @Produce  json
// @Param body body params.RegisterParam true "register"
// @Success 200 {object} datamodel.User
// @Failure 400 {object} viewmodel.Response
// @Router /user/register [post]
func (c *UserController) PostRegister(ctx iris.Context) error {
	param := new(params.RegisterParam)
	if err := ctx.ReadJSON(param); err != nil {
		return consts.ErrorParam
	}

	if err := validator.ValidateF(param); err != nil {
		return err
	}

	if exist, err := services.UserService.Exist(param.Username); err != nil {
		return consts.ErrorDB
	} else {
		if exist {
			return consts.ErrorUserAlreadyExist
		}
	}

	_, err := services.UserService.Register(param)
	if err != nil {
		return err
	}

	ctx.JSON(viewmodel.ResponseSuccess)
	return nil
}

// @Summary 用户信息
// @Description 用户信息
// @Tags user
// @Accept  json
// @Param	Authorization header string true "Bearer JwtToken"
// @Param   user_id     query    int     true        "用户id"
// @Success 200 {object} viewmodel.User
// @Router /user/info [get]
func (c *UserController) GetInfo(ctx iris.Context) error {
	userId, err := ctx.URLParamInt("user_id")
	if err != nil {
		logger.Errorf("ctx.URLParamInt error:%v", err)
		return err
	}

	ctx.Writef("get user info %d", userId)
	return nil
}

func (c *UserController) HandleError(ctx iris.Context, err error) {
	logger.Errorf("user controller handler error:%v", err)
	//ctx.Writef("user controller handler error")

	core.Response(ctx, nil, err)
}

//func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
//	log.Println("BeforeActivation")
//
//	//b.Handle("POST", "/login", "PostLogin")
//	//b.Handle(http.MethodPost, "/register", "PostRegister")
//	//
//	//b.Handle(http.MethodGet, "/info", "GetInfo")
//
//	//b.Handle(
//	//	"POST",
//	//	"/login",
//	//	"PostLogin",
//	//)
//
//	// or even add a global middleware based on this controller's router,
//	// which in this example is the root "/":
//	// b.Router().Use(myMiddleware)
//}
//func (c *UserController) AfterActivation(b mvc.AfterActivation) {
//
//	log.Println("AfterActivation")
//}
