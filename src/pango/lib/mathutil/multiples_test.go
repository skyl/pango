package mathutil

import (
  // "fmt"
  "testing"
  "pango/lib/sliceutil"
)

func TestMultiplesOf(t *testing.T) {
  // fmt.Println("TESTING...")
  cases := []struct {
    multiples, low, high int
    expected []int
  }{
    {3, 7, 20, []int{9, 12, 15, 18}},
  }
  for _, c := range cases {
    got := MultiplesOf(c.multiples, c.low, c.high)
    // fmt.Println("GOT", got)
    // fmt.Println("EXPECTED", c.expected)
    if !sliceutil.EqualInt(got, c.expected) {
      t.Errorf("MultiplesOf(%q, %q, %q) expected %q actual %q", c.multiples, c.low, c.high, c.expected, got)
    }
  }
}
