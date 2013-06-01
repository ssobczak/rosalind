package gorosalind

import (
	"testing"
)

func arrayEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, ak := range a {
		if ak != b[k] {
			return false
		}
	}

	return true
}

func testMatch(t *testing.T, in string, out []int) {
	if ret := match(in); !arrayEq(ret, out) {
		t.Error("match failed: ", in, ret, out)
	}
}

func TestMatch_Find(t *testing.T) {
	testMatch(t, "NSSS", []int{1})
	testMatch(t, "ABCNSTX", []int{4})
	testMatch(t, "QQNSTXYZ", []int{3})
}

func TestMatch_NotFound(t *testing.T) {
	testMatch(t, "", []int{})
	testMatch(t, "N", []int{})
	testMatch(t, "c", []int{})
	testMatch(t, "zsdaq", []int{})
	testMatch(t, "wrNSTPld", []int{})
}

func TestMatch_Multiple(t *testing.T) {
	testMatch(t, "NQSQxxNSSH", []int{1, 7})
	testMatch(t, "qwNHTXNHTXqq", []int{3, 7})
}

func TestMatch_Overlapping(t *testing.T) {
	testMatch(t, "NQSNQSNQSX", []int{1, 4, 7})
	testMatch(t, "qwNXTNXSNXz", []int{3, 6})
}
