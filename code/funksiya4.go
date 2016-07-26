// funksiyalar ichida boshqa funksiyalar yaratish ham mumkin

package main

import "fmt"

/*
func main(){
  qosh := func(x, y int) int {
    return x + y
  }
  fmt.Println(qosh(1,1))
}
qosh o'zgaruvchisi main funksiyasiga tegishli o'z o'zgaruvchisi.
bunaqa yaratiladigan o'zgaruvchilar local(mestniy) deyiladilar
va ushbu o'zgaruvchi orqali yaratilgan funksiya, ushbu main
funksiyasi ichidagi o'zgaruvchilarni qullansa bo'ladi
*/

/*
func main(){
    x := 0
    buyukla := func() int {
    x ++
    return x
  }
  fmt.Println(buyukla())
  fmt.Println(buyukla())
}
har bir martta buyukla funksiyasi yurgizilganda
x o'zgaruvchisini buyuklantiradi.
va ko'rib turganingizdek buyukla funksiyasi o'zinin ichida bo'lmagan
ammo main funksiya ichida bo'lga x o'zgaruvchisini qullaniyapti.
va ushbu yo'lda funksiya aniqlash closure deb ataladi.
closure funksiyasini yaratish bu bir funksiya boshqa funksiya qaytaradi.
*/
// turi uint bo'lgan func() qaytariyapmiz
func juftSonlar() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
func main() {
	keyingiJuft := juftSonlar()
	fmt.Println(keyingiJuft()) // 0
	fmt.Println(keyingiJuft()) // 2
	fmt.Println(keyingiJuft()) // 4
}

// juftSonlar funksiyasi juft sonlarni yozadi. va qayta chaqirilganda
// oldingi songa 2 sonini qo'shadi.
