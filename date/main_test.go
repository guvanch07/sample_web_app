package main

import (
	"testing"
	"time"
)

// начало решения

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}

func isLeapYear1(year int) bool {
	t := time.Date(year, 2, 29, 0, 0, 0, 0, time.UTC)
	return t.Month() == time.February
}

// конец решения

func Test(t *testing.T) {
	if !isLeapYear(2020) {
		t.Errorf("2020 is a leap year")
	}
	if isLeapYear(2022) {
		t.Errorf("2022 is NOT a leap year")
	}
}
