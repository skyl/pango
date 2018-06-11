package mathutil

import "fmt"

// PrintMultiplesOf prints the multiples of a given integer between
// low and high
func PrintMultiplesOf(of int, low int, high int) {
  for i := low; i <= high; i++ {
    if i%of == 0 {
      fmt.Println(i)
    }
  }
}
