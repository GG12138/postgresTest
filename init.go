package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"xorm"
)
var engine *xorm.Engine

func main() {
	engine,_ := xorm.NewEngine("postgres", "user=ssd password=zhukai.. dbname=mydb sslmode=disable")

	sql := "select * from mytable"
	rowArray, _ := engine.Query(sql)
	user := new(User)
	for _,v := range rowArray{
		 for k, vv := range v {
			 if k =="name" {
				 user.name = string(vv)
				 //user.id, _ = strconv.Atoi(id)
			 }
		 }
	}
	fmt.Println(user.name)
}

type User struct {
	id int
	name string
	birthday string
	age int
}
