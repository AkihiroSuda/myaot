package signext

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestSignExt(t *testing.T) {
	assert.Equal(t, 305419896, SignExt(0x12345678, 2))
	assert.Equal(t, -8, SignExt(0x12345678, 4))
}
