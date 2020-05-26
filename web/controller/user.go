package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/liguoqinjim/iris_template/consts"
	"github.com/liguoqinjim/iris_template/logger"
	"github.com/liguoqinjim/iris_template/service"
	"github.com/liguoqinjim/iris_template/validator"
	"github.com/liguoqinjim/iris_template/web/core"
	"github.com/liguoqinjim/iris_template/web/param"
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
// @Router /user/register [post]
func (c *UserController) PostRegister(ctx iris.Context) error {
	p := new(param.RegisterParam)
	if err := ctx.ReadJSON(p); err != nil {
		return consts.ErrParam
	}

	if err := validator.ValidateStruct(p); err != nil {
		return err
	}

	if data, err := service.UserService.Register(p); err != nil {
		return err
	} else {
		core.Response(ctx, data, nil)
	}

	return nil
}

//每个controller各自的error handler可以覆盖总的error handler
func (c *UserController) HandleError(ctx iris.Context, err error) {
	logger.Errorf("user controller handler error:%v", err)

	core.Response(ctx, nil, err)
}
