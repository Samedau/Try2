package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a string
	var b string
	var c string
	fmt.Scanf("%q %s %s", &a, &b, &c)

	if len([]rune(a)) > 10 {
		panic(fmt.Sprintf("Строка не должна иметь больше десяти символов!"))
	}
	if len([]rune(c)) > 10 {
		panic(fmt.Sprintf("Строка не должна иметь больше десяти символов!"))
	}
	//+
	if strings.Contains(b, "+") {

		tem := strings.ReplaceAll(a+c, "\"", "")

		fmt.Print("\"" + tem + "\"")

		///-
	} else if strings.Contains(b, "-") {

		j, k := strconv.Atoi(c)
		if k == nil {
			panic(fmt.Sprintf("Вычитать можно только строку из строки!"))
			os.Exit(1)
			fmt.Println(j)
		}

		tem := strings.Replace(a, strings.Replace(c, "\"", "", 2), "", 1)
		fmt.Print("\"" + tem + "\"")

		///*
	} else if strings.Contains(b, "*") {
		j, k := strconv.Atoi(c)
		if k == nil {

			d, g := strconv.Atoi(c)
			if g == nil {
				if d > 10 {
					panic(fmt.Sprintf("Использовать числа только от 1 до 10!"))
				}
				if d < 0 {
					panic(fmt.Sprintf("Использовать числа только от 1 до 10!"))
				}
			}

			var t string
			y, e := strconv.Atoi(c)
			if e == nil {
				for i := 0; i < y; i++ {
					t = t + a
				}
				tem := strings.ReplaceAll(t, "\"\"", "")

				str := []rune(tem)
				if len(str) > 40 {
					tem = string(str[0:41])
					fmt.Print(tem)
					fmt.Println("...\"")
				} else {
					fmt.Print("\"" + tem + "\"")
				}

			}
		} else {
			panic(fmt.Sprintf("Нельзя умножать не на число!"))
			fmt.Sprint(j)
		}
		//////
	} else if strings.Contains(b, "/") {

		j, k := strconv.Atoi(c)
		if k == nil {
			d, g := strconv.Atoi(c)
			if g == nil {
				if d > 10 {
					panic(fmt.Sprintf("Использовать числа только от 1 до 10!"))
				}
				if d < 0 {
					panic(fmt.Sprintf("Использовать числа только от 1 до 10!"))
				}
			}
			y, e := strconv.Atoi(c)
			a = strings.ReplaceAll(a, "\"", "")
			str := []rune(a)
			if e == nil {
				f := string(str[0 : len(str)/y])

				fmt.Print("\"" + f + "\"")

			}
		} else {
			panic(fmt.Sprintf("Нельзя делить не на число!"))
			fmt.Sprint(j)
		}

	} else {
		panic(fmt.Sprintf("Неверное выражение!"))
	}
}
