package lib

import "log"

func Info(s string, v ...interface{}){
	if s=="" {
		log.Println(v)
	}else if v == nil{
		log.Printf(s)
	}else{
		log.Printf(s,v)
	}
}

