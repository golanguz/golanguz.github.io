// pointerlarning ishlatish farqli yo'llari ham bor.

package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x  1
}

/*
new o'ziga bir tur oladi, u turni moslash uchun yetarlicha xotira ajratadi
va u xotiraga pointer(ko'rsatgich) qaytaradi. pointerlar Go'ning
ichiga qurilgan turlar bilan ko'pincha ishlatilmaydi ammo, farqli
ish jarayonlarida juda ham samarali.
*/
