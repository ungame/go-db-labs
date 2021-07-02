package util

import (
	"math/rand"
	"time"
)

func Random() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func IsDateEqual(a, b time.Time) bool {
	return a.Year() == b.Year() && a.Month() == b.Month() && a.Day() == b.Day()
}

func IsTimeEqual(a, b time.Time) bool {
	return a.Hour() == b.Hour() && a.Minute() == b.Minute() && a.Second() == b.Second()
}

func IsDateTimeEqual(a, b time.Time) bool {
	return IsDateEqual(a, b) && IsTimeEqual(a, b)
}