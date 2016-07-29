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

----

##Channels(Kanallar)
>Kanallar goroutinelar yurguzilishi uchun va sinxronizatsiya qilishi uchun yordam beradi. Quydagi misol channels(kanallar) qullanadi.


```go
package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}

```

>Ushbu programma pong so'zini to'xtovsiz yozib chiqaradi (to'xtatish uchun enter bosing). Kanal turi chan saqlangan so'zi orqali yaratiladi va undan keyin unga o'tkazilayotgan narsalarning turlari yoziladi(misol yuqoridagi programmada biz string o'tkaziyapmiz). chap nayza operatori(<-) kanalga yuborish va qabul qilish manosini bildiradi. c <- "ping" manosi "ping" stringini c'ga yubor. msg := <- c manosi kelayotgan habarni qabul et va o'sha habarni msg'ga yukla. va fmt parchasini quydagidek yozsak bo'lardi, fmt.Println(<-), va bu holda biz undan oldingi chiziqni o'chirsak bo'lardi.

----

>Kanallarni bu yo'lda qullanish ikki goroutinelarni sinxronizatsiya qiladi. Pinger habarni kanalga yo'llashga harakat qilganda u, Printer habarni qabul qilishga tayyor bo'lguncha kutib turadi(bu block deyiladi). Kodimizda yana bitta yuborvuchi aniqlaylik. Quydagi funksiyani aniqlang.

```go
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
```

>Va main funksiyasini quydagidek o'zgartiring

```go
func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
```
>Ushbu programma endi habarlarni  navbat ma navbat yozadi "ping" "pong"

##Kanal yo'nalishi

>Biz kanal turida yonalish aniqlasak bo'ladi, shunday qilib biz uni yuborishdan yoki qabul etishdan muhum qilamiz. Misol pinger funksiyasining imzosi quyidagidek o'zgartirilishi mumkin.

```go
func pinger(c chan <- string)
```
>Endi pinger c ga yuborishga haqqi bor. Agar c'dan qabul qilishga urinsangiz to'plamchi xatosini ko'rasiz. Shunga o'xshab  printer funksiyasini ham o'zgartirsak bo'ladi.

```go
func printer(c chan <- string)
```
>Muhum qilinmagan kanallar **bidirectional**(ko'p yollanmali) deb ataladilar. **bidirectional** kanal qabul faqat qiladigan yoki faqat yuboradigan kanallarga o'tkazilishi mumkin. Ammo chappasi no to'g'ri.

##Select(Tanlash)

>Go'ning ajoyib bir tasdiqlovchisi bor va bu select u xuddi switch'ga o'xshaydi ammo faqat kanallar uchun ishlaydi.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}

```

>Ushbu programma "from 1" habarini har 2 soniyada va "from 2" habarini har 3 soniyada yozib chiqaradi. Select birinchi tayyor bo'lgan kanalni olib qabul etadi (yoki yuboradi). Agar birdan ko'p kanallar tayyor bo'lsa, u xolda qaysi biridan qabul etishni bilmasdan turib tanlaydi, Agar hech bir kanal tayyor bo'lmasa, Tasdiqlovchi yopiladi(bloklanadi) va kutadi kanallardan birini tayyor bo'lishini.

----

>Select tasdiqlovchisi ko'pinshi taymaut'ni bajarish uchun ishlatiladi.

```go
select {
case msg1 := <-c1:
	fmt.Println("Message 1", msg1)
case msg2 := <-c2:
	fmt.Println("Message 2", msg2)
case <-time.After(time.Second):
	fmt.Println("timeout")
}
```

>time.After kanal yaratadi va berilgan vaqtdan keyin, hozirgi vaqtni unga yuboradi(hozirgacha biz vaqt yuborishga qiziqmadig, shunday qilib biz o'zgaruvchiga vaqt yozmadik). Va select uchu default(oddiy) so'zini ishlatsak boladi.

```go
select {
case msg1 := <-c1:
	fmt.Println("Message 1", msg1)
case msg2 := <-c2:
	fmt.Println("Message 2", msg2)
case <-time.After(time.Second):
	fmt.Println("timeout")
default:
	fmt.Println("Tayyor emas")
```
>Default, agar kanallardan hech biri tayyor bo'lmasa o'sha zaxod yuroziladi.

##Buffered Channels(buferlangan kanallar)

>Kanall yaratiyotganda make funksiyasiga ikkinchi parameterni ham yozsangiz bo'ladi.

```go
	c := make(chan int, 1)
```

>Va bu buferlangan kanall yaratadi va uning sig'imi 1. Odatda kanallar sinxronizatsiya'ga qobiliyatli; kanal'ning ikki tomoni ham kutadi toki ularnin bir tomoni tayyor bo'lguncha. Buferlangan kanal esa sinxronizatsiya'siz; habarni yuborish va qabul etish kutib turmaydi toki kanallar alla qachon to'la, Ma'bodo kanal to'la bo'lsa, yuboruvchi u yerda hattoki bitta dona int uchun joy paydo bo'lguncha kutib turadi.

##Misol:

>Ushbu misol programma goroutinelar va kanallar ishlatadi.
u bir necha web saytlardan ularning sahifalarini bir faqtda sog'ib oladi, Va uning uchun net/http paketini qullanadi. Va eng katta sahifani URL'ini yozib chiqaradi(eng ko'p baytlar qaytimi deb aniqlangan).

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}
	results := make(chan HomePageSize)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}
	var biggest HomePageSize
	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}
	fmt.Println("The biggest home page:", biggest.URL)
}
```

>Sahifalarni saqlaydigan bir tur aniqladik

```go
urls := []string{
	"http://www.apple.com",
	"http://www.amazon.com",
	"http://www.google.com",
	"http://www.microsoft.com",
}
```

>Keyin esa URL'larni sog'iydigan kanal yaratdik va u ikkita yangi goroutine yurgizadi har bitta url uchun(urllarni bir vaqtda sog'ish uchun). Har bir URL uchun biz HTTP get so'roqini yuboriyapmiz.

```go
res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
defer res.Body.Close()
```

>Endi esa javob vujudining o'lchamini saqlaymiz.

```go
bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
	panic(err)
	}

results <- HomePageSize{
	URL: url,
	Size: len(bs),
}
```
>Etibor bering bu nomlanmagan funksiya aniqlangan vaqt yugiziladi. Goroutinlarda bunaqa narsalar ko'pincha shilatiladi, etibor bering yagona parametr(url) oliyotgan fuknsiya ham aniqladik. Sabab ushbu funksiya url'ni to'g'ridan to'g'ri yollamaydi, ammo qilsa bo'ladi, Closure quydalari shunaqaki agar biz shunday qilganimizda, hamma to'rta goroutinelarimiz url uchun bir qiymatni ko'rishi mumkin edi. Chunki url for loop orqali o'zgartiriliyapti.

----

>Va oxirida biz kanaldan qaytimlarni tortish uchun 4 martta loop qiliyapmiz, qaytimlarni hozirgi buyuk habarga tenglab ko'ring, va joylarini almashtiring.

```go
var biggest HomePageSize
for range urls {
	result := <-results
	if result.Size > biggest.Size {
		biggest = result
		}
}
```

#Dars tugadi tashrif buyurganingz uchun rahmat.
