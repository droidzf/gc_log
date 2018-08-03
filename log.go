package log

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var isDEBUG = true

func SetDebug(debug bool) {
	isDEBUG = debug
}
func Info(arg ...interface{}) {
	if isDEBUG {
		s := buildString(" [INFO]  ", arg)
		fmt.Printf("%c[1;32m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func Debug(arg ...interface{}) {
	if isDEBUG {
		s := buildString(" [DEBUG] ", arg)
		fmt.Printf("%c[1;34m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func Warn(arg ...interface{}) {
	if isDEBUG {
		s := buildString(" [WARN]  ", arg)
		fmt.Printf("%c[1;33m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func Error(arg ...interface{}) {
	if isDEBUG {
		s := buildString(" [ERROR] ", arg)
		fmt.Printf("%c[1;31m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func buildString(level string, args []interface{}) string {
	var tag []interface{}

	tag = append(tag, time.Now().Format("2006-01-02 15:04:05.000000"), level, getPosition(), " -> ")
	s := fmt.Sprint(tag...) + fmt.Sprint(args...)
	return s
}

func getPosition() string {
	var file string
	var line int
	var ok bool
	_, file, line, ok = runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	path := strings.Split(file, "/")
	index := len(path) - 1
	return path[index] + ":" + strconv.Itoa(line)
}
