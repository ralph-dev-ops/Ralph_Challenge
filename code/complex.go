package main

import (
	"fmt"
	"math/cmplx"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func printComplexLn(number complex128) {
	sign := "+"
	if imag(number) < 0 {
		sign = ""
	}

	fmt.Printf("%.2f%s%.2fi\n", real(number), sign, imag(number))
}

func stringToComplex(input string) complex128 {
	parts := strings.Fields(input)

	a, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	b, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return complex(a, b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("First Complex: ")
	first, _ := reader.ReadString('\n')

	c := stringToComplex(first)

	fmt.Print("Second Complex: ")
	second, _ := reader.ReadString('\n')

	d := stringToComplex(second)

	printComplexLn(c + d)
	printComplexLn(c - d)
	printComplexLn(c * d)
	printComplexLn(c / d)
	printComplexLn(complex(cmplx.Abs(c), 0))
	printComplexLn(complex(cmplx.Abs(d), 0))
}
