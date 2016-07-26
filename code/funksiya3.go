package main

import "fmt"

/*
var x string = "Salom Gopherbek"

// o'zgaruvchi funksiyadan tashqarida bo'lsa, uni boshqa funksiyalar
// ham qullana oladi
func main() {
	fmt.Println(x) // "Salom Gopherbek"
}
func f() { // ushbu funksiya ham x'ni qullansa bo'ladi
	fmt.Println(x) // "Salom Gopherbek"
}
*/
// ayataylik biz x o'zgaruvchisini farqli funksiya ichid yaratdik

func viloyat() {
	var y string = "Salom Gopherbek"
}

func y() {
	fmt.Pritnln(y) // viloyat fuknsiyasi ichidagi y'ni qullanib bo'lmaydi
}

// hullas {}halqalarning ichidagi aniqlangan o'zgaruvchilarni
// boshqa funksiyalar orqali, yani, o'z aniqlangan funksiyasidan
// boshqa qullanolmaydi
