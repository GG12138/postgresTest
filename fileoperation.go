package main

import (
	"./fileoper"
	"fmt"
)

func main(){
	file := new(fileoper.FileOper)
	file.WriteFile("fileoper/zz.txt","aa",nil)
	b := file.ReadFile("fileoper/zz.txt")
	fmt.Println(string(b))
	file.RemoveFile("","fileoper/file")
	file.RemoveFile("fileoper/zz.txt", "")
}

