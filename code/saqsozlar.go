/*
Go'da xuddi C++ ga oxshab comment yozasiz.
Misol ikki oldin slashlardan keyin // hamma so'zingiz komment deb o'qiladi.
ikki slashlar faqat yagona chiziq kommenti, Va ushbu men yoziyotgan so'zlar
slash / yuldizcha * va yuldizcha * slaschni / ichida bunisi esa ko'p chiziqlik
kommentlar uchun ishlatiladi.

Go'da 25 ta saqlangan so'zlari bor:
*/

go	var	  const	     import	 struct

if	type	  range	     package	 default

for	case	  else	     return	 interface

map	goto	  defer	     select	 continue

chan	break	  func	     switch      fallthrough

//----------------------------------------------------------------------------//

var type = select//O'zgaruvchini xato aniqlangan saqlangan so'z type va select

func defer() { } // Funksiya xato aniqlangan saqlangan so'z defer

type else go { } // Struct xato aniqlangan saqlangan so'z else va go
/*
yuqoridagidek saqlangan so'zlarni aniqlovchi qilib ishlatolmaymiz.
aniqlovchi uchun farqli so'zlar ishlatishimiz lozim misol:
*/

type Aniqlovchi struct { // To'g'ri
	Ismi string // To'g'ri
}

var ism string = "Gopherbek" // To'g'ri

func gopherBek() { // To'g'ri

}

/*
Aniqlanuvchi so'zlar juda xam zayif, Misol : GOPHERBEK, Gopherbek, GopherBek,
gopherBek, va gopherbek. Bular ummuman farqli 5'ta aniqlanuvchi so'zlar,
katta xarf bilan boshlanadigan so'zlar "Gopherbek" ular public(ochiq) bo'ladi
ular boshqa fayl ichida xam qullanilsa bo'ladi, ammo kichkina xar bilan
boshlangan so'zlar private(yopiq) va ular boshqa fayllarda qullanib bo'lmaydi.
katta xarf kichkina xarf qoidasi paket nomlariga ta'lluqli emas chunki
hamm package nomlari kichik xarflar bilan yoziladi.
*/

//----------------------------------------------------------------------------//

/*
Bunga qo'shimcha Go'da oldindan aniqlangan so'zlar bor:
*/

// Constants(o'zgarmaydiganlar)
	true	false	iota	nil

// Types(turlar)
	int	int8	int16	int64

	uint	uint8	uint16	uint32	uint64 uintptr

	float32	float64	complex128	complex64

	bool	byte	rune	string	error

// Functions(funksiyalar)

	make	len	cap	new	append	copy	close	delete

	complex	real	imag

	panic	recover

/*
yuqoridagi oldindan aniqlangan so'zlarni aniqlovchi qilib ishlatsangiz bo'ladi
ammo maslahatlanilmaydi chunki sizni noqulayliklargi olib kelishi mumkin.
*/
