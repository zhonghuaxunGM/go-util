package main

import (
	"fmt"
	"runtime"
	"strings"
	"unicode"
)

//Assert checks if err is nil or not, if not, it panics with that err.
func Assert(err error) {
	if err != nil {
		panic(err)
	}
}

//Exception error with stack trace
type Exception []string

func (e Exception) Error() string {
	return strings.Join(e, "\n")
}

//Trace returns the current stack trace
func Trace(msg string, args ...interface{}) (logs Exception) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	msg = strings.TrimRightFunc(msg, unicode.IsSpace)
	if len(msg) > 0 {
		logs = Exception{msg}
	}
	n := 1
	for {
		n++
		pc, file, line, ok := runtime.Caller(n)
		if !ok {
			break
		}
		f := runtime.FuncForPC(pc)
		name := f.Name()
		if strings.HasPrefix(name, "runtime.") {
			continue
		}
		fn := file[strings.Index(file, "/src/")+5:]
		logs = append(logs, fmt.Sprintf("\t(%s:%d) %s", fn, line, name))
	}
	return
}
