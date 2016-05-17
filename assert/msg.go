package assert

import (
	"fmt"
)

func HasArgs(msg ...interface{}) bool {
	return msg != nil && len(msg) > 0
}

func AsMsg(msg ...interface{}) string {
	return fmt.Sprint(msg...)
}

func AsFmt(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func HasFmt(f string) bool {
	return f != ""
}
