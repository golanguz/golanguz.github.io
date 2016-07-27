// error bu (xato) va u string bo'lib chiqadi, va bi o'zimizning error
// habarlarimizni yozishimiz mumkin

package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("xato habar")
	fmt.Println(err)
}
