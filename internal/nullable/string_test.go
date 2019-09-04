package nullable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewString(t *testing.T) {
	assert.Equal(t, String{"", false}, String{}) // invalid

	assert.Equal(t, String{"val", true}, NewString("val")) // valid
}

func TestString_IsDefined(t *testing.T) {
	s := String{}
	assert.False(t, s.IsDefined())

	s = String{String: "val"}
	assert.False(t, s.IsDefined())

	s = NewString("val")
	assert.True(t, s.IsDefined())
}
