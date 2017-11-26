package log

import "log"

func Println(v ...interface{}) {
	log.Print("Log:")
	log.Println(v)
}
