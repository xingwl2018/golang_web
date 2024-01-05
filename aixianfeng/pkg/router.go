package pkg

import (
	"aixianfeng/src/account"
	"aixianfeng/src/activity"
	"aixianfeng/src/brand"
	"aixianfeng/src/exchange_coupons"
	"aixianfeng/src/product"
	"aixianfeng/src/province"
	"aixianfeng/src/rule"
	"aixianfeng/src/shop"
	"aixianfeng/src/tags"
	"aixianfeng/src/unit"
	"aixianfeng/src/vip_member"
	"github.com/kataras/iris/v12"
	"net/http"
	"time"
)

var (
	VERSION = "v0.1.0"
)

// ApplyRouter 服务总路由
func ApplyRouter() *iris.Application {
	//Default returns a new Application.
	app := iris.Default()

	notFound(app)

	// 添加路由
	app.Handle("GET", "/", func(context iris.Context) {
		_ = context.JSON(iris.Map{
			"data": time.Now().Format("2006-01-02 15:04:05"),
			"code": http.StatusOK,
		})
	})
	//
	app.Get("/heart", func(c iris.Context) {
		c.JSON(iris.Map{
			"data": time.Now().Format("2006-01-02 15:04:05"),
			"code": http.StatusOK,
		})
	})
	//定义一组路由
	v1 := app.Party("/v1")
	v1.Get("/version", func(context iris.Context) {
		context.JSON(
			iris.Map{
				"code":    http.StatusOK,
				"version": VERSION,
			},
		)
		return
	})

	app.UseGlobal(LoggerForProject)

	{
		account.Default.RegisterWithOut(app, "/v1")
		rule.Default.RegisterWithout(app, "/v1")
		province.Default.RegisterWithOut(app, "/v1")
		shop.Default.RegisterWithout(app, "/v1")
		activity.Default.Register(app, "/v1", false)
		unit.Default.Register(app, "/v1")
		brand.Default.Register(app, "/v1")
		tags.Default.Register(app, "/v1")
		product.Default.Register(app, "/v1")
	}

	app.Use(TokenForProject)

	{
		account.Default.RegisterWith(app, "/v1")
		vip_member.Default.Register(app, "/v1")
		exchange_coupons.Default.Register(app, "/v1")
		activity.Default.Register(app, "/v1", true)
		order.Default.Register(app, "/v1")
	}

	app.Logger().SetLevel("debug")
	return app
}

func notFound(app *iris.Application) {
	//返回找不到URL
	app.OnErrorCode(http.StatusNotFound, func(context iris.Context) {
		context.JSON(iris.Map{
			"code":   http.StatusNotFound,
			"detail": context.Request().URL.Path,
			"error":  "error found",
		})
	})
	return
}
