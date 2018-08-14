package strev

import (
	"testing"
)

type ReverseTest struct {
	in, out string
}

var ReverseTests = []ReverseTest{
	{"ABCD", "DCBA"},
	{"ZXC-Y", "Y-CXZ"},
	{"Hello 世界", "界世 olleH"},
}

func TestReverse(t *testing.T) {
	for _, r := range ReverseTests {
		exp := Reverse(r.in)
		if r.out != exp {
			t.Errorf("Reverse of %s expects %s, but got %s", r.in, r.out, exp)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	s := "ABCD"
	for i := 0; i < b.N; i++ {
		Reverse(s)
	}
}
