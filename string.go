package tool

import "bytes"

func JoinStringWithComma(lst []string) string {
	var buf bytes.Buffer
	for i, v := range lst {
		if i == 0 {
			buf.WriteString(v)
		} else {
			buf.WriteString(",")
			buf.WriteString(v)
		}
	}
	return buf.String()
}

func JoinStringWithCommaEnter(lst []string) string {
	var buf bytes.Buffer
	for i, v := range lst {
		if i == 0 {
			buf.WriteString(v)
		} else {
			buf.WriteString(",\n")
			buf.WriteString(v)
		}
	}
	return buf.String()
}
