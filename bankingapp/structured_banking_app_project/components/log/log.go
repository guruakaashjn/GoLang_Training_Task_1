package log

import (
	"fmt"
	"time"
)

type Logger interface {
	Print(value ...string)
}
type Log struct{}

func GetLogger() *Log {
	return &Log{}
}

func (l *Log) Print(value ...interface{}) {
	fmt.Println(value)
}

func (l *Log) PrintError(er error) {
	fmt.Println("ERROR : ", er.Error(), " at : ", time.Now())
}
