package controller

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/consts"
	"github.com/liguoqinjim/iris_template/logger"
	"github.com/liguoqinjim/iris_template/service"
	"github.com/liguoqinjim/iris_template/validator"
	"github.com/liguoqinjim/iris_template/web/core"
	"github.com/liguoqinjim/iris_template/web/param"
	"time"
)

type UserController struct{}

// @Summary 用户登录
// @Description 用户登录
// @Tags user
// @Accept  json
// @Produce  json
// @Param body body param.LoginParam true "login"
// @Success 200 {object} datamodel.User
// @Failure 400 {object} viewmodel.Response
// @Router /user/login [post]
func (c *UserController) PostLogin(ctx iris.Context) error {
	p := &param.LoginParam{}
	if err := ctx.ReadJSON(p); err != nil {
		return consts.ErrParam
	}

	if err := validator.ValidateStruct(p); err != nil {
		return err
	}

	if user, err := service.UserService.Login(p); err != nil {
		return err
	} else {
		//jwt
		token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": time.Now().Add(time.Hour).Unix(),
			"iat": time.Now().Unix(),
			"u":   user.Id,
		})
		if t, err := token.SignedString([]byte(config.Config.Secret.Jwt)); err != nil {
			logger.Errorf("jwt token", "err", err)
			return err
		} else {
			user.Token = t
		}

		//更新jwt到redis或mysql

		core.Response(ctx, user, nil)
	}

	return nil
}

// @Summary 用户注册
// @Description 用户注册
// @Tags user
// @Accept  json
// @Produce  json
// @Param body body param.RegisterParam true "register"
// @Success 200 {object} datamodel.User
// @Failure 400 {object} viewmodel.Response
// @Router /user/register [post]
func (c *UserController) PostRegister(ctx iris.Context) error {
	p := new(param.RegisterParam)
	if err := ctx.ReadJSON(p); err != nil {
		return consts.ErrParam
	}

	if err := validator.ValidateStruct(p); err != nil {
		return err
	}

	if exist, err := service.UserService.Exist(p.Username); err != nil {
		return consts.ErrDB
	} else {
		if exist {
			return consts.ErrUserNotFound
		}
	}

	if _, err := service.UserService.Register(p); err != nil {
		return err
	}

	core.Response(ctx, nil, nil)
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

//每个controller各自的error handler可以覆盖总的error handler
func (c *UserController) HandleError(ctx iris.Context, err error) {
	logger.Errorf("user controller handler error:%v", err)

	core.Response(ctx, nil, err)
}
