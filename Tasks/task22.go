package main

import (
	"fmt"
	"math/big"
)

func main() {

	// работа с библиотекой math/big

	bignum1, _ := new(big.Int).SetString("1000000000000", 10)
	bignum2, _ := new(big.Int).SetString("3", 10)

	fmt.Println("Сложение: ", сложение(bignum1, bignum2).Text(10))
	fmt.Println("Вычитание: ", вычитание(bignum1, bignum2).Text(10))
	fmt.Println("Умножение: ", умножение(bignum1, bignum2).Text(10))
	fmt.Println("Деление: ", деление(bignum1, bignum2).Text(10))

}

func сложение(a *big.Int, b *big.Int) *big.Int {
	out := new(big.Int)
	return out.Add(a, b)
}

func вычитание(a *big.Int, b *big.Int) *big.Int {
	out := new(big.Int)
	return out.Sub(a, b)
}

func умножение(a *big.Int, b *big.Int) *big.Int {
	out := new(big.Int)
	return out.Mul(a, b)
}

func деление(num1 *big.Int, num2 *big.Int) *big.Int {
	out := new(big.Int)
	return out.Div(num1, num2)
}
