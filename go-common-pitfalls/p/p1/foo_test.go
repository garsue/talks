package p1

import (
	"fmt"
	"testing"

	"github.com/garsue/talks/p"
)

func TestName(t *testing.T) {
	p.Slice = append(p.Slice, "p1")
	fmt.Println(p.Slice)
}
