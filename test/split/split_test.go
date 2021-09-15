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
