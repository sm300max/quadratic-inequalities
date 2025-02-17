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

	fmt.Println("Enter a")
	fmt.Scan(&a)

	fmt.Println("Enter b:")
	fmt.Scan(&b)

	fmt.Println("Enter c:")
	fmt.Scan(&c)

	fmt.Println("Choose [</<=/=/=>/>]:")
	fmt.Scan(&is)

	if a == 0 {
		fmt.Println("A square expression cannot contain a coefficient a equal to 0")
		return
	}

	if is != "<" && is != "<=" && is != "=" && is != "=>" && is != ">" {
		fmt.Println("A mistake was made in choosing the sign, try again.")
		return
	}

	var q string

	fmt.Println(fmt.Sprint(a) + "x²" + "+" + fmt.Sprint(b) + "x" + "+" + fmt.Sprint(c) + is + "0")
	fmt.Println("The expression is correct? [Y/n]")
	fmt.Scanln(&q)

	if q != "y" && q != ""{
		fmt.Println("Try again")
		return
	}
	var ud int8
	fmt.Println("In view of the function:")
	fmt.Println("y=" + fmt.Sprint(a) + "x²" + "+" + fmt.Sprint(b) + "x" + "+" + fmt.Sprint(c))
	fmt.Println("The square function")
	fmt.Println("The graph is a parabola")
	if a > 0 {
		ud = 1
		fmt.Println("Branches up")
	} else {
		ud = 2
		fmt.Println("Branches down")
	}

	fmt.Println("Function zeros: " + fmt.Sprint(a) + "x²" + "+" + fmt.Sprint(b) + "x" + "+" + fmt.Sprint(c) + "=0")

	var x1, x2 float64
	var x float64
	var oot int8

	D := b*b - 4*a*c
	if D < 0 {
		fmt.Println("D = b² - 4ac")
		fmt.Println("D = " + fmt.Sprint(b) + "² - 4*" + fmt.Sprint(a) + "*" + fmt.Sprint(c))
		fmt.Println("D = " + fmt.Sprint(D))
		fmt.Println("D < 0")
		fmt.Println("The expression has no roots")
		fmt.Println("It's decided")
	} else if D == 0 {
		oot = 1
		x = -b / (2 * a)
		fmt.Println("D = b² - 4ac")
		fmt.Println("D = " + fmt.Sprint(b) + "² - 4*" + fmt.Sprint(a) + "*" + fmt.Sprint(c))
		fmt.Println("D = " + fmt.Sprint(D))
		fmt.Println("D = 0")

		fmt.Println("The expression has one root")
		fmt.Println("x = -b / 2a")
		fmt.Println("x = -" + fmt.Sprint(b) + "/2" + fmt.Sprint(a))
		fmt.Println("x = " + fmt.Sprint(x))
		fmt.Println("It's decided")
	} else if D > 0 {
		oot = 2
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println("D = b² - 4ac")
		fmt.Println("D = " + fmt.Sprint(b) + "² - 4*" + fmt.Sprint(a) + "*" + fmt.Sprint(c))
		fmt.Println("D = " + fmt.Sprint(D))
		fmt.Println("D > 0")
		fmt.Println("The expression has two roots")
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
					fmt.Println("Answer: (-∞;" + fmt.Sprint(start) + ")∪(" + fmt.Sprint(end) + ";+∞)")
				} else if is == "=>" {
					fmt.Println("Answer: (-∞;" + fmt.Sprint(start) + "]∪[" + fmt.Sprint(end) + ";+∞)")
				}
			} else if is == "<" || is == "<=" {
				if is == "<" {
					fmt.Println("Answer: (" + fmt.Sprint(start) + ";" + fmt.Sprint(end) + ")")
				} else if is == "<=" {
					fmt.Println("Answer: [" + fmt.Sprint(start) + ";" + fmt.Sprint(end) + "]")
				}
			}
		} else if ud == 2 {
			if is == ">" || is == "=>" {
				if is == ">" {
					fmt.Println("Answer: (" + fmt.Sprint(start) + ";" + fmt.Sprint(end) + ")")
				} else if is == "=>" {
					fmt.Println("Answer: [" + fmt.Sprint(start) + ";" + fmt.Sprint(end) + "]")
				}
			} else if is == "<" || is == "<=" {
				if is == "<" {
					fmt.Println("Answer: (-∞;" + fmt.Sprint(start) + ")∪(" + fmt.Sprint(end) + ";+∞)")
				} else if is == "<=" {
					fmt.Println("Answer: (-∞;" + fmt.Sprint(start) + "]∪[" + fmt.Sprint(end) + ";+∞)")
				}
			}
		}
	} else if oot == 1 {
		if ud == 1 {

			if is == "=>" || is == "<=" {
				if is == "=>" {
					fmt.Println("Answer: (-∞;+∞)")
				} else if is == "<=" {
					fmt.Println("Answer: There are no solutions(⌀)")
				}
			} else if is == ">" || is == "<" {
				if is == ">" {
					fmt.Println("Answer: (-∞;" + fmt.Sprint(x) + ")∪(" + fmt.Sprint(x) + ";+∞)")
				} else if is == "<" {
					fmt.Println("Answer: There are no solutions(⌀)")
				}
			}
		} else if ud == 2 {
			if is == "=>" || is == "<=" {
				if is == "=>" {
					fmt.Println("Answer: There are no solutions(⌀)")
				} else if is == "<=" {
					fmt.Println("Answer: (-∞;+∞)")
				}
			} else if is == ">" || is == "<" {
				if is == ">" {
					fmt.Println("Answer: There are no solutions(⌀)")
				} else if is == "<" {
					fmt.Println("Answer: (-∞;" + fmt.Sprint(x) + ")∪(" + fmt.Sprint(x) + ";+∞)")
				}
			}
		}
	}
	var wait string
	fmt.Scan(&wait)
}
