//Stringlar

package main

import "fmt"

func main() {
	fmt.Println("Hello, World") // Hello, World yozib chiqariladi
// ikki ust belgi(") bu string deyiladi
// va bu so'zlarni yozib chiqarish uchun ishlatiladi
	fmt.Println("Hell'o, World")// Hell'lo world
// ikkitalig belgini ichida bittali belgi ishlatsa bo'ladi ammo ikkitalikning
// ichida ikkitalik bo'lamydi. Misol "Hell"o World" xato bo'ladi
// ammo backtickning (`)ichida xoxlaganday ishlatsa bo'ladi. Misol
	fmt.Println(`He"ll'o, Wo"rl'd`) // He"ll'o, Wo"rl'd
	fmt.Println(`He"ll'o, 
			      Wo"rl'd`) // backtick bilan bunaqasi ham bo'ladi

// string ichidagi sonlarni tengligini ko'rsak ham bo'ladi
	str := "2"
	ikk := "3"
	fmt.Println(str == ikk) // false

// stringlarni bir biriga qo'shsak ham bo'ladi
	str1 := "birinchi string "
	str2 := "ikkinchi string"
	fmt.Println(str1 + str2) // birinchi string ikkinchi string

// stringlarni o'z chizig'iga tushursak ham bo'ladi
	str3 := "birChiziq\nikkChiziq" // \n stringning ichidagi har bir 
// narsani o'z chiziqiga qoyadi.
	fmt.Println(str3)// birChiziq
			 // ikkChiziq

// Go ichiga qurilgan len funksiyasi bilan string harf sonini bilsangiz bo'ladi
	fmt.Println(len(str3)) // 19
}
