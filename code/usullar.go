// method(usul) usullar hudii odiy funksiyalardek aniqlanadilar
// ammo bit farq bor u, funksiyaning nomidan oldin bitta qo'shimcha
// parameter yaratiladi, va shu parameter funksiyani o'zi qaratilgan
// turga yozadi(yopishtiradi).

package main

import "fmt"

type IsmSharif struct {
	Ism  string // bo'sh aniqlanmagan(qiymat yozilmagan)
	Yosh int    // bo'sh aniqlanmagan (qiymat yozilmagan)
}

// *IsmSharif structiga ko'rsatib uni funksiyani ichidagi i o'zgaruvchisini
// yaratib unga yozyapmi qulaylik uchun
// funksiyaning nomi yoz() string va int qaytaradi
func (i *IsmSharif) yoz() (string, int) {
	return i.Ism, i.Yosh // i o'zgaruvchisi orqali IsmSharifning
	// ichiga kirib i.Ism va i.Yosh qaytaryapmiz
}

func main() {
	x := IsmSharif{"Shavkat", 1989} // qiymat yozilyapti
	fmt.Println(x)                  // terminalga
	fmt.Println(x.yoz())            // method(usul) chaqirish boshqa yo'li
	// yoz methodini(usulini) x'ga yopishtirib chaqiryapmiz
	// x o'zgarucchisiga yuqoridagi IsmSharif'ni yoziyapmiz
	// va yuqoridagi usul o'sha biz x'ga yopishtirgan IsmSharif structidan
	// i.Ism, i.Yosh ni qaytariyapti.

}
