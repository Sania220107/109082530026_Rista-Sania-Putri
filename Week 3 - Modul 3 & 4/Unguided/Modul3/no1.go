package main

import "fmt"

func factorial(n int) int {
	var hasil int
	hasil = 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}

	return hasil
}

func permutation(n, r int) int {
	var perm int

	perm = factorial(n) / factorial(n-r)
	return perm
}

func combination(n, r int) int {
	var comb int

	comb = factorial(n) / (factorial(r) * factorial(n-r))
	return comb
}

func main() {
	var a, b, c, d int
	var perm1, com1, perm2, com2 int

	fmt.Scan(&a, &b, &c, &d) 

	perm1 = permutation(a, c)
	com1 = combination(a, c)

	perm2 = permutation(b, d)
	com2 = combination(b, d)

	fmt.Println(perm1, com1)
	fmt.Println(perm2, com2)

}