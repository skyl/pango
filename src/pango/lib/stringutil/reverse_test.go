package stringutil

import (
  "testing"
)

func TestReverse(t *testing.T) {
  cases := []struct {
    in, want string
  }{
    {"abc", "cba"},
    {"foobar", "raboof"},
  }
  for _, c := range cases {
    got := Reverse(c.in)
    if got != c.want {
      t.Errorf("Reverse(%q) expected %q actual %q", c.in, c.want, got)
    }
  }
}
