package util

import (
	"math/rand"
	"time"
)

func SmallSleep(a, b int) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(a+rand.Intn(b)) * time.Millisecond)
}
