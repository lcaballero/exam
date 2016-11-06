package assert

import (
	"fmt"
)

// HasArgs asserts that the variadic args is empty.
func HasArgs(msg ...interface{}) bool {
	return msg != nil && len(msg) > 0
}

// AsMsg joins the values without a format string.
func AsMsg(msg ...interface{}) string {
	return fmt.Sprint(msg...)
}

// AsFmt delegates the format interface to Sprintf(f, ...).
func AsFmt(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

// HasFmt checks that is not empty.
func HasFmt(f string) bool {
	return f != ""
}
