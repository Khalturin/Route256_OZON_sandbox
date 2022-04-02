package main

import (
	"fmt"
	"sort"
)

/*
Оригинальный пароль
	Каждый пароль должен состоять из четного количества символов,
	Первая половина пароля должна быть «анаграммно меньше» второй его половины.

	Нужен алгоритм, проверяющий, удовлетворяет ли придуманный пароль заданному условию.
	Вам дано n пар строк одинаковой длины (каждая строка – половина пароля).

	Строка s считается анаграммно меньше строки t, если существуют строка s' и строка t' такие что:
	s' получается из s перестановкой букв
	t' получается из t перестановкой букв
	s' лексикографически меньше, чем t'

	Для каждой пары строк si, ti определите правда ли, что si анаграммно меньше чем ti.

	_______________________________________________________________________________________________________

Формат входных данных
	В первой строке задано целое число n - число пар строк в тесте (1 < n < 100).
	В следующих 3n строках содержатся описания пар строк. Каждое описание состоит из трех последовательных строк.
	В первой строке каждого описания дана длина строк m_i (1 < m_i < 100), затем в следующих двух строках записаны
		строки  s_i, t_i длины m_i, состоящие из строчных латинских букв.

Формат выходных данных
	Для каждой пары строк из входных данных, выведите в i-й строке Yes, если строка s_i анаграммно меньше строки t_i.
	Иначе выведите No.

Примеры
Входные данные:
3
3
cba
aaz
2
bc
ab
1
a
a
Выходные данные:
Yes
No
No

*/

type myStruct struct {
	n string
	f string
	s string
}

type Ascending []byte

func (s Ascending) Len() int           { return len(s) }
func (s Ascending) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Ascending) Less(i, j int) bool { return s[i] < s[j] }

type Descending []byte

func (s Descending) Len() int           { return len(s) }
func (s Descending) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Descending) Less(i, j int) bool { return s[i] > s[j] }

func main() {
	var n int
	fmt.Scan(&n)
	inpStr := make([]myStruct, n, n)
	for i := 0; i < n; i++ {
		var mstr myStruct
		fmt.Scan(&mstr.n)
		fmt.Scan(&mstr.f)
		fmt.Scan(&mstr.s)
		inpStr[i] = mstr
	}

	for _, val := range inpStr {
		f := []byte(val.f)
		s := []byte(val.s)
		sort.Sort(Ascending(f))
		sort.Sort(Descending(s))
		val.f = string(f)
		val.s = string(s)

		//fmt.Println( val.f , val.s)
		//fmt.Println( val.f < val.s, isNormL, isNormR, isEqualsStr)

		if val.f < val.s {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
