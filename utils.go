package ImageGOrruptor

import "math/rand"

func posOrNeg() int {
	num := rand.Intn(2)
	if num == 1 {
		return -1
	}
	return 1
}