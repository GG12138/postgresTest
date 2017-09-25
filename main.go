package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"./server"
	"log"
	"./users"
)

func main() {
	server.OpenPostgres()
	engine := server.NewEngine()
	err := engine.Engine.Sync(new(users.Student))
	if  err != nil{
		log.Fatal(err)
	}
	//sql, err:=engine.Engine.QueryString("select * from student")
	var s []users.Student

	err = engine.Engine.Find(&s,users.Student{})
	if err!= nil {
		fmt.Printf("%v",err)
	}

	for _,v:= range s {
		fmt.Println(v.Id)
	}
	//for _,v := range sql{
	//	for k, vv := range v {
	//		fmt.Print(k+ " : ")
	//		fmt.Println(vv)
	//	}
	//}
}




