package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type (
	mapFunc[E any]    func(E) E
	keepFunc[E any]   func(E) bool
	reduceFunc[E any] func(cur, next E) E
)

func main() {
	s := []string{"a", "b", "c"}

	n := []int{1, 2, 3, 4}

	fmt.Println(Map(s, strings.ToUpper))

	evenNumbers := Filter(n, isEven[int])

	oddNumbers := Filter(n, isOdd[int])

	sum := Reduce(n, 0, func(cur, next int) int {
		return cur + next
	})

	p := Reduce(n, 1, func(cur, next int) int { return cur * next })

	j := Reduce(s, "", func(cur, next string) string { return cur + next })

	fmt.Println(evenNumbers)

	fmt.Println(oddNumbers)

	fmt.Println(sum)

	fmt.Println(p)

	fmt.Println(j)
}

func Map[S ~[]E, E any](s S, f mapFunc[E]) S {
	result := make(S, len(s))

	for i := range s {
		result[i] = f(s[i])
	}

	return result
}

func Filter[S ~[]E, E any](s S, f keepFunc[E]) S {
	result := S{}

	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func Reduce[E any](s []E, init E, f reduceFunc[E]) E {
	cur := init

	for _, v := range s {
		cur = f(cur, v)
	}

	return cur
}

func isEven[T constraints.Integer](v T) bool {
	return v%2 == 0
}

func isOdd[T constraints.Integer](v T) bool {
	return v%2 > 0
}
