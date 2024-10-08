// Mengubah nilai rupiah menjadi dollar
package main

import "fmt"

func main() {
	var rupiah, dollar float64
	fmt.Print("Masukan Nominal Rupiah: ")
	fmt.Scan(&rupiah)
	dollar = (rupiah / 15000)
	fmt.Print("Jadi ", rupiah, " rupiah = ", dollar, " dollar")
}
