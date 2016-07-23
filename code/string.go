/*
String aniq uzunligi belgilari ketma ketligi matnni ko'rsatish uchun ishlatiladi.
Go stringlari individual bayt orqali tashkil etilgan, odatda har bir bayt uchun
bir belgi(xitoy belgilari uchun esa birdan ko'p bayt ishlatiladi).
String literallari qo'shaloq kuvota("") "Salom, Gopherbek" bilan yaratiladi yoki
backtick bilan (``) `Salom Gopherbek` bularning farqi qo'shaloq "" kuvotalar
yangi chiziq ishlatolmaydi, ammo ular mahsus belgi qullanadilar Misol \n
yangi chiziqqa tushiradi. va \t tab belgisi bu, Belgilar o'rtasida 8'ta 
bo'shlik yaratadi. Misol stringlar qullangan programma quydagidek.
*/

package main


import "fmt"

func main() {
	
	fmt.Println("Salom, Gopherbek") // oddiy xabar 
	fmt.Println("Salom\t", "Gopherbek") // tab(8'ta bo'shlik) qo'shilgan
	fmt.Println("Salom " + "Gopherbek") // ikki string qo'shilgan 
	fmt.Println("Salom\n", "Gopherbek") // yangi chiziqqa tushirilgan
	fmt.Println(`Salom,	
				Gopherbek`) // backtick bilan yangi chiziqqa
	fmt.Println(len("Salom, Gopherbek")) // belgilar uzunligini hisoblash
// len() funksiyasini qullanib belgilarni uzunligini saniyapmiz qo'shaloq
// kuvotalarni ichida bo'shlik belgi deb hisoblanadi.
// agar string'ning ichidagi belgilardan birini terminalga yozmoqchi
// bo'lsangiz
	fmt.Println("Salom, Gopherbek"[1]) // deb yozishingiz mumkin ammo 
// terminalda xarf o'rniga son(97) ko'razis chunki:
// qo'shaloq kuvotalar ichidagi beliglar ular multi-byte(ko'plik-bayt) UTF-8
// kodlash individual belgilari qullaniyotgan bo'lishi mumkin, 
// va byte bu intejer(son). 
//Demak string'ning ichidagi belgini yozib chiqarish uchun:

	fmt.Printf("%c\n", "Salom, Gopherbek "[1]) // a xarfini terminalga yozadi
// chunki string indeksiyatsiyasi 0'dan boshlanadi quydagidek
// esingizda bo'lsa  bo'shlik xam sanaladi
// "S   a   l   o   m   ,     G   o   p   h   e   r   b   e   k" 
//  0   1   2   3   4   5  6  7   8   9   10  11  12  13  14  15
//
//      bir sonini a yozib chiqaradi.

	fmt.Println(string([]rune("Salom, Gopherbek")[1])) // UTF-8

// bular haqida  kechroq o'rganamiz

// yuqoridagi len() funksiyasi 1234.... deb sanashni boshlaydi
// chunki biz string'ning ichidan xarf ajratiyotganimiz yo'q ularning miqdorini
// bilmoqchimiz.
}
