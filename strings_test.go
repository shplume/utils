package utils

import "testing"

func TestIsASCII(t *testing.T) {
	cases := []struct {
		str    string
		expect bool
	}{
		{"ae12&6^5", true},
		{"我是一个中文字符串", false},
	}

	for _, c := range cases {
		isASCII := IsASCII(c.str)
		if isASCII != c.expect {
			t.Errorf("IsASCII(%s) is %v", c.str, isASCII)
		}
	}
}
