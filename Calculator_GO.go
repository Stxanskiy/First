package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func romanToInt(roman string) int {
	if roman == "I" {
		return 1
	} else if roman == "II" {
		return 2
	} else if roman == "III" {
		return 3
	} else if roman == "IV" {
		return 4
	} else if roman == "V" {
		return 5
	} else if roman == "VI" {
		return 6
	} else if roman == "VII" {
		return 7
	} else if roman == "VIII" {
		return 8
	} else if roman == "IX" {
		return 9
	} else if roman == "X" {
		return 10
	} else {
		return -1
	}
}

func intToRoman(num int) string {
	if num == 1 {
		return "I"
	} else if num == 2 {
		return "II"
	} else if num == 3 {
		return "III"
	} else if num == 4 {
		return "IV"
	} else if num == 5 {
		return "V"
	} else if num == 6 {
		return "VI"
	} else if num == 7 {
		return "VII"
	} else if num == 8 {
		return "VIII"
	} else if num == 9 {
		return "IX"
	} else if num == 10 {
		return "X"
	} else {
		return "Ошибка"
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
		return "Такакого логического оператора не существует"
	}

	parts := strings.Split(input, operator)
	if len(parts) != 2 {
		return "Неправильный формат примера"
	}
	a := strings.TrimSpace(parts[0])
	b := strings.TrimSpace(parts[1])

	isRoman := false
	aInt := romanToInt(a)
	bInt := romanToInt(b)

	if a == "0" || b == "0" {
		log.Fatal("0 не принимаеться программой")
	}

	if aInt != -1 && bInt != -1 {
		isRoman = true
	} else {
		aInt, _ = strconv.Atoi(a)
		bInt, _ = strconv.Atoi(b)
	}
	var result int
	if operator == "+" {
		result = aInt + bInt
	} else if operator == "-" {
		result = aInt - bInt
	} else if operator == "*" {
		result = aInt * bInt
	} else if operator == "/" {
		result = aInt / bInt
	}
	if isRoman {
		if result < 1 {
			return "Римские числа не могут быть меньше 1 "
		}
		return intToRoman(result)
	} else {
		return strconv.Itoa(result)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите пример?:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Результат:", calculate(strings.TrimSpace(input)))
	defer log.Println("Супер! Задание завершено.")
}
