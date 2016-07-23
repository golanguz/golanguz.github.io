// variable(o'zgaruvchi) bu mahsus turlik sistema xotirasining joyi.
// oldingi yozgan salom gopherbek programmasini o'zgartiramiz,
// va ushbu salom gopherbek stringini o'zgaruvchiga joylaymiz.

package main


import "fmt"

var (
	x = 1
	y = 2
	z = 3
	q = "Qaley>?"
	boolean = true
	funksiya = func(x,y int) int{return  x + y }  // anonim funksiyasi
)

func main() {
	
	 // h bu habar so'zinig birinchi harfi uchun ishlatilgan
	   // string bu o'zgaruvchi h'ning turi
	var h string = "Salom Gopherbek" // bu esa qiymat(value)

	fmt.Println(h) // h'ning "Salom Gopherbek" qiymatini terminalga yozamiz
// Xabarni to'g'ri Println funksiyasini ichiga yozishni o'rniga uni,
// h o'zgaruvchisiga joyliyabmiz.
// o'zgaruvchi aniqlayotgan zamon uning qiymatini aniqlash u sizga bog'liq
// yani qiymatsiz xam yaratsak bo'ladi quydagidek:
	var habar string // qiymatsiz
	habar = "Salom Gopherbek" // qiymat joylandi
	fmt.Println(habar) // terminalga yozildi
// o'zgaruvchilar xuddi nomidek ular o'zgaradi misol
	habar = "Salom Shavkat" // habar o'zgaruvchisining qiymatini o'zgartirdik
	fmt.Println(habar) // "Salom Shavkat" yozib chiqaradi

// funksiyalarning ichida qisqa yo'l bilan o'zgaruvchi aniqlash mumkin:

	a := "Salom"
	b := "Dunyo"
	fmt.Println(a, b) // "Salom Dunyo"

// qisqa aniqlash faqat funksiyalar ichida ishlaydi va funksiyadan tashqarida
// aniqlash xato beradi. Yuqoridagi a va b o'zgaruvchilari avtomatik xolatda
// Go to'plamchisi ularga tur beradi va u "string" turi chunki a va b'ning
// qiymati string.

// Agar birdaniga ko'p o'zgaruvchi aniqlamiqchi bo'lsangiz, ularni qavs ichiga
// yozasiz xuddi yuqoridagi funksiyadan tashqaridagi o'zgaruvchilarga o'xshab
	
	fmt.Println(x, y, z, q, boolean, funksiya(2,3)) // 1 2 3 Qaley>? true 5
// ko'rib turganingizdek funksiyadan tashqaridagi var zo'sining ichida
// bir dan ko'p o'zgaruvchi aniqladig son, string, boolean, va anonim funksiya
// bu anonim funksiyasi bor yo'g'i ko'rgazma va buni biz funksiyalar'ga 
// qarashli sahifada o'rganamiz.

}
