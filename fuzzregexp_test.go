package main_test

import (
	"regexp"
	"regexp/syntax"
	"testing"
)

func FuzzRegexp(f *testing.F) {
	f.Add(".")
	f.Fuzz(func(t *testing.T, pat string) {
		_, err := regexp.Compile(pat)
		if err != nil {
			t.Skip()
		}
		_, err = syntax.Parse(pat, 0)
		if err != nil {
			t.Errorf("%q: %v", pat, err)
		}
	})
}
