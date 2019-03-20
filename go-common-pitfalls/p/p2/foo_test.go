package p2

import (
	"fmt"
	"testing"

	"github.com/garsue/talks/p"
)

func Test2(t *testing.T) {
	p.Lock.Lock()
	p.Lock.Unlock()
	fmt.Println("done in Test2")
}
