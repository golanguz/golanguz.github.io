// Boolean bu oddiy bir qiymat. (George Boole)'ga atab qo'yilgan.
// mahsus 1 bit'lik intejer tur, true(rost) va false(yolg'on) bilintiradi
// ko'pincha uchta ishlamchi bilan ishlatiladi ular: or(yoki) and(va) not(yo'q)

package main


import "fmt"


func main() {

		      // && and(va) manosini bildiradi
	fmt.Println(true && true)  // qaytim true
	fmt.Println(false && true) // qaytim false

		      // || or(yoki) manosini bildiradi
	fmt.Println(true || true)  // qaytim true
	fmt.Println(true || false) // qaytim true
		
		 // ! not(yoq) manosini bildiradi
	fmt.Println(!true)         // qaytim false
//--------------------------------------------------------------------------//	
	a := "1" // o'zgaruvchi aniqlaymiz uning qiymati "1"
	b := "2" // "2"
	fmt.Println(a == b)// false chunki 1 va 2 bir biriga teng emas
	fmt.Println(a <= b)// true chunki 1, 2'dan kam(kichkina) 
	fmt.Println(a >= b)// false chunki 1, 2'dan katta yoki teng emas
	fmt.Println(a != b)// true chunki 1, 2'ga teng emas
	fmt.Println(a > b) // false chunki 1, ikkidan katta emas
	fmt.Println(a > b && b >= a) // false chunki:
		//  fasle &&  true   // flase
	fmt.Println(a > b || b >= a)
}		//  false ||  true   // truecl

// yuqoridagilarni yurgizsangiz quyidagi xabarni ko'rasiz:
/*
	true
	false
	true
	true
	false
	false
	true
	false
	true
	false
	false

tushunmiyotgan bo'lsangiz o'qishni davom ettirin chunki kod yozishni o'rganish
oson emas kod yozishni tez o'rganish bu ko'p o'qish va o'qigan kodlaringizni
eng kam bir martta yozish. Farqli dasturchilardan kod yozishni eng oson va tez o'rganadigan yo'li so'rasangiz bu PRACTICE(AMALIYOT) kodni yozing ko'proq, 
Va qancha ko'p xato qilsangiz shuncha yaxshi. xato qilishdan qo'ramang.
*/
