package mathutil

// MultiplesOf gathers the multiples of a given integer between
// low and high and returns the answer as a slice
func MultiplesOf(of int, low int, high int) []int {
  ret := make([]int, 0)
  for i := low; i <= high; i++ {
    if ((i % of) == 0) {
      ret = append(ret, i)
    }
  }
  return ret
}
