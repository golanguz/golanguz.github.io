// aytib o'tgandek maplarning qiymatlari farqli bo'lishi mumkin
// ushbu map2 darsida bis stringlik map yaratamiz va uning qiymati
// stringlik map bo'ladi quydagidek:

package main

import "fmt"

func main() {
	// birinchi map so'zi bu tashqi map. ikkinchisi ichgi.
	moshin := map[string]map[string]string{ // map literal yaratdig make'siz
		"N": map[string]string{ // tashqi mapga qarashli.
			"nom": "nexia",
			"gaz": "benzin",
		},
		"M": map[string]string{ // tashqi mapga qarashli.
			"nom": "malibu",
			"gaz": "benzin",
		},
	}
	if nom, ok := moshin["N"]; ok {
		fmt.Println(nom["nom"], nom["gaz"]) // nexia benzin
	}
	/*
					  ko'rib turganingizdek biz stringlik map yartdik va uning ichida yana
					  bitta stringlik map bor va uning qiymati string, tashqi map faqat qidiruv
				    uchun ishlatiladi, "N" yoki M so'zi uchun ishlatiladi. va ichki map
		        "nom": "malibu",
		        "gaz": "benzin",
		        informatsiyasini saqlash uchun ishlatiladi.
	*/
}
