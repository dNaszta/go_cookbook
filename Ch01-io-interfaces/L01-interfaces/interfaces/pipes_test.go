package interfaces

import (
	"github.com/dNaszta/go_cookbook/Ch01-io-interfaces/L01-interfaces/interfaces"
	"testing"
)

func TestPipeExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interfaces.PipeExample()
		})
	}
}
