package main

import (
	"github.com/kataras/iris"
)

func main() {
	//Initialize App
	app := iris.New()

	//Register views
	app.RegisterView(iris.HTML("formHandling/views", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.View("form.html")
	})

	//Setting up form submit function
	app.Post("/add", func(ctx iris.Context) {
		username := ctx.PostValue("username")
		message := ctx.PostValue("message")
		ctx.ViewData("Username", username)
		ctx.ViewData("Message", message)
		ctx.View("form.html")
	})
	app.Run(iris.Addr(":8080"))
}
