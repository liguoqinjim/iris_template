package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/liguoqinjim/iris_template/config"
	"github.com/liguoqinjim/iris_template/web/core"
)

func Jwt() context.Handler {
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (i interface{}, err error) {
			return []byte(config.Config.Secret.Jwt), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		Expiration:    true,
		ErrorHandler: func(ctx context.Context, err error) {
			ctx.StopExecution()
			ctx.StatusCode(iris.StatusUnauthorized)

			core.Response(ctx, nil, err)
		},
	}).Serve
}
