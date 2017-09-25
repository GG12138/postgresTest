package server

import (
	"github.com/go-xorm/xorm"
	"log"
	"fmt"
)

var engine *xorm.Engine

//连接数据库
func OpenPostgres(){
	var err error
	engine, err = xorm.NewEngine("postgres",
		"user=ssd password=zhukai.. dbname=mydb sslmode=disable")

	if err != nil {
		log.Println(err,"数据库连接错误")
		fmt.Println(engine.Ping())
	}
}
type Engine struct {
	Engine *xorm.Engine
}
func NewEngine() Engine{
	return Engine{
		Engine:engine,
	}
}


