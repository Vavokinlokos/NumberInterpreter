package main

import (
	"fmt"
	"strings"
)

var inputAlphabeth = [31]string{"один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять", "десять",
	"одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать",
	"девятнадцать", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят",
	"девяносто", "сто", "двести", "ста", "сот"}

func main() {
	input := "пять сот двадцать один"
	input = "три ста пять"
	input = "девятнадцать"
	input = "сто тринадцать"
	input = "два два"
	input = "три сот пять"
	input = "три"
	fmt.Println("Input is ", input)

	matrix := [5][32]int{
		{5, 5, 1, 1, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 0, 0, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 5},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 0, 0, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 5},
	}

	words := make(map[string]int)

	words["один"] = 1
	words["два"] = 2
	words["три"] = 3
	words["четыре"] = 4
	words["пять"] = 5
	words["шесть"] = 6
	words["семь"] = 7
	words["восемь"] = 8
	words["девять"] = 9
	words["десять"] = 10
	words["одиннадцать"] = 11
	words["двенадцать"] = 12
	words["тринадцать"] = 13
	words["четырнадцать"] = 14
	words["пятнадцать"] = 15
	words["шестнадцать"] = 16
	words["семнадцать"] = 17
	words["восемнадцать"] = 18
	words["девятнадцать"] = 19
	words["двадцать"] = 20
	words["тридцать"] = 30
	words["сорок"] = 40
	words["пятьдесят"] = 50
	words["шестьдесят"] = 60
	words["семьдесят"] = 70
	words["восемьдесят"] = 80
	words["девяносто"] = 90
	words["сто"] = 100
	words["двести"] = 200

	splittedInput := strings.Split(input, " ")

	i, j := 0, 0
	result := 0
	for _, v := range splittedInput {
		if i == 5 {
			fmt.Println("Вы допустили ошибку. Введите строку снова.")
			return
		}
		if v == "ста" || v == "сот" {
			result *= 100
		} else {
			result += words[v]
		}

		j = GetJ(v)
		if j == -1 {
			fmt.Println("Вы допустили ошибку. Введите строку снова.")
			return
		}
		i = matrix[i][j]
		if i == 0 {
			fmt.Println("Вы допустили ошибку. Введите строку снова.")
			return
		}
	}
	fmt.Println("Result is ", result)
	return
}

func GetJ(word string) int {
	for i, v := range inputAlphabeth {
		if word == v {
			return i
		}
	}
	return -1
}
