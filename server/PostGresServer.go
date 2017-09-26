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
func(e Engine) Sync(beans interface{}){
	err := e.Engine.Sync(beans)
	if  err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("----------------->同步数据库成功")
	}
}
//忽略主键添加（主键自增）
func(e Engine) OmitInsert(id string, bean interface{}) {
	_ , err :=e.Engine.Omit(id).Insert(bean)
	if err!= nil {
		log.Fatal("添加失败：%v", err)
	}else{
		fmt.Println("---------------->添加数据成功<--------------")
	}

}
//根据id查找
func(e Engine) GetId(id int64,bean interface{}) bool{
	has , err:= e.Engine.Id(1).Get(bean)
	if !has {
		log.Fatal("查询失败：%v",err)
	}
	return has
}
//查询所有
func(e Engine) GetAllSelect(bean interface{}){
	//bean 必须是指针
	err := e.Engine.Find(bean)
	if err!= nil {
		fmt.Printf("----: %v",err)
	}
}
//删除
func(e Engine) Delete (bean interface{}) int64 {
	ret, err := e.Engine.Delete(bean)
	if err !=nil {
		log.Fatal("删除错误 ：",err)
	}else{
		return ret
	}
	return -1
}

