// program yang mencari volume dan luas pada bola
package main

import "fmt"

func main() {
	var r float64
	var volume, luas float64
	fmt.Scan(&r)
	fmt.Printf("jejari :", r)
	volume = (4.0 / 3.0) * 3.1415926535 * r * r * r
	luas = 4 * 3.1415926535 * r * r
	fmt.Printf("Bola dengan jejari %.f memiliki volume %.4f dan luas kulit %.4f", r, volume, luas)
}
