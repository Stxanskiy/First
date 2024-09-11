package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func romanToInt(roman string) (int, bool) {
	switch roman {
	case "I":
		return 1, true
	case "II":
		return 2, true
	case "III":
		return 3, true
	case "IV":
		return 4, true
	case "V":
		return 5, true
	case "VI":
		return 6, true
	case "VII":
		return 7, true
	case "VIII":
		return 8, true
	case "IX":
		return 9, true
	case "X":
		return 10, true
	default:
		return -1, false
	}
}

func validateNumberRange(num int) {
	if num < 1 || num > 10 {
		log.Panic("Программа принимает на вход числа от 1 до 10")
	}
}

func calculate(input string) string {
	input = strings.TrimSpace(input)
	var operator string
	if strings.Contains(input, "+") {
		operator = "+"
	} else if strings.Contains(input, "-") {
		operator = "-"
	} else if strings.Contains(input, "*") {
		operator = "*"
	} else if strings.Contains(input, "/") {
		operator = "/"
	} else {
		log.Panic("Неправильный оператор")
	}

	parts := strings.Split(input, operator)
	if len(parts) != 2 {
		log.Panic("Такой формат не оддерживаетсяя")
	}
	a := strings.TrimSpace(parts[0])
	b := strings.TrimSpace(parts[1])

	aInt, isRomanA := romanToInt(a)
	bInt, isRomanB := romanToInt(b)

	var aNum, bNum int

	if isRomanA && isRomanB {
		aNum = aInt
		bNum = bInt
		validateNumberRange(aNum)
		validateNumberRange(bNum)
	} else if !isRomanA && !isRomanB {
		var errA, errB error
		aNum, errA = strconv.Atoi(a)
		bNum, errB = strconv.Atoi(b)

		if errA != nil || errB != nil {
			log.Panic("Числа не распознаны")
		}
		validateNumberRange(aNum)
		validateNumberRange(bNum)
	} else {
		log.Panic("сложение арабского и римского числа не поддерживается")
	}

	var result int
	switch operator {
	case "+":
		result = aNum + bNum
	case "-":
		result = aNum - bNum
	case "*":
		result = aNum * bNum
	case "/":
		if bNum == 0 {
			log.Panic("Не делиться на ноль ")
		}
		result = aNum / bNum
	default:
		log.Panic("В школе не учился? такой оператор не поддерживаеться")
	}

	if isRomanA && (result < 1) {
		log.Panic("Римские числа -1 не принимаются")
	}

	return strconv.Itoa(result)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введи пример:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Результат :", calculate(strings.TrimSpace(input)))
	defer log.Println("Красавчик!")
}
