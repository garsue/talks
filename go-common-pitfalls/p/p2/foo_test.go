package p2

import (
	"fmt"
	"testing"

	"github.com/garsue/talks/p"
)

func TestName(t *testing.T) {
	p.Slice = append(p.Slice, "p2")
	fmt.Println(p.Slice)
}
