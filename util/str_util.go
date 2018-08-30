package util

import "strconv"

func String2Int(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func String2Int64(str string) int64  {
	val, _ := strconv.ParseInt(str, 10, 64)
	return val
}

func Int2String(in int) string {
	return strconv.Itoa(in)
}

func Int642String(in int64) string {
	return strconv.FormatInt(in,10)
}