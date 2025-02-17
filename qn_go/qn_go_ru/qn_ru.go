package main

import (
	"fmt"
	"math"
)

func main() {
	// made by sm300max
	fmt.Println("made by sm300max")

	var a, b, c float64
	var is string

	fmt.Println("Введите первый коэффициент:")
	fmt.Scan(&a)

	fmt.Println("Введите второй коэффициент:")
	fmt.Scan(&b)

	fmt.Println("Введите третий коэффициент:")
	fmt.Scan(&c)

	fmt.Println("Введите знак [</<=/=/=>/>]:")
	fmt.Scan(&is)

	if a == 0 {
		fmt.Println("Квадратное выражение не может содержать а раное нулю")
		return
	}

	if is != "<" && is != "<=" && is != "=" && is != "=>" && is != ">" {
		fmt.Println("Допущена ошибка в выборе знака, попробуйте ещё раз")
		return
	}

	var q string

	fmt.Println(fmt.Sprint(a) + "x²" + "+" + fmt.Sprint(b) + "x" + "+" + fmt.Sprint(c) + is + "0")
	fmt.Println("Выражение верное? [Y/n]")
	fmt.Scanln(&q)

	if q != "y" && q != ""{
		fmt.Println("Попробуйте ещё раз")
		return
	}
	var ud int8
	fmt.Println("Введу функцию:")
	fmt.Println("y=" + fmt.Sprint(a) + "x²" + "+" + fmt.Sprint(b) + "x" + "+" + fmt.Sprint(c))
	fmt.Println("Квадратная функция")
	fmt.Println("График — парабола")
	if a > 0 {
		ud = 1
		fmt.Println("Ветви вверх")
	} else {
		ud = 2
		fmt.Println("Ветви вниз")
	}

	fmt.Println("Нули функции: " + fmt.Sprint(a) + "x²" + "+" + fmt.Sprint(b) + "x" + "+" + fmt.Sprint(c) + "=0")

	var x1, x2 float64
	var x float64
	var oot int8

	D := b*b - 4*a*c
	if D < 0 {
		fmt.Println("D = b² - 4ac")
		fmt.Println("D = " + fmt.Sprint(b) + "² - 4*" + fmt.Sprint(a) + "*" + fmt.Sprint(c))
		fmt.Println("D = " + fmt.Sprint(D))
		fmt.Println("D < 0")
		fmt.Println("Выражение не имеет корней")
		fmt.Println("Решено")
	} else if D == 0 {
		oot = 1
		x = -b / (2 * a)
		fmt.Println("D = b² - 4ac")
		fmt.Println("D = " + fmt.Sprint(b) + "² - 4*" + fmt.Sprint(a) + "*" + fmt.Sprint(c))
		fmt.Println("D = " + fmt.Sprint(D))
		fmt.Println("D = 0")

		fmt.Println("Выражение имеет один корень")
		fmt.Println("x = -b / 2a")
		fmt.Println("x = -" + fmt.Sprint(b) + "/2" + fmt.Sprint(a))
		fmt.Println("x = " + fmt.Sprint(x))
		fmt.Println("Решено")
	} else if D > 0 {
		oot = 2
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println("D = b² - 4ac")
		fmt.Println("D = " + fmt.Sprint(b) + "² - 4*" + fmt.Sprint(a) + "*" + fmt.Sprint(c))
		fmt.Println("D = " + fmt.Sprint(D))
		fmt.Println("D > 0")
		fmt.Println("Выражение имеет два корня")
		fmt.Println("x = (-b ± √D) / (2a)")

		if a < 0 && b < 0 {
			fmt.Println("(x1=-(" + fmt.Sprint(b) + ")+√" + fmt.Sprint(D) + ") / (2*(" + fmt.Sprint(a) + "))")
			fmt.Println("(x1=-(" + fmt.Sprint(b) + ")-√" + fmt.Sprint(D) + ") / (2*(" + fmt.Sprint(a) + "))")
		} else if b < 0 {
			fmt.Println("(x1=-(" + fmt.Sprint(b) + ")+√" + fmt.Sprint(D) + ") / (2*" + fmt.Sprint(a) + ")")
			fmt.Println("(x1=-(" + fmt.Sprint(b) + ")-√" + fmt.Sprint(D) + ") / (2*" + fmt.Sprint(a) + ")")
		} else if a < 0 {
			fmt.Println("(x1=-" + fmt.Sprint(b) + "+√" + fmt.Sprint(D) + ") / (2*(" + fmt.Sprint(a) + "))")
			fmt.Println("(x1=-" + fmt.Sprint(b) + "-√" + fmt.Sprint(D) + ") / (2*(" + fmt.Sprint(a) + "))")
		} else {
			fmt.Println("(x1=-" + fmt.Sprint(b) + "+√" + fmt.Sprint(D) + ") / (2*" + fmt.Sprint(a) + ")")
			fmt.Println("(x1=-" + fmt.Sprint(b) + "-√" + fmt.Sprint(D) + ") / (2*" + fmt.Sprint(a) + ")")
		}
		fmt.Println("x1 = " + fmt.Sprint(x1))
		fmt.Println("x2 = " + fmt.Sprint(x2))
	}

	var start, end float64
	switch {
	case x1 > x2:
		start = x2
		end = x1
	case x1 < x2:
		start = x1
		end = x2
	}

	if oot == 2 {

		if ud == 1 {
			if is == ">" || is == "=>" {
				if is == ">" {
					fmt.Println("Ответ: (-∞;" + fmt.Sprint(start) + ")∪(" + fmt.Sprint(end) + ";+∞)")
				} else if is == "=>" {
					fmt.Println("Ответ: (-∞;" + fmt.Sprint(start) + "]∪[" + fmt.Sprint(end) + ";+∞)")
				}
			} else if is == "<" || is == "<=" {
				if is == "<" {
					fmt.Println("Ответ: (" + fmt.Sprint(start) + ";" + fmt.Sprint(end) + ")")
				} else if is == "<=" {
					fmt.Println("Ответ: [" + fmt.Sprint(start) + ";" + fmt.Sprint(end) + "]")
				}
			}
		} else if ud == 2 {
			if is == ">" || is == "=>" {
				if is == ">" {
					fmt.Println("Ответ: (" + fmt.Sprint(start) + ";" + fmt.Sprint(end) + ")")
				} else if is == "=>" {
					fmt.Println("Ответ: [" + fmt.Sprint(start) + ";" + fmt.Sprint(end) + "]")
				}
			} else if is == "<" || is == "<=" {
				if is == "<" {
					fmt.Println("Ответ: (-∞;" + fmt.Sprint(start) + ")∪(" + fmt.Sprint(end) + ";+∞)")
				} else if is == "<=" {
					fmt.Println("Ответ: (-∞;" + fmt.Sprint(start) + "]∪[" + fmt.Sprint(end) + ";+∞)")
				}
			}
		}
	} else if oot == 1 {
		if ud == 1 {

			if is == "=>" || is == "<=" {
				if is == "=>" {
					fmt.Println("Ответ: (-∞;+∞)")
				} else if is == "<=" {
					fmt.Println("Ответ: Решений нет(⌀)")
				}
			} else if is == ">" || is == "<" {
				if is == ">" {
					fmt.Println("Ответ: (-∞;" + fmt.Sprint(x) + ")∪(" + fmt.Sprint(x) + ";+∞)")
				} else if is == "<" {
					fmt.Println("Ответ: Решений нет(⌀)")
				}
			}
		} else if ud == 2 {
			if is == "=>" || is == "<=" {
				if is == "=>" {
					fmt.Println("Ответ: Решений нет(⌀)")
				} else if is == "<=" {
					fmt.Println("Ответ: (-∞;+∞)")
				}
			} else if is == ">" || is == "<" {
				if is == ">" {
					fmt.Println("Ответ: Решений нет(⌀)")
				} else if is == "<" {
					fmt.Println("Ответ: (-∞;" + fmt.Sprint(x) + ")∪(" + fmt.Sprint(x) + ";+∞)")
				}
			}
		}
	}
	var wait string
	fmt.Scan(&wait)
}
