package bytestrings

import (
	"github.com/dNaszta/go_cookbook/Ch01-io-interfaces/L02-bytestrings/bytestrings"
	"testing"
)

func TestWorkWithBuffer(t *testing.T) {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		t.Errorf("unexpected error")
	}
}
