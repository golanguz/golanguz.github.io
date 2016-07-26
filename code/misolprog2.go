// aytaylik bir to'da sonlardan eng kichkinasini va eng kattasini
// terminalga yozadigan algoritm lozim, uni quyidagidek yozamiz:

package main

import "fmt"

func main() {
	x := []int{ // x o'zgaruvchisida int turi bo'lga sonlar aniqliyapmiz
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	kichik := x[0] // eng kichik sonni ushlaydigan kichik o'zgaruvchisini
	// aniqliyapmiz, va uning qiymati 0 deb o'yliyapmiz hozircha.
	buyuk := x[0] // va eng buyukning qiymati 0 deb o'yliyapmiz

	for _, qiymat := range x { //indeks'siz for yaratiyapmiz faqat qiymati bor
		// keyin x sonlar'ning uztidan loop qilyapmiz

		if qiymat > buyuk { // agar forning qiymati buyukdan katta bo'lsa
			buyuk = qiymat
			// qiymatni buyukga yoz.
		} else if qiymat < kichik { // agar qiymat kichikdan kichik bo'lsa
			kichik = qiymat
			// qiymatni buyukga yoz.
		}

	}
	fmt.Printf("eng kichik: %d, eng buyuk: %d\n", kichik, buyuk)

	//Eng kichik son: 9, Eng buyuk son: 97
	/*
		anglatma. hammasi for loopdan boshladani.
		for x'ning ustidan loop qiladi 16 martta, va tekshiradi.
		shunday qilib bor sonlarning ustidan aynaladi
		va hamma aynalib o'tilgan sonlarni kichik va buyuk
		o'zgaruvchilariga soladi("yozadi"). keyin
		if qiymat > buyukdan kattami yo'qmi?
		u bunday qilib so'riyapti aggar buyukdan
		katta bo'lsa qiymatni buyukga sol("yoz")
		va shunday qilib if bayonoti aynalib aynalib
		hamma sonlarni tekshirib va eng kattasini 97'ni
		buyuk'ga soladi.

		else if qiymat < kichikdan kichkinami yo'qmi?
		kichik o'zgaruvchisi ham huddi buyuk o'garuvchisidek
		faqat o'zining nomi aytib turgandek eng kichik sonni yozadi.
	*/
}
