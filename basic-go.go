package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	a = 10
	b = 4
	c = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d, %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
}

func trangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func main() {
	fmt.Println("helloWord!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(a, b, c)
	euler()
	trangle()
}
