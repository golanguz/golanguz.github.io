// argumenti(parametiri) bor funksiyani chaqirganda(yurgizganda)
// u argument funksiyaga kopiyalanadi

package main

import "fmt"

/*
func zero(x int) {
		x = 0
}

finc main() {
		x := 5
		zero(x)
		fmt.Println(x) // x haliyam 5
}
ushbu programmada zero funksiyasi main funksiyasining ichidagi 5'ni
o'zgartirmaydi, agar biz shuni xoxlaganimizda quydagi kodni yozardik
*/
func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // 0
}

/*
pointers(ko'rsatgichlar) qiymatning o'zi turgan manzilga emas,
qiymat joylashgan manzilga ma'lumotnoma ushlab turadi pointer qullanib
(*int), zero funksiyasi main funksiyasidagi x o'zgaruvchisini
modifikatsiya qila oladi.

* Go'da pointer bu(*) yulduzcha(asterisk) va yuduzchadan keyin tur(*int)
	zero funksiyasida xPtr bu int'ga ko'rsatgich.

	asterisk, o'zgaruvchilardan ma'lumotnomani yechib uchun ham ishlatiladi.
	dereferens qilish ko'rsatgich ko'rsatiyotgan qiymatni foydalanishga
	ruxsat beradi. *xPtr = 0 deb yozganimizda, biz aytiyapmiz,
	"xPtr ma'lumotnoma ushlab turga xotira joyiga int 0'ni joyla"
	agar biz xPtr = 0 deganimizda, to'plamchi xatosi ko'rardik
	chunki xPtr bu int  emas, bu *int, bunga hmma vaqt *int qilib yozishimiz lozim

	va biz & ampersand(&) operatorini o'zgaruvchining manzilini topish
	uchun ishlatiyapmiz. &x qaytimi *int(int'ga ko'rsatgich) chunki x bu int.
	ushbu operator main funksiyasining ichidagi o'zgaruvchini o'zgartirishga
	yordam beradi. main'dagi &x va zero'dagi xPtr bitta xotira manziliga qarashadi.
*/
