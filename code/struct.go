/*
struct bu jam to'plamchisi, u nol yoki nolda ziyod nomlangan qiymatlarni
bitta idorada ushlaydi, har bir qiymatning nomi field(maydon)
*/

package main

import "fmt"

type IsmSharif struct { // turi IsmSharif degan struct yaratdik
	Ism  string // birinchi maydoning turi string nomi Ism
	Yosh int    // ikkinchi maydoning turi int nomi Yosh
}

func main() {
	ismsharif := IsmSharif{Ism: "Savkat", Yosh: 1989} // IsmSharif
	// structini ismsharif o'zgaruvchisiga yoziyapmiz
	fmt.Println(ismsharif) // va terminalga chiqaryapmiz
}
