

// O'zgaruvchilarni aniqlash.

package main // agar python'da dasturlagan bo'lsangiz, pythonga o'xshash paket
// nomlash


import "fmt" // Go'ning ichki kodidan fmt degan paketini import qilamis va,
// Keyin pastda uning("fmt"ning) ichidan "Printf" funksiyasini qullanib
// farengeyt issiqligini terminalga yozib chiqaramiz.


const farengeyt = 212.0 // bu yerda const yani(o'zgarmaydigan) aniqlaymiz va,
// Uni farengeyt deb nomlaymiz uning qiymayi esa "float64" 212.0
// keyin main() funksiyani ichida buni o'zgaruvchi var f'ga butunlay
// topshiramiz



func main() { // funksiya main() faqat to'planadigan programmalar uchun lozim
// va muhim, Butun to'planadigan programmangizda bitta main paket va uning
// bitta main() funksiyasi bo'lishi lozim. Agar kutubxona yoki paket
// yaratiyotgan bo'lsangiz main() funksiyasiz xam bo'ladi.

	var f = farengeyt // Bu ayni funksiyadan tashqaridagi farengeyt
// o'zgarmaydigan constant buning qiymayi 212.0.		


	var c = (f - 32) * 5 / 9 // Bu yerda esa biz yangi c o'zgaruvchi 
// aniqlaymiz, Keyin uni qiymatini f deb aniqlaymiz qavs ichida f'dan 
// 32'ni ayramiz("qavslarni Go to'planchisi birinchi bo'lib yurguzadi")
// bu degani biz birinchi 212.0'dan 32 ayiramiz bu = 180 bo'ladi
// keyin u 180 qiymatni 5'ga ko'paytiramiz bu = 900 bo'ladi
// va oxirda 900'ni 9'ga bo'lamiz bo = 100 bo'ladi


	fmt.Printf("farengeyt issiqligi %g°F yoki selsium'da %g°C\n", f, c)
// bu yerda esa yuqoridagi fmt paketini import qilib uning Printf funksiyasini
// qullanib to'planchi bajargan hamma qushish va bo'lishlarni terminalga yozib
// chiqaramiz. Ushbu programmani yurgizsangiz terminalga
// farengeyt issiqligi 212°F yoki selsium'da 100°C degan habar ko'rsatiladi.

}

// O'zgaruvchilar

// var so'zi oddiy bir o'zgaruvchi turini yaratadi, Va uning ismini aniqlaydi,
// Uning qiymatini ham aniqlaydi
	var ism tur = qiymat
// tur yoki qiymat'siz o'zgaruvchilar aniqlash mumkin, Agar tur'siz aniqlanadigan
// bo'lsa o'zgaruvchining turi qiymat orqali aniqlanadi, Agar qiymatsiz aniqlansa
// o'zgaruvchining qiymati sonlar uchun 0 bo'ladi, false booleanlar uchun,
// va "" bo'sh string stringlar uchun, nil interfacelar va referens turlari
// (slice, pointer, map, channel, function) uchun.

// nol-qiymatli mexanizm o'zgaruvchi'ning o'zining turiga yarasha qiymati
// borligini aniqlaydi
// Misol
	var s string // bu qiymatsiz lekin yuqorida aytib utilgandek
// Go to'plamchisi bunin qiymatini turiga qarab aniqlaydi va bizda xozir 
// s deb nomlangan o'zgaruvchi bor uning turi string bu degani uning qiymati
// "" bo'sh string 
	fmt.Println(s) // s'ni chaqiramiz va yurgizami terminalda xato emas
// boshlik ko'ramiz.

// Bir o'zgaruvchida bir dan ko'p qiymatlar aniqlasak bo'ladi

	var a, b, v int // a,b,v uchisining turi // int, int, int
// Xuddi shunday qilib agar bularni tur'siz aniqlasak ular o'zlari aniqlaydi
	var a, b, v = true, 2.2, "uch" // bool, float64, string bo'ladi

// main() funksiyasi bo'lgan faylingizda xar bir o'zgaruvchi main'dan oldin 
// aniqlanishi lozim


var (
	ism tur = qitmat
	ism tur = qiymat
	ism tur = qiymat
	ism tur = qiymat
// bunaqasiham bo'ladi
)

// qisqacha o'zgaruvchi aniqlash bular funksiyalarning ichida ishlatiladi
// va bularning turi ularning qiymati orqali aniqlanadi
// Misol
	func qisqaOzgaruvchilar() {
	
	menQisqaOzgaruvchi := "salom"
	son := 0.0
// masalan funksiyanin ichida aniqlangan qisqa o'zgaruvchilar 
// aniqalangan bo'lsa, keyingi aniqlanyotgan o'zgaruvchilar aniqlanmaydi
// Misol
	bir, ikki := "bir aniqlandi, ikki aniqlandi"
	uch, ikki := "uch aniqlandi, ikki yuqoridagi ikkiga yopishtirildi"
// qisqa o'zgaruvchilar kamida bitta yangi o'zgaruvchi aniqlashi lozim,
// Misol 
	tort, besh := "zo'r gap yo'q"
	tort, besh := "xato to'plamchi buzuldi"		
// buni to'g'rilash uchun ikkinchi aniqlanuvchini oddiy o'zgaruvchi qilib
// aniqlang
	}





































