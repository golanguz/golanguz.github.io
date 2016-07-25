/*
map bu tartibsiz asosiy-qiymat juftligi, map boshqa dasturlash tillarida
dictionaries("kutubxonalar"), va hash tables("chopilgan jadvallar") deb atalishadi
maplar qiymatni unga bog'liq bo'lgan asosi'ni qidisrish uchun ishlatiladi
*/
package main

import "fmt"

func main() { //assosiy qiyimat juftligi[string] va uning turi int intejer
	aktorlar := map[string]int{
		// assosiy qiymat juftligi "Nicolas Cage". tur in 50
		"Nicolas Cage":     50,
		"Scarlet Johanson": 29,
		"Jude Law":         41,
	}
	fmt.Println(aktorlar)
	fmt.Println(aktorlar["Nicolas Cage"]) // 50
	// aktorlarning qiymatlarini qullanish to'rt burchak qavs bilan
	// agar yuqoridagi literal map yaratish ishlatmasangiz u xolda
	// make funksiyasini ishlatishingiz lozim

	ak := make(map[string]int)
	ak["Jennifer Lawrence"] = 1990
	fmt.Println(ak["Jennifer Lawrence"])

	/*
					  quydagi kod xato
					  var x map[string]int
					  x["key"] = 10
					  fmt.Println(x)
					  panic: runtime error: assignment to entry in nil map
				    chunki maplar qullanishidan oldin yurguzilishi(boshlanishi) lozim
		        ushbu kodni quydagidek yozish lozim
	*/
	var x = make(map[string]int)
	x["key"] = 10
	fmt.Println(x)        // map[key:10]
	fmt.Println(x["key"]) // 10

	// maplarning assosini ham qiymatini ham int yoki float64 yoki string yoki
	// yangi map qilsangiz bo'ladi

	// misol programma assosi va qiymatini string qilamiz

	string := make(map[string]string)
	string["bir"] = "ikki"
	string["uch"] = "to'rt"
	fmt.Println(string["bir"], string["uch"])

	// map qiymat turi uchun 0 nol qaytaradi va stringlar uchun bu "" bo'sh string.
	// va biz mapning ichidagi qiymatlarni tekshirsak bo'ladi quyidagidek
	ism, ok := string["bir"]
	fmt.Println(ism, ok) // ikki true(to'g'ri) chunki haqqatan ham "bir" assosiy qiymati bor
	// ushbu kodning boshqa ossonroq yo'li ham bor quydagidek

	if ism, ok := string["bir"]; ok {
		fmt.Println(ism, ok) // ikki true
	}
	// oldin biz map'dan qiymat olishga harakat qilamiz. agar ushbu "bir" qiymati
	// mapning ichida bo'lsa halqalar ichidagi kod fmt.Println(ism, ok) yurgiziladi
	// aks xolda hato

	// yoki intejerlik map
	intejer := make(map[int]int)
	intejer[1] = 2
	intejer[3] = 4
	fmt.Println(intejer[1], intejer[3])

	// maplarning ichida map[string]string
}
