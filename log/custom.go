package log

import (
	"fmt"
	"time"
)

func Println(v ...interface{}) {
	fmt.Print(time.Now())
	fmt.Println(v)
}
