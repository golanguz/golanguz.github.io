/*
slicelar bular oddiy arrayning parchasi, huddi arayga o'xshab bular
indeksatsiyali va uzunligi bor, va ularning uzunligini uzaytirsa
bo'ladi, chunki slice(slays) bosh arayga ma'lumot(ko'rsatgich),
bu degani slice, oddiy arayning ustida qurilgan, agar slice'ga slice
yozsangiz u ikkisi ham ustida qurilgan arrayka ma'lumot ushlab turadi
tushunmagan bo'lsangiz Misol:
telefoningizdagi bir insonga qarashli 2ta nomer.
sllayslar bosh array'ga o'rov.
*/
package main

import "fmt"

func main() {
	// []T bu turi T bo'lgan slice, bizning misolda u int
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s) // [1,2,3,4,5]

	// slice(tarjimasi "kesish") huddi o'zining ismi tarjimalangandek
	// slayslarni kessa bo'ladi quyidagidek

	fmt.Println(s[1:3]) // [2 3] [past : balan] pastdan yuqoriga
	fmt.Println(s[:3])  // [1 2 3] past yo'q bu degani pas qiymati 0
	// va bu boshidan kesishni boshlaydi
	fmt.Println(s[4:]) // [5] chunki balandligi yo'q 4'chi dan boslab
	// oxirigacha kesadi.
	//[past] bu kesishni boshlash indeksi va [baland] bu indeksni
	// tugashi ammo indekslar o'zlari sanalmaydilar
	// boshlayotgan indeksingiz sanalmaydi va tugayotgani ham.

	// slicelarni birdaniga qiymati bilan yartgandan ko'ra
	// make funksiyasini ishlatib bo'sh slice yaratsangiz
	// bo'ladi va uning qiymatlarini keyinroq yozsangiz bo'ladi.

	hayvonlar := make([]string, 3) //hayvonlar va ularning uzunligi 3
	hayvonlar[0] = "kuchuk"
	hayvonlar[1] = "mushuk"
	hayvonlar[2] = "sichqon"
	fmt.Println(hayvonlar)
	// make bilan yaratilgan slice bor yo'qi nollik(0) array
	// qaytariyotgan slice va shu qaytiyotgan slice o'sha qaytariyotgan
	// arrayga ma'lumot ushlab turibdi

	// slice'ning yordamchi funksiyalari ham bor ular append va copy
	// append(qo'shish) slice'ning oxiriga qiymat qo'shadi
	// copy(ko'chirish) slicening qiymatini boshqa slice'ga ko'chiradi

	shaharlar := []string{"Samarqand"}
	shaharlar2 := append(shaharlar, "Toshkent", "Buxoro")
	fmt.Println(shaharlar, shaharlar2) // [Samarqand] [Samarqand Toshkent Buxoro]
	/*
							   aytib o'tgandek append funksiyasi slaysning oxiriga elementar qo'shadi
							   agar bosh arrayda yetarli imkoniyat bo'lsa element slaysning oxirgi
							   elementidan keyin qo'yiladi va uning uzunligi uzayadi.
							   agar uning uzunlik imkoniyati bo'lmasa, yangi array yaratiladi
							   va hamma yaratilgan elementlar oxirga qo'shilib kopiyalanadi.
						     va yangi slice qaytariladi xuddi shunday [Samarqand] [Samarqand Toshkent Buxoro]

				         yuqoridagi kod yurgizilgandan keyin append yaratilgan slaysni olib
		             ("Samarqand") unga ikkinchi slaysni qo'shadi ("Toshkent", "Buxoro")
	*/

	// slice'ga slice qo'shsangiz ham bo'ladi
	boshqaShahar := []string{"Shaxrisabz", "Andijon", "Xorazm"}
	shaharlar = append(shaharlar2, boshqaShahar...) //ellipsis qullanyapmiz
	fmt.Println(shaharlar)

	/*
							  ellipsis bu Go'ning ichiga qurilgan xususiyat, ellipsis bu degani
							  elementlar to'plami. biz string turini []string turiga qo'sholmaymiz
						    ammo slice'ning oxirida elipsis qullanib biz buni qo'shiyapmiz
				        biz Go to'plamchisiga ko'rastiyapmiz biz slaysimizning elementlarini
				        bir biriga qo'shmoqchimiz, shuning uchun to'plamchi bizga xato bermaydi
				        chunki biz bir string turli slaysdan boshqa slaysga string qiymatlarini
				        qo'shiyapmiz. ammo turi []int bo'lgan slaysni turi []string bo'lgan
		            slicega qo'sholmaysiz ummuman. to'plamchi baqiradi.
	*/
	// xoxlagan vaqt slaysning uzunligini tekshirib ko'rishingiz mumkin len() funksiyasi bilan
	fmt.Println(len(shaharlar))
	shahar := make([]int, 6)
	fmt.Println(len(shahar))

	//slaysning 0 nol qiymati bu nil. nil slaysning uzunligi va imkoniyati 0
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	/*
		   copy funksiyasi esa ko'chirish ishini bajaradi. copy funksiyasi
		   ikkita ragument ishlatadi dst va src. src'nin hamma qiymatlari
		   dst'ga ko'chiriladi, va distdagi hamma qiymatni o'chirib ustidan yozadi
			 agar ikki slaysning uzunligi bir bo'lmasa, eng qisqa uzunlig slays ishlatiladi
	*/
	slays := []int{1, 2, 3}
	slays2 := make([]int, 2)
	copy(slays2, slays)
	fmt.Println(slays, slays2)

	/*
		programmani yurgizgandan keyin slays [1,2,3] va slays2[1,2] ko'rasiz
		slays'ning qiymati slays2'ga ko'chirilgan, ammo slays2 faqat 2'ta qiymat
		uchun joyi bo'lagni uchun, slays'ning faqat 2'ta qiymati ko'chirilgan
		slayslar odatda ko'p miqdorda parchalarni anglatadi va shu uchun qullanadi
		misol o'yinchi #33, yoki 10'chi eng ko'p qidirilgan internetda narsa
	*/
}
