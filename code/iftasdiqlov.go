/*
if(agar) statement(tasdiqlov) if tasdiqlovchisi programmadaa "agar" manosini 
bildiradi, Misol programmangiz bir shar bajarishi lozim aytaylik qo'shish.
2 + 2 = 4, va shar to'g'ri bo'ladigan bo'lsa siz xoxlaysiz ushbu if tasqilovchisi
terminalga to'g'ri deb yozsin, agar to'g'ri bo'lmasa noto'g'ri deb yozsin.
if statementi quyidagidek yoziladi:
*/
// go yo'lchasidagi src quttisida yangi qutti yaratamiz uning nomi if-statement
// va u quttining ichida main.go faylini yaratamiz va fayl ichida quydagi kod:

package main


import "fmt"

var qosh int = 2 + 2 // o'zgaruvchi qosh intejer qiymati 2+2. 2'ni 2'ga qo'shish

func main() {
	
	if qosh == 4 { // agar qosh o'zgaruchisi 4'ga teng bo'lsa
		fmt.Println("2 + 2 = ", qosh, "to'g'ri") // ushbu habar yozilsin
	} else { // teng bo'lmasa
		fmt.Println("noto'g'ri") // ushbu
	}

/*
ko'rib turganingizdek if statementini ishlatib biz tekshiryapmiz qosh 4'ga tengmi
yo'qmi, va ikki teng belgisini ishlatyapmiz, dasturlashda huddi shunday ikki
tenglik belgisi ishlatiladi, va yagona tenglik belgisining ma'nosi bu assignment
(yozish). agar esingizda bo'lsa boolean sahifasida bularni qullangan edik a==b.
if dan keyin biz else so'zini qullaniyapmiz, else(bo'lmasa) degani, 
else javob ummuman topilmaganda ishlatiladi, ammo ifdan keyin else if, yoki else albatta bo'lishi lozim aks xolda xato.
if(agar) qosh == 4'ga teng bo'lsa "to'g'ri" habarini yoz, bo'lmasa else'ni
qullanib "noto'g'ri". bularning uchunchi si ham bor bu (else if)agar yana.
bu if tasdiqlovchisini uzunlashtirishga yordam beradi. Misol aytaylik
biz for loop qullanib 0'dan ikkigacha sanyapmiz va ushbu for loopining ichida
if tasdiqlovchisini yaratiyapmiz va unga aytiyapmiz if(agar) i bir bo'lsa 
terminalda "bir" habarini yoz, else if(agar yana) i ikkiga teng bo'lsa
"ikki" habarini yoz, quyidagidek
*/
	for i := 1; i <= 2; i++ { // i <= 2'dan kam yoki tengmi? 
		if i == 1 { // i bir bo'lsa "bir" habari
			fmt.Println("bir")		
		} else if i == 2 { // i ikki bo'lsa "ikki"
			fmt.Println("ikki")		
		}
	}
}
// if statementi yuqoridan pastga qarab o'qiladi(yurigziladi) 
// agar if statementinig birinchi bloki(halqalar ichidagi kod) to'g'ri bo'ladigan
// bo'lsa demak u yurgiziladi va boshqalari yurgizilmaydi, birinchi to'g'ri
// blokdan keyingi blok  to'g'ri bolsa ham yurgizilmaydi.py
// go run main.go buyruqini yurgizsangiz quydagi habarni ko'rishingiz lozim:
// 2 + 2 =  4 to'g'ri
// bir
// ikki

