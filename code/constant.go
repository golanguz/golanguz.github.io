/*
constants(o'zgarmaydiganlar) ular huddi var dek ammo ularning qiymati o'zgarmaydi
*/

package main


import "fmt"


func main() {

	const a string = "Salom Gopherbek"
	fmt.Println(a)
	a = "Salom Shavkat" // xato beradi "cannot assign to a" (a'ga yangi 
// qiymat yozolmaymiz).
	fmt.Println(a)
 
}
