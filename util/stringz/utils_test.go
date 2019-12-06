package stringz

import (
	"fmt"
	"testing"
)

func TestPermutations(t *testing.T) {
	Permutations([]string{"a", "b", "c", "d"}, func(val []string) {
		fmt.Printf("%v\n", val)
	})
}
