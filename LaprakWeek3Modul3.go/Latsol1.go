// program digunakan untuk menghitung nilai 𝑥 dalam persamaan
package main

import "fmt"

func main() {
	var x, fungsi float64
	fmt.Scan(&x)
	fungsi = (2 / (x + 5)) + 5
	fmt.Print(fungsi)
}
