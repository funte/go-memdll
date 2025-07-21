package memdll

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sum.dll
var data []byte

func Test_memdll(t *testing.T) {
	if d, err := NewDLL(data, "sum.dll"); err != nil {
		t.Error("Failed to create memory dll, err", err)
	} else if p, err := d.FindProc("sum"); err != nil {
		t.Error("Failed to find sum procdure, err", err)
	} else {
		r, _, _ := p.Call(1, 2)
		assert.Equal(t, int(r), 3)
	}
}
