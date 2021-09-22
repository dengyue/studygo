package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	ret := Split("abcdabcd", "a")
	want := []string{"", "bcd", "bcd"}
	if !reflect.DeepEqual(ret, want) {
		t.Fatalf("want:%v but got: %v\n", want, ret)
	}
}

func Test2Split(t *testing.T) {
	ret := Split("abc,,,,abc", ",")
	want := []string{"abc", "", "", "", "abc"}
	if !reflect.DeepEqual(ret, want) {
		t.Fatalf("want:%v but got: %v\n", want, ret)
	}
}

func Test3Split(t *testing.T) {
	type testCase struct {
		input string
		sep   string
		want  []string
	}
	testGroups := []testCase{
		{"abcdabcd", "a", []string{"", "bcd", "bcd"}},
		{"abc,,,,abc", ",", []string{"abc", "", "", "", "abc"}},
	}
	for _, tg := range testGroups {
		ret := Split(tg.input, tg.sep)
		if !reflect.DeepEqual(ret, tg.want) {
			t.Fatalf("want:%v but got: %v\n", tg.want, ret)
		}
	}
}
