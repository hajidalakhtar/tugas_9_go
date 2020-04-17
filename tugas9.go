package main

import (
	"fmt"
	"strconv"
)

func main() {
	defer pulihkan_saya()
	var input string
	fmt.Print("Masukan angka :")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		println(number, " ini adalah angka")
	} else {
		fmt.Println(input, "ini bukan angka")
		panic(err.Error())
		fmt.Println("haloo")
	}
}

func pulihkan_saya() {
	if r := recover(); r != nil {
		fmt.Println("Haloo saya kembali")
	} else {
		fmt.Println("lancar jayaa")
	}
}

// fungsi panic adalah dia akan berenti di baris kode yang terdapat panic dan kode di bawah nya tidak akan di jalankan
// recover itu untuk mengcek apakan terdapat panic di dalem sebuah kode jika terdapat panic maka recover akan bernilai tidak nil
