package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"./server"
	"log"
	"./users"
)

func main() {
	//开启服务
	server.OpenPostgres()
	engine := server.NewEngine()

	//同步数据库字段
	err := engine.Engine.Sync(new(users.Users))
	if  err != nil{
		log.Fatal(err)
	}

	user := make([]users.Users,0)
	u := users.Users{}.AddUsers("女","王绿绿", 1, 17)
	engine.Engine.ShowSQL(true)
	_, err =engine.Engine.Omit("id").Insert(u)

	if err != nil{
		fmt.Println(err)
	}

	err = engine.Engine.Find(&user,users.Users{})
	if err!= nil {
		fmt.Printf("%v",err)
	}
	fmt.Printf("%+v",user)
	//for _,v:= range s {
	//	fmt.Println(v.Id)
	//}
}




