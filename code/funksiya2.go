package main

import "fmt"

func qosh(args ...int) int {
	jami := 0
	for _, qiymat := range args {
		jami += qiymat
	}
	return jami // jaamini qaytar
}

func main() {
	fmt.Println(qosh(1, 2, 3, 4)) // 10
}

// ushbu kod parchasida qosh funksiyasi ellipsi'lik int'i bo'lgani
// uchun unga xoxlagancha son yozishingiz mumkin. qosh'ga yozilgan
// sonlarning ustidan range loop qilinib sonlar bir birlariga
// qo'shiladi
