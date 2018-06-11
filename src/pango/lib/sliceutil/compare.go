package sliceutil

// EqualInt is a deprecated function that takes two slices of ints and returns
// a boolean of whether the two slices have equal values at every position
// instead, use reflect.DeepEqual https://stackoverflow.com/a/15312182
func EqualInt(a, b []int) bool {
    if a == nil && b == nil {
        return true;
    }
    if a == nil || b == nil {
        return false;
    }
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
