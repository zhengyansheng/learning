package schedule

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSelectHost(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Println(rand.Intn(2))
	}
}
