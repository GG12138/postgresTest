package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"./server"
	"./users"
)

func main() {
	//开启服务
	server.OpenPostgres()
	engine := server.NewEngine()
	user := new(users.Users)
	has, _ := engine.Engine.IsTableExist(user)
	_, _ = engine.Engine.IsTableEmpty(user)
	if !has {
		fmt.Println("------>同步数据库字段<--------")
		engine.Sync(user)
	}
	//删除数据库
	//engine.Engine.DropTables(user)

	uu := user.AddUsers("男", "朱凯", 18)
	engine.OmitInsert("id",uu)
	//user.Name = "王绿绿"
	//user.Sex = "女"
	//user.Age = 17
	//session := engine.Engine.NewSession()
	//defer session.Close()
	//err := session.Begin()
	//if err !=nil{
	//
	//}
	//_, err = session.Id(1).Update(user)
	//if err != nil {
	//	fmt.Println("--------------->修改错误 session.Rollback")
	//	session.Rollback()
	//	return
	//}
	//err = session.Commit()
	//if err != nil{
	//	return
	//}
	//engine.Engine.Id(1).Update(user)
	u := make([]users.Users,0)
	engine.GetAllSelect(&u)
	for _, v := range u {
		fmt.Println(v)
	}



}




