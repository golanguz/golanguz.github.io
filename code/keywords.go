/*

"Saqlangan zo'zlar"

Go'ning funksiyalar(func), o'zgaruvchilar(var), o'zgarmaydiganlar(constant),
turlar(type), 
bayonot belgilari("if 0 < 1", agar 0 birdan < kam, bu yani 
statement "bayonot" degani, if, va switch saqlangan so'zlari 
ko'pincha bayonotda qullaniladi.), 
va paketlar, Oddiy qoidalar bilan qullaniladi. 
Masalan har bir ism harf bilan boshlanadi, 
Yoki past chizgi'si _ bilan, Va yani harflar bilan yoki sonlarbilan davom etadi, 
Va ularning bosmaligi muhum: 'Golang' va 'goLang' bular umuman ikki xil so'z.

*/

// Go'da 25'ta saqlangan so'zlar bor, Ular quyidagidek:


go	default		select		interface	for

goto	package		switch		fallthrough	const

case	defer		struct		return		case

var	range		import		defer		chan

type	if		func		else		map


// Qo'shimcha Go'da o'nlab oldindan aniqlangan so'zlar bor: 'int' va 'true', 
// va Go'ni ichiga qurilgan: 'constants'(o'zgarmaydiganlar), types(turlar),
// va funksiyalar

// O'zgarmaydiganlar:
		nil	true	false	iota

// Turlar:
		int	int8	int16	int32	int64
		uint	uint8	uint16	uint32	uint64	uintptr
		float32  float64  complex128  complex64
		bool 	byte	string	rune	error

// Funksiyalar
		make len cap new append copy close delete panic recover
		complex real imag

/*

Bu isimlar saqlanmagan va bularni aniqluvchilar qilib ishlatsangiz bo'ladi. Ammo 
ogoh bo'ling, aniqluvchi qilib ishlatganingizda sizni chalg'itish mumkin.
Agar mavjudot funksiyaning ichida aniqlangan bo'lsa demak bu mavjudot o'sha
funksiyaga ta'luqli va funksiyadan tashqarida qullanib bo'lmaydi, Agar funksiyadan
tashqarida aniqlangan bo'lsa demak bu o'z paketining ichidagi hamma fayllarga 
ko'rinadi.
Ismining birinchi harfi aniqlangan mavjudotni boshqa paketlarga va fayllarga 
kurinishini va ko'rinmasligini aniqlaydi, Agar ism Bosma harf bilan boshlansa 
demak u butun programmangizning ichida ko'rinadi va qullansa bo'ladi, 
Lekin paket isimlar faqat kichik harflar bilan yoziladi. Go'da aniqlanuvchi 
mavjudotlarning nomlanishining uzunligi uchun tusiqlik yo'q, Ammo Go dasturchilari 
ko'pincha qisqa aniqlanuvchi ismlarini ishlatadilar, Go kodini o'rganiyotganingizda 
ko'pincha aniqlanuvchi ism TheLoopIndex'ni o'rniga i'ni ko'rishingiz mumkin.
Va ko'pincha Go dasturchilar tuyaga o'xshash aniqlanuvchi ismlar qullanadilar.
Misol MyUzbekFunction, myUzbegimFunction, Ammo my_uzbek_function qullanmaydilar.
		
*/
		

			
