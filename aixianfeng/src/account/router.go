package account

import "github.com/kataras/iris/v12"

// 度量单位的路由

type ControllerForAccount struct {
}

var Default = ControllerForAccount{}

func (controller ControllerForAccount) RegisterWithOut(app *iris.Application, path string) {
	middleware := func(context iris.Context) {
		context.Next()
	}

	account := app.Party(path, middleware)
	{
		account.Post("/register", registerHandle)
		account.Post("/sign", signHandle)

	}

}

func (controller ControllerForAccount) RegisterWith(app *iris.Application, path string) {
	middleware := func(context iris.Context) {
		context.Next()
	}

	account := app.Party(path, middleware)
	{
		account.Post("/logout", logoutHandle)
		account.Get("/account/{id:uint}", getAccountHandle)
	}
}
