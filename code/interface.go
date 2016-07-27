// interfacelar bular nomlangan usullar imzolarining to'plami.
// interface turi o'zida, o'sha metodlar ishlatiyotgan qiymatlarni
// ushlab turoladi

package main

import "fmt"

// hammasi oddy
type User struct {
	FirstName, LastName string // ikki si ham string qaytaryapti
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	// yuqoridagi user structni qullanib u.FirstName, u.LastName terminalga
	// yoziyapti va bu tayyor ammo ism falimiya yozmaganimz hozircha, keyin yozamiz main funksiyasida
}

type Namer interface {
	Name() string // yuqoridagi name() metodini ushlab turibdi
}

func Greet(n Namer) string { // yuqoridagi namer inrefacini o'ziga oladi
	return fmt.Sprintf("Hurmatli %s", n.Name()) //va u orqali yuqoridagi
	// name() metodini chaqiradi, va undan oldin "Hurmatli" so'zini qo'yadi
}

func main() {
	u := &User{"Shavkat", "Izbasar"} // eng yuqoridagi user structiga ism familiya yozyapmiz
	fmt.Println(Greet(u))            // greet orqali ushbu u o'zgaruvchisini yozib chiqariyapmiz
	// terminal: Hurmatli Shavkat Izbasar
	// kodni organiyotganda o'qiyotgan kodingizni asstalik xolatda o'qish lozim va
	// ushbu kodni juda ham sekin o'qisangiz be'malol tushinib olasiz, faqat shoshmang!
}

/*
bu safar Greet funksiyasini aniqladik, va u, interface turi Namer bo'lgan
parametirini oladi, Namer bu biz aniqlagan yangi interface, va u,
bitta dona metod ushlab turibdi Name(). shunday qilib Greet() uni
Name() nomli farqli turlik parameterdek qabul qiladi.
*/
