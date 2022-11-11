package utils

import (
	
	"math/rand"
)


func RandomMessage() string {

	var s string
	for i := 0; i < 65; i++ {

		val := rand.Intn(62) + 55;
		if val < 65 {

			val -= 7
		} else if val > 90 {

			val += 6
		}
		s = s + string(val)
	}
	return s
}