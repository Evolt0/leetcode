package main

import (
	"fmt"
	"strings"
	"testing"
)

func oneEditAway(first string, second string) bool {
	if len(first)-len(second) >= 2 || len(first)-len(second) <= -2 {
		return false
	}
	return false
}

func TestOneEditAway(t *testing.T) {
	fmt.Println(strings.Contains("plae",""))
}
