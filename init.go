package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/go-xorm/xorm"
	"log"
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
type Student struct {
	Id int64
	Name string `xorm:"unique"`
	Email string
}
type Users struct {
	Id int64 `xorm:"pk"`
	Name string `xorm:"unique"`
	Age int64
	Sex string
}

func main() {
	OpenPostgres()
	err := engine.Sync(new(Users))
	if  err != nil{
		log.Fatal(err)
	}
	//users := &Users{
	//	Name:"朱凯",
	//	Age:18,
	//	Sex:"男",
	//}
	//ret , err := engine.Insert(users)
	//fmt.Println(ret)
	//if err!= nil {
	//	fmt.Printf("%v",err)
	//}
	//students := &Student{
	//	Id:1,
	//	Name:"朱凯",
	//	Email:"1842159520@qq.com",
	//}
	//ref , err :=engine.Insert(students)
	//if err!= nil {
	//	fmt.Printf("%v",err)
	//}
	//var s []Student
	//err = engine.Find(&s,&Student{Id:1})
	//if err!= nil {
	//	fmt.Printf("%v",err)
	//}
	//for _,v := range s{
	//	fmt.Println(v.Name)
	//	fmt.Println(v.Email)
	//}
}




