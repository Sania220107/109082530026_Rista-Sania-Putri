package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	var i int = 0
	var status bool = false

	for i < n && !status {
		status = T[i] == val
		i++
	}

	return status
}

func inputSet(T *set, n *int) {
	*n = 0
	var bilangan int

	fmt.Scan(&bilangan)

	for *n < 2022 && !exist(*T, *n, bilangan) {
		T[*n] = bilangan
		*n++

		fmt.Scan(&bilangan)
	}
}

func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	var j int = 0
	*h = 0

	for j < n {
		if exist(T2, m, T1[j]) {
			T3[*h] = T1[j]
			*h++
		}
		j++
	}
}

func printSet(T set, n int) {
	var i int = 0

	for i < n {
		fmt.Print(T[i], " ")
		i++
	}
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int

	inputSet(&s1, &n1)
	inputSet(&s2, &n2)

	findIntersection(s1, s2, n1, n2, &s3, &n3)

	printSet(s3, n3)
}