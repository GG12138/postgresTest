package main

import (
	"github.com/devfeel/dotweb"
	"github.com/astaxie/beego"
	"github.com/echo/labstack/echo"
	s "../server"
	u "../users"
	"fmt"
)

type Echo struct {
	echo.Context
}

func main(){
	e := echo.New()

	e.POST("/",Login)
	e.Logger.Fatal(e.Start(":8080"))
	//StartServer()
}

func Login(e echo.Context)error{
	username := e.FormValue("username")
	password := e.FormValue("password")
	fmt.Println(username)
	engine := s.NewEngine()
	user :=&u.Users{
		Name     : username,
		Password : password,
	}
	has ,_ := engine.Engine.Get(user)
	if !has {
		fmt.Println("登录失败")
		return e.String(200,"err.....")
	}


	return e.String(200,"echo start....")
}




type MainController struct {
	beego.Controller
}

func (m MainController) Get(){
	m.Ctx.WriteString("sb")
}

func StartServer(){
	//beego.Router("/",&MainController{})
	//beego.Run()
	app := dotweb.New()
	app.HttpServer.GET("/:zhukai",func(c dotweb.Context) error{
		z := c.GetRouterName("zhukai")
		c.WriteString(z)
		return nil
	})
	app.StartServer(8080)

}