package assert

// ToAny changes string values to interface values.
func ToAny(args ...string) []interface{} {
	if args == nil || len(args) < 1 {
		return make([]interface{}, 0)
	}
	any := make([]interface{}, len(args))
	for i, arg := range args {
		any[i] = arg
	}
	return any
}
