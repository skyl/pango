package sliceutil

import (
  "testing"
)

type testCase struct {
  a []int
  b []int
}

func TestEqualInt(t *testing.T) {
  cases := []testCase{
    {
      []int{9, 12, 15, 18},
      []int{9, 12, 15, 18},
    },
  }
  for _, c := range cases {
    if !EqualInt(c.a, c.b) {
      t.Errorf("%q and %q should be equal", c.a, c.b)
    }
  }

  badcases := []testCase{
    {
      []int{1, 2, 3},
      []int{2, 3, 4},
    },
  }
  for _, c := range badcases {
    if EqualInt(c.a, c.b) {
      t.Errorf("%q and %q should not be equal", c.a, c.b)
    }
  }
}
