package main

import "github.com/kataras/iris"

type User struct {
	Name 	string `json:"name" xml:"name" form:"name"`
	Email 	string `json:"email" xml:"email" form:"email"`
}

func getSchool(ctx *iris.Context) {
    ctx.HTML(iris.StatusOK, "<h1> Hello World!</h1>")
}

func getUser(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, User{Name: "zhangchengfei", Email: "chengfeizh@gmail.com"})
	//ctx.Render(iris.StatusOK, User{Name: "zhangchengfei", Email: "chengfeizh@gmail.com"}, iris.RenderOptions{"charset": "UTF-8"})
}

func main() {
    server := iris.New()

	// routers
	server.Get("/", getSchool)
	server.Get("/user", getUser)
	//server.Post()
	//server.Put()
	//server.Delete()

	// server configuration
	server.Listen(":8888")
}
