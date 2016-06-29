/*
Go bu statik dasturlash tili. Bu degani o'zgaruvchilar'ning hamma vaqt turalari 
bor, Va u tur hech qachon o'zgarmaydi, Statik yozish bu boshida no qulay
bo'lishi mumkin, Ammo statik yozish bizga programmamizni yaxshi tushunishga
imkoniyat beradi.

Birnchi sahifada aytib utganday Go'da saqlangan soz'lar bor, Shular xaqida o'rganamiz
*/
package main

import ("fmt")

// Sonlar
// Asosan go da sonlar uchun ikki xil turlar bor.
// Ular integers va floating-point sonlar

// #Integerlar bular nuqtasiz va minus'siz bo'lgan sonlar (-2,3.4,-4)
// Go'da to'rt xil uzunlik integerlar bor ular: 8, 16, 32, va 64 bit
// Ular o'zlarining tularibilan namoyish etiladi va 
// turlari: int8, int16, int32, int64
// va O'ziga xos yozilmagan integerlar bor ular:
// uint8, uint16, uint32, uint64
// int va uint deganlari xam bor bular uchun uzinligini aniqlash shart emas.
// eng ko'p ishlatiladigani bu int bularning ikkisi xam, 32 yoki 64 bit bo'ladi
// ammo siz fol ochishingiz shart emas, Chunki to'plamchilar farqli to'playdilar
// va int bilan int32 bular ummuman farqli turlar, int bu -128 127'ga,
// int32 esa 0 255'ga

// Quydagilar goning arifmetik operatorlari

	*	/	%	<<	>>	&	&^
	+	-	|	^
	==	!=	<	<=	>	>=
	&&
	||	

// operatorlar chap tomondan boshlanadi(xisoblanadi) Agar shart qavs ichida
// va o'ng tomonda bo'lsa demak birinchi qavs ichidagi shart yurgiziladi
// misol, qo'shish & (2+2) birinchi qavs ichi xisoblanadi.

func main() {
	  // (*) Ko'pincha ko'paytirish manosini bildiradi, Ko'rsatgichlar
	  //bilan xam ishlatiladi. 
				fmt.Println(2 * 7)// 14 
	  // Ko'rsatgich.
				k := 1
				q := &k
				fmt.Println(*q) // 1
	  // (/) Bo'lish 
				fmt.Println(7. / 2.) // 3.5 float64
				fmt.Println(7 / 2)   // 3   int
	  // (%) Qoldiq 
				fmt.Println(5 % 2)   // 1
	  // (<<) Chap'ga ozgartirish manosini bildiradi
				var x uint8 = 1 << 5
				var y uint8 = 1 << 3
				fmt.Printf("%08b\n", x) // 00100000 
				// birni besh chapga o'tkazdi
				fmt.Printf("%08b\n", y) // 00001000
				// birni uch marta chapga o'tkazdi
	// (>>) O'ng tomongo o'zgartirish xuddi yuqoridagidek ishlaydi

	// (&&) Va(and) 
				fmt.Println(true && true) // true
				fmt.Println(true && false) // false
	// (||) yoki 
				fmt.Println(true || true) // true
				fmt.Println(true || false) // true
	// (!) yoq 
				fmt.Println(!true) // false
	// (teng) 
				fmt.Println(2 == 3) // false
				fmt.Println(2 == 2) // true
	// (<=) oz yoki teng 
	// (>=) ko'p yoki teng 
	// (<) oz
	// (>) ko'p
	// (!=) teng emas
}				

// Floating point sonlar
// (1.23, 32.43, 0.00000343, 1233432)
// floattin point sonlar ikki turda keladi float32 fa float64
// ko'pincha dasturchilar float64'ni qullanishadi xullas float bu int'ga 
// o'xshash, Nuqta bilan ishlatsa bo'ladi ammo int'da nuqta mumkin emas.

