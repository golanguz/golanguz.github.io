package main

import "fmt"

type Person struct { // Person struct
	Name string // string tur
}

func (p *Person) Talk(ism string) {
	fmt.Println("Salom mening ismim", p.Name, ism)
}

type Android struct { // Android struct
	Person        // yuqoridagi Person structga yopishtirilgan(ko'rsatilgan)
	Model  string // string tur
}

func (a *Android) Soyla(model string) {
	fmt.Println("Mening modelim", a.Model, model)
}

func main() {
	a := Android{}           // Android structini a o'zgaruvchisiga yozyapmiz
	a.Person.Talk("Shavkat") // a orqali Androidni ichiga kirib va talk()
	// funksiyasi orqali talk()funksiyasi ichidagi fmt.Println() ichidagi
	// habarni yozib chiqaryapmiz, talk(ism, string) funksiyasinin parameteri
	// bo'lgani uchun biz a orqali personi chaqiriyotganda talk()'ga "Shavkat"
	// yozsak go'ladi
	a.Talk("Shavkat") // yuqoridagini chaqirishning boshqa yo'li
	a.Soyla("045")
}
