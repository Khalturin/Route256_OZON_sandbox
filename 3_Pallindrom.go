package main

import (
	"fmt"
	"math"
)

/*
Счастливый заказ
	Условия
	Лена узнала, что на ее любимом маркетплейсе действует интересная акция.
	Согласно её условиям, организаторы раздают подарки всем клиентам, впервые сделавшим заказ, номер которого
	окажется палиндромом. Лена заметила, что номера заказов формируются в порядке возрастания, и решила перехитрить
	организаторов: в нужный момент очень быстро оформить подряд несколько мелких заказов, чтобы гарантированно «выбить»
	счастливый номер.

	Помогите Лене определить, насколько близко номер текущего заказа к ближайшему за ним «счастливому» номеру.

	Примечание: Число является палиндромом, если читается слева направо так же, как справа налево.
	Пример палиндрома: 123321

    _______________________________________________________________________________________________________________

Формат входных данных
	Дано целое положительное число n, не являющееся палиндромом (гарантируется, что n - не палиндром), n ≤ 10^13.

Формат выходных данных
	Выведите разницу между ближайшим палиндромом p (p > n) и входным числом n.
*/

// Увеличение порядка числа
func exponent(num, amount int) int {
	return num * int(math.Pow10(amount))
}

// Преобразовать число в массив
func numToMass(number int) (mass []int) {
	// Заполнить массив
	for n := number; n > 0; n /= 10 {
		mass = append(mass, n%10)
	}
	// Перевернуть массив
	length := len(mass)
	for i := 0; i < length/2; i++ {
		mass[i], mass[length-i-1] = mass[length-i-1], mass[i]
	}
	return
}

// Преобразовать массив в число
func massToNum(mass []int) (resVal int) {
	for i, val := range mass {
		resVal += exponent(val, i)
	}
	return
}

// Получить левую часть массива, если правая половина полиндрома больше левой, последнюю увеличить на 1
func getHalfPallindrom(mass []int) (resMass []int) {
	var leftN, rightN, realLeft int

	j := len(mass) - 1
	half := len(mass) / 2 // Половина массива

	// Если четное число, уменьшаем кол-во итераций на 1
	rIdx := 0
	if len(mass)%2 == 0 {
		half -= 1
		rIdx = 1
	}

	// Переводим левую и правую часть в число, для проверки
	for i := 0; i <= j; i++ {
		idxInsert := half - i

		rightN += exponent(mass[half + i + rIdx], idxInsert)
		leftN += exponent(mass[idxInsert], idxInsert)

		realLeft += exponent(mass[i], idxInsert)

		j--
	}
	// Если правая сторона больше левой
	if rightN > leftN {
		realLeft++
	}
	//Переводим число в массив
	resMass = numToMass(realLeft)

	return
}

func nearestPall(number int) (resVal int) {
	mass := make([]int, 0, 0)
	mass = numToMass(number)
	isOdd := len(mass)%2 != 0

	// получаем половину паллиндрова
	mass = getHalfPallindrom(mass)

	// Добавить элементы в результирующий
	lastIdx := len(mass) - 1
	for i := lastIdx; i >= 0; i-- {
		// Если кол-во эл-ов массива нечетно, пропустить добавление последнего элемента - он будет средним
		if isOdd && i == lastIdx {
			continue
		}
		mass = append(mass, mass[i])
	}
	resVal = massToNum(mass)

	return resVal
}

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(nearestPall(number) - number)

	// Тестовые данные:
	//fmt.Println("654: ", nearestPall(654))
	//fmt.Println("5548: ", nearestPall(5548))
	//fmt.Println("5448: ", nearestPall(5448))
	//fmt.Println("5443: ", nearestPall(5443))
	//fmt.Println("10: ", nearestPall(10))
	//fmt.Println("12: ", nearestPall(12))
	//fmt.Println("21: ", nearestPall(21))
	//fmt.Println("19: ", nearestPall(19))
	//fmt.Println("91: ", nearestPall(91))
	//fmt.Println("100: ", nearestPall(100))
	//fmt.Println("109: ", nearestPall(109))
	//fmt.Println("900: ", nearestPall(900))
	//fmt.Println("123: ", nearestPall(123))
	//fmt.Println("321: ", nearestPall(321))
	//fmt.Println("789: ", nearestPall(789))
	//fmt.Println("987: ", nearestPall(987))
	//fmt.Println("798: ", nearestPall(798))
	//fmt.Println("897: ", nearestPall(897))
	//fmt.Println("1000: ", nearestPall(1000))
	//fmt.Println("1009: ", nearestPall(1009))
	//fmt.Println("1234: ", nearestPall(1234))
	//fmt.Println("4321: ", nearestPall(4321))
	//fmt.Println("6789: ", nearestPall(6789))
	//fmt.Println("9876: ", nearestPall(9876))
	//fmt.Println("7998: ", nearestPall(7998))
	//fmt.Println("8997: ", nearestPall(8997))
	//fmt.Println("12345: ", nearestPall(12345))
	//fmt.Println("54321: ", nearestPall(54321))
	//fmt.Println("56789: ", nearestPall(56789))
	//fmt.Println("98765: ", nearestPall(98765))
	//fmt.Println("79998: ", nearestPall(79998))
	//fmt.Println("89997: ", nearestPall(89997))
}
