package gtime

import "time"

func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Timestamp() int64 {
	return time.Now().Unix()
}
