package main

import (
	"fmt"
	"lession1"
	"lession2"
	"lession3"
	"lession4"
	"lession5"
	"lession6"
)

func main() {
	if false {
		lession1.Lession1()
	}
	if false {
		lession2.Lession2()
	}
	if false {
		lession3.Lession3()
		lession3.Lession3_1()
		lession3.Lession3_2()
		lession3.Lession3_3()
	}
	if false {
		lession4.Lession4()
		lession4.Lession4_1()
		lession4.Lession4_2()
	}
	if false {
		lession5.Lession5()
		lession5.Lession5_1()
		lession5.Lession5_2()
		lession5.Lession5_3()
	}
	if true {
		lession6.Lession6()
		fmt.Println(lession6.Pi) //3.14
		lession6.Lession6_1()
	}

}
