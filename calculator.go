package main

import (
	"fmt"
)

func main() {
	var a, z string
	var b, c, d, e float32
	a = "y"
	e = 1

	fmt.Println("selamat menggunakan operator aritmatika")

	for a != "n" {
		fmt.Print("apakah anda akan melakukan perhitungan ( y / n) = ")
		fmt.Scan(&a)

		if a == "y" {
			fmt.Print("input operand 1 : ")
			fmt.Scan(&b)
			fmt.Print("input operator aritmatika : ")
			fmt.Scan(&z)
			fmt.Print("input operand 2 : ")
			fmt.Scan(&c)

			if z == "+" {
				d = b + c
			} else if z == "-" {
				d = b - c
			} else if z == "*" {
				d = b * c
			} else if z == "/" {
				d = b / c

				for e < c {
					d = d * b
					e++
				}
			} else {
				fmt.Print("operator tidak terdefinisi, silahkan mulai dari awal lagi .")
			}

			fmt.Print("hasil")
			fmt.Print(b)
			fmt.Print(z)
			fmt.Print(c)
			fmt.Print(" = ")
			fmt.Println(d)

		} else if a != "y" {
			fmt.Println("Terima kasih telah menggunakan operator arimatika ini")

		}
	}
	fmt.Println(" ")
	fmt.Println("Demikian tugas mengenai aplikasi operator aritmatika")
	fmt.Println("Terimah kasih")
}