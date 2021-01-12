package lesson3

import (
	"strconv"
 	"go/lesson1"
)

// NewVersion Returns the float version
func NewVersion() float64 {
	if s, err := strconv.ParseFloat(lesson1.MyVersion, 64); err == nil {
		return s
	}
	return 0
}