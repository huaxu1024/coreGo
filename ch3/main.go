package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	_ = 1 << (10 * iota)
	kib
	mib
	gib
	tib
	pib
	eib
	zib
	yib
)

func main() {
	fmt.Println(kib)
	fmt.Println(mib)
	fmt.Println(gib)
	fmt.Println(tib)
	fmt.Println(pib)

	formatIS()

	fmt.Println(intsToString([]int{1,2,3,4,5}))

	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))
	fmt.Println(basename("ab/sc.") == basename2("ab/sc."))
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i -- {
		if  s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i -- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}


func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func formatIS() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(y == strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 2))
	s := fmt.Sprintf("x=%b", x)
	fmt.Println(s)
}

