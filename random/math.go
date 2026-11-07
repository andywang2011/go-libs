package random

import (
	"math/rand"
	"time"
)

// package-local PRNG (not crypto-safe).
var irand = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	upper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower  = "abcdefghijklmnopqrstuvwxyz"
	number = "0123456789"
	char62 = upper + lower + number
)

func Uppers(n int) string {
	return Chars(upper, n)
}

func Lowers(n int) string {
	return Chars(lower, n)
}

func Numbers(n int) string {
	return Chars(number, n)
}

func Strings(n int) string {
	return Chars(char62, n)
}

func Chars(chars string, n int) string {
	length := len(chars)
	buf := make([]byte, n)
	for i := range buf {
		buf[i] = chars[irand.Intn(length)]
	}
	return string(buf)
}

// RangeNum returns a uniform random int in [min, max] (inclusive); min may exceed max.
func RangeNum(min, max int) int {
	if min > max {
		min, max = max, min
	}
	num := irand.Intn(max - min + 1)
	return num + min
}

// Shuffle permutes sli in place using the package PRNG.
func Shuffle[T any](sli []T) {
	irand.Shuffle(len(sli), func(i, j int) {
		sli[i], sli[j] = sli[j], sli[i]
	})
}
