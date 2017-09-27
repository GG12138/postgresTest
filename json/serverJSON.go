package main

import (
	"fmt"
	"encoding/json"

)

type Server struct {
	ServerName string
	ServerIP string
}
type Serverslice struct {
	Servers []Server
}
func main() {
	//var s Serverslice
	//str := `{"servers":
	//		[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
	//		{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	//json.Unmarshal([]byte(str),&s)
	//fmt.Printf("%+v", s)
	var f interface{}
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	err := json.Unmarshal(b,&f)
	if err!= nil {
		fmt.Println(err)
	}
	m := f.(map[string]interface{})
	//js ,_ := simplejson.NewJson(b)
	//
	//fmt.Println(js.Get("Name").String())


	//fmt.Println(arr)

	for k,v := range m  {
		switch vv:= v.(type){
		case string:
			fmt.Println(k,"is string",vv)
		case int:
			fmt.Println(k,"is int",vv)
		case []interface{}:
			fmt.Println(k,"is []interface{}",vv)
			for kk, vvv := range vv {
				fmt.Println(kk, vvv)
			}
		default:
			fmt.Println(k,vv)
			//fmt.Println(k, "is of a type I don't know how to handle")
		}

	}

}

