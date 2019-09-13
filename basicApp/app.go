package main

import (
	"github.com/kataras/iris"
)

func main() {
	//Initialize App
	app := iris.New()

	//set base get method
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h2>Hello world this is Iris</h2>")
	})
	//Text
	app.Get("/text", func(ctx iris.Context) {
		ctx.Text("Hello world this is Text")
	})
	//String
	app.Get("/string", func(ctx iris.Context) {
		ctx.WriteString("Hello world this is string format")
	})
	//JSON
	app.Get("/api", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"Msg": "Hello this is JSON"})
	})
	//Handler
	app.Handle("GET", "/simple", func(ctx iris.Context) {
		ctx.HTML("<h2>Hello world this is Iris</h2>")
	})

	app.Run(iris.Addr(":5001"))
}
