package main

import "fmt"

func far(f float64) float64 { // farengeytdan tselsiyusga konverter funksiyasi
	 return (f-32) * 5 / 9 // hisoblash bajarmasi
}

func main() {
	
	fmt.Print("Son yozing: ") // terminalda habar

	var kirish float64 // terminalda yoziladiga float64 sonini ushlaydi

	fmt.Scanf("%f", &kirish) // kirishga yozilgan sonni skaner qiladi
	
	chiqish := far(kirish) // yangi o'zgaruvchi chiqish.
// yuqoridagi far funksiyasini ichiga kirish'ni qoyib ikkisini chiqishga
// joyliyaymiz
// misol terminalda 100 ni yozsangiz 37.77777777 javobini olasiz, shu 100 soni
// kirish o'zgaruvchisiga kiradi, keyin skanner qilinadi, keyin far funksiyasi
// ichiga joylanadi u funksiya 100 sonidan  32'ni ayiradi(68).
// keyin 5'ni 9'ga bo'ladi(0.5555555) keyin ikkisining qiymatini ko'paytiradi
// 0.55555 * ko'paytiramiz 68 = 37.777777. keyin bu javob(qiymat) chiqish
// o'zgaruvchisiga joylanadi va Println funksiyasi orqali terminlaga yoziladi.	
	fmt.Println(chiqish)
}
/*
Harakat qilin bitta shunday programma yozing va u selsiyusdan farengeytga konvert
qilsin: Misol yuqoridagi kodni qisqa yo'li ham bor.

	fmt.Print("Son yozing: ")
	var kirish float64
	fmt.Scanf("%f", &kirish)
	chiqish := (kirish -32) * 5/9
	fmt.Println(chiqish)

huddi yuqoridagi kod'dek bu ham farengeytdan selsiyusga konvert qiladi.
*/
