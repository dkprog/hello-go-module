package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Daniel", "leinaD"},
		{"Hello, World", "dlroW ,olleH"},
		{"", ""},
	}

	for _, c := range cases {
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
