package nullable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInt(t *testing.T) {
	assert.Equal(t, Int{2, true}, NewInt(2))
	assert.Equal(t, Int{-1, true}, NewInt(-1))
}

func TestLegacyInt(t *testing.T) {
	assert.Equal(t, Int{-1, false}, LegacyInt(-1))
	assert.Equal(t, Int{-2, true}, LegacyInt(-2))
	assert.Equal(t, Int{0, true}, LegacyInt(0))
	assert.Equal(t, Int{10, true}, LegacyInt(10))
}

func TestInt_IsDefined(t *testing.T) {
	nullInt := Int{}
	assert.False(t, nullInt.IsDefined())

	nullInt = Int{
		Int: 12,
	}
	assert.False(t, nullInt.IsDefined())

	nullInt = NewInt(1)
	assert.True(t, nullInt.IsDefined())

	nullInt = NewInt(-1)
	assert.True(t, nullInt.IsDefined())

	nullInt = LegacyInt(-1)
	assert.False(t, nullInt.IsDefined())

	nullInt = LegacyInt(1)
	assert.True(t, nullInt.IsDefined())
}

func TestInt_String(t *testing.T) {
	nullInt := Int{}
	str := nullInt.String()
	assert.Equal(t, "", str)

	nullInt = Int{
		Int:   12,
		Valid: true,
	}
	str = nullInt.String()
	assert.Equal(t, "12", str)
}
