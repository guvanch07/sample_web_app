package main

import (
	"errors"
	"fmt"
	"time"
)

// TimeOfDay описывает время в пределах одного дня
type TimeOfDay struct {
	hour   int
	minute int
	second int
	loc    *time.Location
}

// Hour возвращает часы в пределах дня
func (t TimeOfDay) Hour() int {
	return t.hour
}

// Minute возвращает минуты в пределах часа
func (t TimeOfDay) Minute() int {
	return t.minute
}

// Second возвращает секунды в пределах минуты
func (t TimeOfDay) Second() int {
	return t.second
}

// String возвращает строковое представление времени
// в формате чч:мм:сс TZ (например, 12:34:56 UTC)
func (t TimeOfDay) String() string {
	return fmt.Sprintf("%02d:%02d:%02d %s", t.hour, t.minute, t.second, t.loc)
}

// Equal сравнивает одно время с другим.
// Если у t и other разные локации - возвращает false.
func (t TimeOfDay) Equal(other TimeOfDay) bool {
	return t.loc.String() == other.loc.String() &&
		t.hour == other.hour &&
		t.minute == other.minute &&
		t.second == other.second
}

// Before возвращает true, если время t предшествует other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) Before(other TimeOfDay) (bool, error) {
	if t.loc.String() != other.loc.String() {
		return false, errors.New("different locations")
	}

	if t.hour < other.hour ||
		(t.hour == other.hour && t.minute < other.minute) ||
		(t.hour == other.hour && t.minute == other.minute && t.second < other.second) {
		return true, nil
	}

	return false, nil
}

// After возвращает true, если время t идет после other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) After(other TimeOfDay) (bool, error) {
	if t.loc.String() != other.loc.String() {
		return false, errors.New("different locations")
	}

	if t.hour > other.hour ||
		(t.hour == other.hour && t.minute > other.minute) ||
		(t.hour == other.hour && t.minute == other.minute && t.second > other.second) {
		return true, nil
	}

	return false, nil
}

// MakeTimeOfDay создает время в пределах дня
func MakeTimeOfDay(hour, min, sec int, loc *time.Location) TimeOfDay {
	return TimeOfDay{
		hour:   hour,
		minute: min,
		second: sec,
		loc:    loc,
	}
}

// func main() {
// 	t1 := MakeTimeOfDay(17, 45, 22, time.UTC)
// 	t2 := MakeTimeOfDay(20, 3, 4, time.UTC)

// 	fmt.Println(t1.Hour(), t1.Minute(), t1.Second())
// 	// 17 45 22

// 	fmt.Println(t1)
// 	// 17:45:22 UTC

// 	fmt.Println(t1.Equal(t2))
// 	// false

// 	before, err := t1.Before(t2)
// 	fmt.Println(before, err)
// 	// true <nil>

// 	after, err := t1.After(t2)
// 	fmt.Println(after, err)
// 	// false <nil>
// }
