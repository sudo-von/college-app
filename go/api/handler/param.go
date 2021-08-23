package handler

import (
	"net/url"
	"strconv"
	"time"
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

func ParamToDate(param string, url url.Values) *time.Time {
	if s := url.Get(param); s != "" {
		value, err := time.ParseInLocation("2006-01-02", s, time.Local)
		if err != nil {
			return nil
		}
		return &value
	}
	return nil
}
