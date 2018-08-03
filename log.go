package log

import (
	"bytes"
	"time"
	"strconv"
	"os"
	"runtime"
	"fmt"
)

var isDEBUG = true;

func SetDebug(debug bool) {
	isDEBUG = debug
}
func Info(arg ...interface{}) {
	if isDEBUG {
		s := buildString("[INFO] ", arg)
		fmt.Printf("%c[1;32m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func Debug(arg ...interface{}) {
	if isDEBUG {
		s := buildString("[DEBUG]", arg)
		fmt.Printf("%c[1;34m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func Warn(arg ...interface{}) {
	if isDEBUG {
		s := buildString("[WARN] ", arg)
		fmt.Printf("%c[1;33m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func Error(arg ...interface{}) {
	if isDEBUG {
		s := buildString("[ERROR]", arg)
		fmt.Printf("%c[1;31m%s%c[0m\n", 0x1B, s, 0x1B)
	}
}

func buildString(level string, args []interface{}) string {
	var tag []string
	var file string
	var line int
	var ok bool
	_, file, line, ok = runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	tag = append(tag, time.Now().Format("2006-01-02 15:04:05"), level, strconv.Itoa(os.Getpid()), file+":"+strconv.Itoa(line), ":")
	var buffer bytes.Buffer
	for _, value := range args {
		tag = append(tag, fmt.Sprint(value))
	}
	for _, value := range tag {
		buffer.WriteString(value)
		buffer.WriteString(" ")
	}
	s := buffer.String()
	return s
}
