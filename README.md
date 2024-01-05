# golang_web
搭建golangWeb服务


### simple 简单web服务

## [iris web框架](https://www.iris-go.com/docs/#/)

````
import "github.com/kataras/iris/v12"

func main() {
  app := iris.New()
  app.Use(iris.Compression)

  app.Get("/", func(ctx iris.Context) {
    ctx.HTML("Hello <strong>%s</strong>!", "World")
  })

  app.Listen(":8080")
}
````
