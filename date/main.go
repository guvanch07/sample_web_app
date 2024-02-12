package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func main() {
	//MainFormat()
	str := asLegacyDate1(time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC))
	fmt.Println(str)
	t, _ := parseLegacyDate1("3600.123456789")
	fmt.Println(t)
}
func asLegacyDate1(t time.Time) string {

	precision := len(strconv.Itoa(t.Nanosecond()))

	format := fmt.Sprintf("%%.%df", precision)
	return fmt.Sprintf(format, float64(t.UnixNano())/1e9)
}

func parseLegacyDate1(d string) (time.Time, error) {
	_, err := fmt.Sscanf(d, "%f", new(float64))
	if err != nil {
		return time.Time{}, errors.New("error parsing legacy date")
	}
	secs, err := time.Parse(time.RFC3339, d)
	if err != nil {
		return time.Time{}, errors.New("error parsing legacy date")
	}

	// Split seconds into integer and fractional parts without using math.Floor
	intPart := int64(secs.Second())
	fracPart := int64((float64(secs.Second()) - float64(intPart)) * 1e9)

	t := time.Unix(intPart, fracPart)
	return t, nil
}
