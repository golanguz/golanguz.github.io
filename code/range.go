/*
for loop'nin range formasi slays, va map ustilaridan aynaladi.
slays va map'ning hamma qiymatining ustidan aynaladigan qobiliyati
bo'lgani uchun, range, qiymatlarning ustidan aynalishni juda ham qulay qiladi
*/

package main

import "fmt"

var sonlar = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for indeks, qiymat := range sonlar {
		fmt.Println(indeks, qiymat)
	}
	/*
		  kodni yurgizsangiz quydagi habarni ko'rasiz
				0 1
				1 2
				2 4
				3 8
				4 16
				5 32
				6 64
				7 128
		  ushbu habar qiymatlarning indeksi bilan ko'rsatilgan ammo
		  indeks'siz yozib chiqarsangiz ham bo'ladi uning uchun _, ishlatiladi
	*/
	for _, qiymat := range sonlar { // _, indeksini aniqlamashlik
		fmt.Println(qiymat)
	}
	/*
				  1
		  		2
		  		4
		  		8
		  		16
		  		32
		  		64
		  		128
		  xoxlasangiz yimatni o'rniga _, ishlatishingiz mumkin.
	*/
	/*
	  agar kerakli qiymatni topibganda slays ustidan loop qilishni
	  to'xtatmoqchi bo'lsangiz break so'zini ishlatasiz
	*/

	toxtash := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range toxtash {
		fmt.Println(i)
		if i >= 5 {
			break
		}

	}
	/* kodni yurgizsangiz quydagi habarni ko'rasiz
			0
			1
			2
			3
			4
			5
	  ko'rib turganingizdek to'xtash slaysimizning qiymati o'nta intejer
	  va biz uning ustidan range qilib o'tyapmiz va keyin esa aytiyapmiz
	  agar to'xtashning ustidan loop qilayotgan i o'zgaruvchisi 5 sonidan
	  katta yoki teng bo'lsa "break"(to'xta) deb aytiyapmiz.
	*/
	// continue ham huddi shunday ishlaydi ammo u o'tkazib yuborish ishlatadi

	otkazish := make([]int, 10)
	for i := range otkazish {
		fmt.Println(i)
		if i >= 5 {
			break
		}
	}
	/*
	  0
	  1
	  2
	  3
	  4
	  5
	*/

	/*
		range'ni map'ning ustidan ham loop qilish uchun ishlatsangiz bo'ladi
		u xolda map'nin birinchi parametiri bu intejer emas, so'z
	*/

	shahar := map[string]int{
		"Samarqand": 0,
		"Buxoro":    1,
		"Termiz":    2,
	}
	for soz, qiymat := range shahar {
		fmt.Println(soz, qiymat)
	}

	/*
		Samarqand 0
		Buxoro 1
		Termiz 2

	*/

}
