// array bu uzunligi limitlangan noldan ko'p mahsus turlar ketmaketligi.
// ularning uzunligi limitlangan uchun arraylar ko'p qullanilmaydi
// ularning o'rniga slice, va map, qullaniladi lekin slice va map'ni o'rganish
// uchun biz oldin arayni o'rganishimiz lozim. arraylar quyidagidek aniqlanadilar:

package main 

import "fmt"

func main() {
	var x [4]int // qiymati 4'lik intejer arrayi
	fmt.Println(x) // arayning ichidagi birinchi qismni terminalga yoz.
// array indeksatsiyasi(sanalishi) to'rt-burchak qavs orqali [] 0 dan boshlanadi,
// bu degani arrayning 1'chi qismi bu [0], va oxirgisi [-1].
// arrayga qiymat yozmoqchi bo'lsangiz quyidagidek
	x[0] = 1
	x[1] = 2	
	x[2] = 3
	x[3] = 4
	fmt.Println(x) // [1 2 3 4]


// array'ning aniqlashning boshqa yo'llari quyidagidek:
	y := []int{1,2,3} 
	fmt.Println(y) //[1 2 3]

// ikkinchi array da biz qiymatni birdaniga yozib aniqliyabmiz

// array literal orqali qiymatli array aniqlasak ham bo'ladi quydagidek:

	var z [5]int = [5]int{1,2,3}
	fmt.Println(z[1]) // 2

/*agar yuqoridagi x arrayga yana qiymat yozmoqchi bo'lsangiz to'plamchi xato
	x[0] = 1
	x[1] = 2	
	x[2] = 3
	x[3] = 4
	x[4] = 5
	fmt.Println(x) //invalid array index 4 (out of bounds for 4-element array)
*/

// arraylarning ustidan for loop bilan o'tsak bo'ladi.
	for k, s := range x { // range loop boshqa darsda o'rganamiz
		fmt.Println(k,s) // 0 1 
	}			 // 1 2
				 // 2 3
				 // 3 4
// array, slice, maplarning ustinda loop qilish uchun faqat loop ishlatiladi
}

