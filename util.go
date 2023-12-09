package main

import (
	"strconv"
	"strings"
)

func Lines(s string) []string {
	return strings.Split(s, "\n")
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func All[T any](tt []T, f func(T) bool) bool {
	for _, t := range tt {
		if !f(t) {
			return false
		}
	}
	return true
}

func Any[T any](tt []T, f func(T) bool) bool {
	for _, t := range tt {
		if f(t) {
			return true
		}
	}
	return false
}

func Map[T any, U any](t []T, f func(T) U) []U {
	var u []U
	for _, tt := range t {
		u = append(u, f(tt))
	}
	return u
}
func MapIndex[T any, U any](t []T, f func(T, int) U) []U {
	var u []U
	for i, tt := range t {
		u = append(u, f(tt, i))
	}
	return u
}

func SumInt(ii []int) int {
	var t int
	for _, i := range ii {
		t += i
	}
	return t
}

func Trim(s string) string {
	return strings.Trim(s, " ")
}

func MustInt(s string) int {
	return Must(strconv.Atoi(s))
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Filter[T any](tt []T, f func(T) bool) []T {
	var filtered []T
	for _, t := range tt {
		if f(t) {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

func Find[T any](items []T, f func(T) bool) (T, bool) {
	for _, e := range items {
		if f(e) {
			return e, true
		}
	}
	var def T
	return def, false
}
func FindAll[T any](items []T, f func(T) bool) []T {
	var all []T
	for _, e := range items {
		if f(e) {
			all = append(all, e)
		}
	}
	return all
}

func IndexAll[T any](items []T, f func(T) bool) []int {
	var indexes []int
	for i, e := range items {
		if f(e) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
