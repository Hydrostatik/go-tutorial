package businesslayer

import (
	"fmt"
	"math/rand"
	"time"
)

var generator = rand.New(rand.NewSource(time.Now().UnixNano()))

type Die struct{}

func (x Die) Roll1To(max int) int {
	return generator.Intn(max) + 1
}

func DieCheck() {
	for _, v := range []int{4, 6, 8, 10, 20} {
		fmt.Printf("The result of a %d sided die is: %d\n", v, Die{}.Roll1To(v))
	}
}
