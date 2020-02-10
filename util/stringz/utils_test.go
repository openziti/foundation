package stringz

import (
	"fmt"
	"testing"
)

func TestPermutations(t *testing.T) {
	Permutations([]string{"a", "b", "c", "d"}, func(val []string) {
		if len(val) > 100000 {
			fmt.Printf("%v\n", val)
		}
	})
}
