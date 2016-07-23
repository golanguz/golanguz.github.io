//Go'da birdan ko'p farqli turlar bor, Ko'pincha biz ularni ikkiga bo'lamiz
//ular integers va floating-point sonlar.


// Integerlar

/*
intergerlar xuddi ularnin matematik qarindoshlariga oxshaydi desimal komponentsiz.

Bizning kompyuter sistemamiz 10'ta farqli dijitdan iborat. Misol biz sanayotgan
vaqta bizarning fikrimizdagi dijitlar tugab qolsa biz, kattaroq sonlar ishlatishni
boshlaymiz, 2, 3, 4, 5, 6, 7, 8, 9.
Misol: 9'dan keyin 10, 99'dan keyin 100. Kompyuterlar xam xuddi shunday ammo
ular 10'tani o'rniga faqat 2 dijit ishlatishadi.
Va kompyuterlar uchun sanash quyidagidek: 0, 1, 10, 11, 100, 101, 110, 111 ...
biz ishlatadigan va kompyuterlar ishlatadigan sonlar o'rtasida farq, kompyuterlar
intejer turlarining aniqlangan uzunligi bor. Misol: 4-bit integer quyidagidek:
0000, 0001, 0010, 0011, 0100. Ammo bir oz vaqtdan keyin bunaqa sanashda
bo'sh joy qolmaydi, ko'pincha kompyuterlar boshidan boshlaydi lekin bu yo'l
ko'p noqulayliklarga olib kelishi mumkin
*/

//Go'ning intejer turlari bu:


//	int, uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64

/*
8, 16, 32, 64 bular yuqoridagi turlarning bit uzunligi.  uint bu imzolanmagan
integer, va int imzolangan integer, imzolanmagan integerlar faqat pozitive
sonlarni yozoladi Misol: 0 nol va noldan ziyod. qo'shimcha ikkita turlar bor
byte(esingizda bo'lsa oldindan aniqlangan so'zlar)  bu xuddi uint8 ga o'xshagan,
rune(bu xam oldindan aniqlangan so'zlar) bu xuddi int32 o'xshagan
Baytlar bullar judda xam kompyuterlarda muhim uint uzunlik o'lcham turlari
Misol:(1 byte = 8 bit, 1,024 byte = 1 kilobayt, 1,024 kilobayt = 1 megabayt ...)
shuning uchun Go'ning byte data turi ko'pincha boshqa turlarda aniqlanib 
ishlatiladi. Yana uchta moshina(sistema) qarashli intejer turlar bor
uint, int, va uintptr. Bular moshinaga qarashli chunki ularning uzunligi
moshina(sistema) arxitekturasiga qarashli.
Agar intejerlar bilan ishlayotgan bo'lsangiz  bor yoqi int turini ishlatishingiz
lozim.
*/


// Floating-Point sonlar


/*
floating-point tarjimasi, float bu havoda uchib turish xuddi kosmosdagidek,
poing bu ko'rsatgich, bir shaxsga barmoqingiz bilan ko'rsatsangiz, point,

floating-point sonlar bularning desimal komponenti bor, (xaqiqiy sonlar)
Misol: 1.432, 1234.6, 000000.54, 0.4353453, va 123459994. Bular floating-point
sonlar, Ularning kompyuterda korgazmasi juda xam qiyin va ularni ishlatish uchun
ularning detallarini bilish juda muhim emas,
xozircha quydagilarni eslab qolsanigz yetadi.

-- Floating-point bular nomalum sonlar. Va bir sonni ko'rastib bo'lmaydi
Misol 1.01 - 0.99 floating-point orqali tekshirsangiz javob
0.020000000000000018 to'g'ri javobja juda xam yaqin ammo to'g'ri emas.
-- Interjerlarga o'xshab, floating-point sonlarning aniqlangan uzunligi bor
(32 bit yoki 64 bit).
-- Bir necha boshqa qiymatlar bor Ular: "not a number(son emas)(NaN, 0/0 uchun)
pozitiv va negativ infiniti(oxirisiz) (+∞ va −∞)

Go'ning ikkita floating-point turlari bu float32 va float64(yakka va qo'shaloq
aniqlik deb ataladilar mos ravishda). Hamda qo'shimcha complex(murakkab) 
sonlari vakiligi ikkita turi bor (xayoliy sonlar qisimlari) complex64 va
complex128, Odatda floating-point sonlarni ishlatiyotganda float64 ishlatishimiz
lozim, Qulayroq.

Misol: ikki sonni qo'shadigan programma yarataylil
src quttingizni ichida sonlar degan quttini yaratig va uning ichida main.go 
faylini yaratig. Quydagi yozing.
*/

package main

import "fmt"

func main() {
	fmt.Println("1 + 2 =", 1 + 2 )
}
/*
kodni yurgizish uchun sonlar quttisini ichida terminal ochib go run main.go
buyruqini yurgizing
va terminalda "1 + 1 = 3" xabarini ko'rishingiz lozim.
ushbu programma xuddi salom gopherbek programmasiga o'xshagan 
bunda xam package main, import "fmt", va func main() { } bor
ammo salom gopherbek'ning o'rniga biz 1 + 1 = ifodasini yozib chiqaryapmiz
buning qisimlari iborati, sonlik ifoda 1(buning turi int) va + operatori
(qo'shmani bilintiradi) va boshqa sonlik ifoda 1, xuddi shu programmani 
floating-point bilan yozish mumkin quydagidek:
fmt.Println("1 + 2 =", 1.0 + 2.0)
bunda faqat bitta o'zgarish bor bu .(nuqta) Go to'planchisi buni floating-point 
deb biladi, ushbu programmani yurgizsangiz qaytimi xuddi birinchisidek bo'ladi

Go'ning arifmetik operatorlari (qisqacha)

+ qo'shma

- ayirma

* ko'paytirma

/ bo'lma

% qoldiq

operatorlarini xozircha eslab qolsangiz bo'ladi keyinroq bular xaqida ko'proq
o'rganamiz.
*/
