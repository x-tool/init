package tool

import "strconv"

func STI(s string) (i int, err error) {
	i, err = strconv.Atoi(s)
	return
}

func STI64(s string) (i int64, err error) {
	i, err = strconv.ParseInt(s, 10, 64)
	return
}

func STF64(s string) (f float64, err error) {
	f, err = strconv.ParseFloat(s, 64)
	return
}
