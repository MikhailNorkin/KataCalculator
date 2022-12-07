package main

import (
	"fmt"
	"strconv"
	"strings"
)

func RomeToArabic(st string) int {
	var value int = 0
	for _, rune := range st {
		if string(rune) == "I" {
			value = value + 1
		} else if string(rune) == "V" {
			value = 5 - value
		} else if string(rune) == "X" {
			value = 10 - value
		}
	}
	return value
}

func ArabicToRome(value int) string {
	st := strconv.Itoa(value)
	res := ""
	var k int = 0
	for _, rune := range st {
		if ((k == 1) && (value > 10) || (value < 10)) && (value != 100) {
			if string(rune) == "1" {
				res = res + "I"
			} else if string(rune) == "2" {
				res = res + "II"
			} else if string(rune) == "3" {
				res = res + "III"
			} else if string(rune) == "4" {
				res = res + "IV"
			} else if string(rune) == "5" {
				res = res + "V"
			} else if string(rune) == "6" {
				res = res + "VI"
			} else if string(rune) == "7" {
				res = res + "VII"
			} else if string(rune) == "8" {
				res = res + "VIII"
			} else if string(rune) == "9" {
				res = res + "IX"
			}
		} else if value != 100 {
			if string(rune) == "1" {
				res = res + "X"
			} else if string(rune) == "2" {
				res = res + "XX"
			} else if string(rune) == "3" {
				res = res + "XXX"
			} else if string(rune) == "4" {
				res = res + "XL"
			} else if string(rune) == "5" {
				res = res + "L"
			} else if string(rune) == "6" {
				res = res + "LX"
			} else if string(rune) == "7" {
				res = res + "LXX"
			} else if string(rune) == "8" {
				res = res + "LXXX"
			} else if string(rune) == "9" {
				res = res + "XC"
			}
		} else {
			res = "С"
		}
		k = k + 1
	}
	return res
}

func main() {

	var st1, operator, st2, st3 string
	var count1, count2, countS3, res int
	mistake := 0 // здесь будем хранить ошибку

	fmt.Println("Введите выражение:")
	fmt.Scanf("%s %s %s %s", &st1, &operator, &st2, &st3)
	if st1 == "" || operator == "" || st2 == "" {
		mistake = 1
	}

	// проверим, есть ли еще что-то после выражения (третье число)
	if st3 != "" {
		mistake = 2
	}

	//проверим есть в выражении операторы "+-/*" и найдем что находится до и после оператора:
	if mistake == 0 {
		operators := "+-/*"
		countOperator := strings.IndexAny(operator, operators)
		if countOperator != 0 {
			mistake = 3
		}
	}

	//проверим являются ли первое и второе чило совпадающими по системам исчисления:
	if mistake == 0 {
		numbers1 := "1234567890"
		countS1 := strings.IndexAny(st1, numbers1)
		countS2 := strings.IndexAny(st2, numbers1)
		numbers2 := "IVX"
		countS3 = strings.IndexAny(st1, numbers2)
		countS4 := strings.IndexAny(st2, numbers2)
		if !(((countS1 == 0 && countS2 == 0) && (countS3 == -1 && countS4 == -1)) || ((countS1 == -1 && countS2 == -1) && (countS3 == 0 && countS4 == 0))) {
			mistake = 4
		}
		//проверим, если числа римские, то не должно быть вычитания:
		if countS3 == 0 && operator == "-" {
			mistake = 5
		} else if (st1 == "0") || (st2 == "0") {
			mistake = 6
		}
	}

	//если числа римские, переведем их в арабские:
	if mistake == 0 && countS3 == 0 {
		count1 = RomeToArabic(st1)
		count2 = RomeToArabic(st2)
	} else if mistake == 0 {
		countTemp1, err := strconv.Atoi(st1)
		if err != nil {
			panic(err)

		}
		count1 = countTemp1
		countTemp2, err := strconv.Atoi(st2)
		if err != nil {
			panic(err)
		}
		count2 = countTemp2
	}

	if count1 > 10 || count2 > 10 {
		fmt.Println("Числа не должны быть больше 10")
	} else if mistake == 1 {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
	} else if mistake == 2 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).\n")
	} else if mistake == 3 {
		fmt.Println("Вывод ошибки, не правильный опрератор.")
	} else if mistake == 4 {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
	} else if mistake == 5 {
		fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
	} else if mistake == 6 {
		fmt.Println("Вывод ошибки, так как числа не должны быть равны нулю.")
	} else {
		//посчитаем результат:
		if operator == "+" {
			res = count1 + count2
		} else if operator == "-" {
			res = count1 - count2
		} else if operator == "/" {
			res = count1 / count2
		} else if operator == "*" {
			res = count1 * count2
		}
		//если числа римские переведем результат в римские:
		if countS3 == 0 {
			fmt.Println(ArabicToRome(res))
		} else {
			fmt.Println(res)
		}
	}
}
