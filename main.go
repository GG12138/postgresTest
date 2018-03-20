package main

import (
	"fmt"

	"./server"
	u "./users"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {

	//开启服务
	echo := echo.New()
	server.OpenPostgres()

	//同源跨域
	echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://192.168.0.101:3000"},
		AllowCredentials: true,
	}))

	echo.POST("/", Login)
	echo.Logger.Fatal(echo.Start(":8080"))
}
func Login(e echo.Context) error {

	username := e.FormValue("username")
	password := e.FormValue("password")
	fmt.Println(e.FormParams())
	engine := server.NewEngine()
	user := &u.Users{
		Name:     username,
		Password: password,
	}
	has, _ := engine.Engine.Get(user)
	if !has {
		fmt.Println("登录失败")
		return e.String(200, "err")
	}
	return e.JSON(200, user)
}
