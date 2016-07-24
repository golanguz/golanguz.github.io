//for tasdiqlovchisi bir kodning parchasini birdan ko'p martta yurgizadi

package main


import "fmt"


func main() {

	i := 1 // o'zgaruvchi aniqlaymiz uning qiymati 1
	for i <= 10 { // for saqlangan so'zi bilan loop(halqa) yaratamiz
	fmt.Println(i) // i'ni terminlaga yozib chiqaramiz
	i = i + 1 // o'ng tomondan i'ga 1'ni qo'shamiz va qayta ayga yozamiz.
}

/*
o'zgaruvchi yaratamiz uning qiymati 1
so'ng for loop orqali tasdiqlovchi kod parchasi yaratamiz <= 10 va loop
aynalishni boshlaydi.

i <= 10 i o'ndan kammi?  to'g'ri kam

fmt.Println(i) i'ni terminalga yozadi

1'ni i'ga qo'shadi va bir qo'shilgan i'ni  qayta i'ga qo'shadi, qiymat 2

i <= 10 i 10'dan kammi ? to'g'ri kam

fmt.Println(i) i'ni terminalga yozadi

1'ni i'ga qo'shadi va bir qo'shilgan i'ni  qayta i'ga qo'shadi, qiymat 3

va shunday davom etadi toki i 10'ga teng bo'lguncha, teng bo'lganda to'taydi
*/
	for j := 0; j <= 10; j++ { // dasdiqlovchilar shartli ifodaning ichida 
		fmt.Println(j) 
	}


// ikkinchi for loop ham huddi birinchidek faqat dasdiqlovchilar shartli
// ifodaning ichida, Va i = i + 1'ning o'rniga i++ ishlatiyapmiz
// dasturlashda bu i = i + 1 juda ham ko'p ishlatiladi shuning uchun
// dasturchilar qisqa yo'lini o'ylab topgan i++'ni i--'si ham bor bu faqat 
// chappachasiga xisoblaydi 10'dan pastga qarab ketadi quydagidek:

	for c := 10; c >= 0; c-- {
		fmt.Println(c)
	}

/*
uchunchi for loopda c'ning qiymati 10, va keyin biz tekshiriyapmiz c >= 0
c 0 nolga tengmi ? yo'q teng emas. c-- c'dan(qiymati 10) birmatta ayirim
ishlatyapmiz yani 1 sonini va uning qiymati 9 bo'lyapti, keyin loop qilib
aynalib kelyapmiz, yana tekshiryapmiz. toki c'ning qiymati 0 nolga teng 
bo'lmaguncha loop ishlaydi.
	terminalga quydagi habarni yozib chiqaradi:
		1
		2
		3
		4
		5
		6
		7
		8
		9
		10
		0
		1
		2
		3
		4
		5
		6
		7
		8
		9
		10
		10
		9
		8
		7
		6
		5
		4
		3
		2
		1
		0

*/

}


