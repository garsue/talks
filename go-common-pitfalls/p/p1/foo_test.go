package p2

import (
	"fmt"
	"testing"

	"github.com/garsue/talks/p"
)

// START OMIT
func Test1(t *testing.T) {
	p.Lock.Lock()
	p.Lock.Unlock()
	fmt.Println("done in Test1")
}
// END OMIT
