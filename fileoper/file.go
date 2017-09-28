package fileoper

import (
	"os"
	"fmt"
)


type FileOper struct {


}
func(f FileOper) Mkdir(filename, path string){
	if filename != "" && path == "" {
		err := os.Mkdir(filename,0777)
		if err != nil {
			fmt.Printf("mkdir error ... %v",err)
			return
		}
	}else if filename == "" && path != ""{
		err := os.MkdirAll(filename,0777)
		if err != nil {
			fmt.Printf("mkdirAll error ... %v", err)
			return
		}
	}else{
		fmt.Println("filename和path不能都为空")
	}
}

func(f FileOper) CreatFile(filename string){
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		fmt.Printf("create error ... :%v",err)
		return
	}

}
func(f FileOper) WriteFile(filename ,content string,b []byte){
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("write error ... :%v",err)
		return
	}
	defer file.Close()
	if content != "" && b == nil {
		//fmt.Printf("content :%v",file.Name())
		file.WriteString(content)
	}else if content == "" && b != nil {
		//fmt.Printf("byte : %v",file.Name())
		file.Write(b)
	}else{
		fmt.Println("content和[]byte都为空")
	}


}
func(f FileOper) ReadFile(filename string) []byte{
	ff, _ := os.Open(filename)
	defer  ff.Close()
	b := make([]byte, 1024)
	for{
		n, _ := ff.Read(b)
		if 0 == n {
			break
		}
	}
	return b
}
func(f FileOper)RemoveFile(filename ,path string){
	if filename != "" && path == ""{
		err := os.Remove(filename)
		if err != nil {
			fmt.Printf("remove error ... : %v",err)
		}
	}else if filename == "" && path != "" {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Printf("removeAll error ... : %v",err)
		}
	}else{
		fmt.Println("filename和path不能都为空")
	}

}
