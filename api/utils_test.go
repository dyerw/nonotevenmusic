package api

import "testing"

func TestDefMatch(t *testing.T) {
	s1 := DefMatch("")
	if s1 != ".*" {
		t.Error("Expected \".*\" got ", s1)
	}

	s2 := DefMatch("abc")
	if s2 != "abc" {
		t.Error("Expected \"abc\" got ", s2)
	}
}
