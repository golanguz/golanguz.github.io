/*
O'zgaruvchilarni o'rgangandan keyin kichkina programmalar yaratib boshlasag
bo'ladi, Masalan bir programma yozaylik u 1'dan 10'gacha sanashni bilsin
bizdagi bilim bilan biz quydagi kodni yozishimiz mumkin.
*/
package main


import "fmt"

func main() {

	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)
}
// yoki buni

package main

import "fmt"

func main() {

	fmt.Println(`
	1
	2
	3
	4
	5
	6
	7
	8
	9
	10`)
	
}

// Lekin bu ikkisi ham noqulay, Bizga bir narsani bir'dan ko'p marta yurgizadigan
// programma kerak. Misol: For loop

//package main

//import "fmt"

//func main() {

	// i := 1  qiymatlarni ushlash uchun o'zgaruvchi yaratamiz ismi 
// (i) qiymati (1)

	//for i <= 10 {  keyin loop yaratamiz for so'zi bilan keyin
// unga vazifa beramiz bu i < 10'dan oz yoki tengmi, Buning qaytimi boolean
// true(rost) yoki false(yolg'on) keyin forning badanini yaratamiz halqalar bilan

	// fmt.Println(i)  birdan o'ngacha yozib chiqaradi
	// i = i + 1  bu juda ham muhum chunki yuqoridagi for shartimiz
	//}
// bo'lmaganda
// bu to'xtamastan ishlardi va uni infinite loop deyishadi. Biz bu yerda i'ni
// i'ga qo'shyapmiz keying u ikki qo'shilgan i'larga 1'ni qo'shyapmiz, 
// Chunki yuqorida for loop tekshiryapti i <= 10'dan oz tengmi?, 
// Xullas programma toki for loop
// 10'ga teng bo'lguncha ishlaydi keyin esa for loop true habarini olgandan
// keyin terminalga yozishni to'xtatadi.
	

// Biz i o'zgaruvchisini aniqladig uning qiymati 1

// keyin usha i o'zgaruvchisini tekshirdik u i <= 10'dan oz yoki teng u 10'gacha
// tekshiradi for loop'ning  vazifasi i'ni 10'ga yetguncha terminalga yozib
// chiqaradi

// keyin terminalga yozib chiqardik fmt.Println(i)

// keyin i'ga birni qoshib i'ga qaytatdan qoshdik bu degani i'ning qiymati 2 
// bo'ldi, Shunday qilib xisob 10 bo'lguncha davom etdik


// boshqa dasturlash tillarida farqli sintaks ko'rishingiz mumkin 
// u sintaks xam goda ishlaydi misol
	
//	func main() {
//		for i := 1; i <= 10; i++ { 
//			fmt.Println(i)	
//	}
}// bu ham xuddi yuqoridagi for loop'ga o'xshab vazifani bajaradi
// ammo biz bu yerda i++ deb qisqachasiga yozyapmiz unin o'rniga
// i = i + 1 desag ham programmamiz ishlaydi

// if(agar)

// bu safar programmamizni bir oz o'zgartiramiz u nafaqat birdan o'ngacha
// va ularning juft va toqligini yozib chiqarsin
// birinchi bo'lib biz son juft yoki yotqligini aniqlaydigan yo'l yopishimiz lozim
// eng oson yo'li biz buni 2 bo'lishni o'rgatishimiz kerak, agar qoldiq bo'lmasa
// demak bu juft, bo'lmasa toq. Agar esingizda bo'lsa oldingi sahifalardan
// buning uchun qoldiq % belgisini ishlatsak bo'ladi. 1 % 2 1'ga barogar
// 2 % 2 0'ga barobar, 3 % 2 1'ga barobar
// Misol
/*
	if i % 2 == 0 {
		fmt.Println("juft")
	} else {
		fmt.Println("toq")
	}
*/

// if bayonoti for'ga o'xshaydi ammo uning else bayonoti bor, if bayonotida
// agar true chiqsa keyingi block yuruladi agar chiqmasa else bayonot bo'lsa
// else bayonoti yuruladi, va uni uzunroq qilish uchun unda else if bayonoti bor
// Misol
/*
	if i % 2 == 0 {
		ikkiga bo'linadi	
	} else if i % 3 == 0 {
		uchga bo'linadi
	} else if i % 4 == 0 {
		turtga bo'linadi
	}
shartlar yuqoridan pastga yurgiziladi, Agar birinchi boyonot true(rost) bo'lib
chiqsa demak shu bayonot yurgiziladi, va keyingi bayonotlar yurguzilmaydi
hatto ki ularning biri ham ture bo'lib chiqsa, masalan 8 4bilan bo'linadi
va 2 bilan bo'linadi, ammo bu yurgizilmaydi chunki birinchi true(rost) chiqqanda
programma boshqalarni tekshirishni to'xtatadi

butun kod quyidagidek
	
	func main() {
		for i := 1; <= 10; i++ {
			if i % 2 == 0 {
				fmt.Println(i, "juft")	
			} else {
				fmt.Println(i, "toq")
			}
		}
	}
	// o'zgaruvchi yaratdik i uning turi int qiymati 1
	// i 10'dan kam yoki teng? agar rost bo'lsa keyingisiga sakra
	// i ÷ 2 0'ga tengmi? yoq bo'lsa else blockiga sakra
	// i'ni terminalga toq deb yozib chiqar
	// i'ga 1'ni qo'sh chunki i++ yuqorida
	// i 10'dan oz'mi ? ha keyingi block'ga sakra
	// i ÷ 2 0'ga tengmi ? ha if blocki'ga sakra 
	// i'ni terminalga juft deb yozib chiqar
*/
// Switch(almash)
/*
misol biz bir programma yaratmoqchimiz va u programma hamma sonlarni o'zbek
so'zlari bilan yozib chiqarsin bor bilimimiz bilan bi quydagi programmani
yozardik

	if i == 0 {
	  fmt.Println("Nol")
	} else if i == 1 {
	  fmt.Println("Bir")
	} else if i == 2 {
	  fmt.Println("Ikki")
	} else if i == 3 {
	  fmt.Println("Uch")
	} else if i == 4 {
	  fmt.Println("To'rt")
	} else if i == 5 {
	  fmt.Println("Besh")
	}
bu yo'lda programa yozish juda ham noqulay go'da buni ossonroq qilish uchun
switch(almash) bayonoti bor u switch bayonotibilan quydagidek yozardik

	switch i {
	case 0: fmt.Println("Nol")
	case 1: fmt.Println("Bir")
	case 2: fmt.Println("Ikki")
	case 3: fmt.Println("Uch")
	case 4: fmt.Println("To'rt")
	case 5: fmt.Println("Besh")
	default: fmt.Println("Nomalum son")
	}
switch bayonoti switch so'zi bilan boshlanadi undan keyin bajaruv bu (i)
keyin esa bir necha martta case(hodisa) so'zi bajaruvning qitmati case so'ziga
taqqoslanadi agar bajaruv true bo'lib chiqsa demak :'dan keyingi fmt yurgiziladi
xuddi if bayonoti'ga oxshab switch yuqoridan pastga yurgiziladi, true deb teng 
bo'ladigani yurgiziladi, switch bayonotida default suzi xam qullaniladi u degani
agar yuqoridagilarni birontasixam teng rost chiqmasa demak u default bajaruv
yurgiziladi
*/




