package main

import (
	"strings"
	"testing"
)

func replaceSpaces(S string, length int) string {
	return strings.Replace(S[:length], " ", "%20",-1)
}
func TestReplaceSpaces(t *testing.T) {
	t.Log(replaceSpaces("Mr John Smith    ", 13))
}
