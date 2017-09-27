package main

import (
	"encoding/xml"
	"os"
	"fmt"
	"io/ioutil"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string `xml:",innerxml"`
}
type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main(){
	file, err := os.Open("xml/server.xml")
	if err != nil {
		fmt.Printf("file error : %v",err)
		return
	}
	defer file.Close()

	data,err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("data error : %v",err)
		return
	}

	v := Recurlyservers{}

	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("xml error: %v", err)
		return
	}
	fmt.Println(v.XMLName)
	for _, s :=range v.Svs{
		fmt.Println(s.XMLName)
		fmt.Println(s.ServerIP)
		fmt.Println(s.ServerName)
	}
	fmt.Println(v.Version)
	fmt.Println(v.Description)
}
