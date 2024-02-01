package main

import (
	"errors"
	"fmt"
	"log"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	fmt.Printf("runes: %q\n", r)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func main() {
	var err error
	input := "The quick brown fox jumped over the lazy dog"
	rev, err := Reverse(input)
	if err != nil {
		log.Fatalf(err.Error())
	}
	doubleRev, err := Reverse(rev)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}
