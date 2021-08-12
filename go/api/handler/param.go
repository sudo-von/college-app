package handler

import (
	"net/url"
	"strconv"
)

func ParamToInt(param string, url url.Values) int {
	var value int
	if s := url.Get(param); s != "" {
		value, err := strconv.Atoi(s)
		if err != nil {
			return value
		}
		return value
	}
	return value
}
