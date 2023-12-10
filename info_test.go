package utils

import "testing"

func add(left int, right int) int {
	return left + right
}

func TestNameOfPackage(t *testing.T) {
	const nameOfPackage = "github.com/shplume/utils"
	name := NameOfPackage()

	if name != nameOfPackage {
		t.Errorf("NameOfPackage() is %s", name)
	}
}

func TestQueryName(t *testing.T) {
	sub := func(left int, right int) int {
		return left - right
	}

	cases := []struct {
		object any
		expect string
	}{
		{add, "github.com/shplume/utils.add"},
		{sub, "github.com/shplume/utils.TestQueryName.func1"},
		{1, "1"},
	}

	for _, c := range cases {
		name := QueryName(c.object)
		if name != c.expect {
			t.Errorf("QueryName(%v) is %s", c.object, name)
		}
	}
}
