package goerrors

import "strconv"

func valueToString(val interface{}) string {
	switch v := val.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	case error:
		return v.Error()
	default:
		return "unknowntype"
	}
}
