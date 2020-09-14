package kata

import "fmt"

func breakChocolate(n, m int) int {
  if n < 1 || m < 1 {
    return 0
  }
  return n*m-1
  
}

func main () {
	fmt.Println(breakChocolate(2, 3))
}