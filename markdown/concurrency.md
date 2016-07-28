##Concurrency(Paralellik)

>Buyuk programmalar ko'pincha kichik programmalar yoradmi bilan quriladi Misol, Web server browser qilgan so'roqlarni(talablarni) dastlab, Browserga HTML sahifalarini qaytaradi(yuboradi), Va browserdan kelayotgan so'roq kichiq programma orqali eplanadi(uddalanadi).

>Birdan ko'p so'roqlarni eplash uchun o'zining kichik kod parchalarini yurgizish buyuk programmalar uchun bu ajoyib, Taraqqiyotni birdan ko'p
vazifalarni bir vaqtning o'zida yurgizishi uchun. Va bu "Concurrency" deyiladi (Paralellik). Va Go dasturlash tili uni juda ham qo'llab quvvatlaydi. Concurrency(paralellik) qullanish uchun Go'da goroutine(go muntazami) va channel(kanal) ishlatiladi.


###Goroutines(Go muntazamlar)

>Goroutine bu funksiya, Va boshqa funksiyalar bilan paralellik darajada yurishiga qobiliyati bor. Groutine yaratish uchun biz saqlangan so'z go qullanamiz va undan keyin funksiya yoqilishi. Quyidagidek:


```go
package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)
}
```
>Yuqoridagi programmada ikkita go routine ishlatiliyapti. Birinchi goroutine u shubhasiz(ne yavniy) va u bizning main funksiyamiz.
Va ikkinchi goroutine u, Biz f(0) funksiyasini yurgazganimizda yaratiladi. Odatda funksiya chaqirilganda programma funksiyadagi hamma stasdiqlovchilarni yurigzadi va funksiya chaqirilayotgan chiziqdan(qatordan) pastgi chiziqqa(qatorga) qaytadi. Goroutine bilan esa darhol navbatdagi chiziqqa qaytadi, Va funksiya o'z yurgizmasini tugatishini kutmaydi. Shuning uchun "Scanln" funksiyasi kiritilgan(yozilgan) edi, Uning siz programma sonlarni yozishga ulgurmay, yurishni qo'xtatib chiqardi.

>Goroutinelar ular juda ham yengil va biz ulardan mingtalab yaratsak bo'ladi. 10'ta goroutine yurgizish uchun biz yuqoridagi programmani quyidagidek o'zgartirsak bo'ladi.

```go
func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
```
>Etibor qilgandirsiz, programmani yurgizganda huddi goroutinlar tartdibda yurgizilayotgandek, ammo ular bir faqtda yurishi lozim. Qani time.Sleep va rand.Intn funksiyasini qullanib bir oz ko'p kod qo'shaylik.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

/*
4 : 0
6 : 0
5 : 0
7 : 0
0 : 0
8 : 0
2 : 0
3 : 0
1 : 0
9 : 0
3 : 1
9 : 1
7 : 1
9 : 2
8 : 1
0 : 1
4 : 1
5 : 1
*/
```
>f() funksiyasi 0'dan 10'gacha sonlarni yuzib chiqaradi, Va yoziyotgan faqt har bir chaqiruvda kutib turadi 0 va 250ms(milli sekund)gacha.
