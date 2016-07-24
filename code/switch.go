//switch bu huddi if'ga o'xshaydi, uni quyidagidek qullanamiz.

package main


import "fmt"


func main() {

	for i := 1; i <= 10; i++ {

		switch i {		
		case 1 : fmt.Println("bir")
		case 2 : fmt.Println("ikki")
		case 3 : fmt.Println("uch")
		case 4 : fmt.Println("to'rt")
		case 5 : fmt.Println("besh")
		case 6 : fmt.Println("olti")
		case 7 : fmt.Println("yetti")
		case 8 : fmt.Println("sakkiz")
		case 9 : fmt.Println("to'qqiz")
		case 10 : fmt.Println("o'n")
		}
	}
}

/*
if ning o'rniga switch ishlatyapmiz va else if'ning o'rniga case ishlatyapmiz
*/
