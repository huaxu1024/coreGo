package unsafe

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_unsafeAs(t *testing.T) {
	unsafeAs()
}

func Test_split(t *testing.T) {
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	fmt.Println(reflect.DeepEqual(got, want))

	var a, b []string = nil, []string{}
	fmt.Println(reflect.DeepEqual(a, b))

	var c, d map[string]int = nil, make(map[string]int)
	fmt.Println(reflect.DeepEqual(c, d))
}