package main

import (
	"github.com/kataras/iris"
)

func main() {
	//Initialize App
	app := iris.New()

	//First method, use serve file to render content
	app.Get("/", func(ctx iris.Context) {
		ctx.ServeFile("templatesAndHtml/views/index.html", false)
	})
	//Using a seperate dir for content
	app.Get("/home", func(ctx iris.Context) {
		ctx.ServeFile("templatesAndHtml/views/home.html", false)
	})
	//Register the engine(view enginer)
	app.RegisterView(iris.HTML("templatesAndHtml/views", ".html"))
	//Can utilize handlebars as well
	//app.RegisterView(iris.Handlebars("/templatesAndHtml/views", ".html"))
	app.Get("/about", func(ctx iris.Context) {
		ctx.ViewData("firstname", "Jessie JCharis")
		ctx.View("about.html")
	})

	app.Run(iris.Addr(":8080"))
}
