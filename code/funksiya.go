/*
funksiyalr nol yoki noldan ko'p paremeterlar(argumentlar) olishi mumkin
va funksiyadagi o'zgaruvchilarning turi hamma vaqt o'zgaruvchidan keyin
yoziladi.
*/

package main

import "fmt"

// agar ikki o'zgaruvchining turi bir bo'ladigan holda quydahi aniqlash
// ham to'g'ri qosh(x, y int)
func qosh(x int, y int) int { // va ushbu funksiya int turini qaytaradi
	return x + y // x bilan y'ni qo'shib qaytar.
}

// va funksiyalarning qaytimini ikkita qilib qaytarsag bo'ladi

func ikki() (int, int) { // ikki (int, int) qaytim
	return 1, 2
}

// endi yuqoridagi qo'sh funksiyasini qullanamiz
func main() {
	fmt.Println(qosh(2, 2)) // 4

	x, y := ikki()
	fmt.Println(x, y) // 1 2
}
