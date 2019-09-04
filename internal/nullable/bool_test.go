package nullable

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBool(t *testing.T) {
	assert.Equal(t, Bool{false, true}, NewBool(false))
	assert.Equal(t, Bool{true, true}, NewBool(true))

	assert.Equal(t, Bool{false, true}, False)
	assert.Equal(t, Bool{true, true}, True)
}

func TestBool_BoolPtr(t *testing.T) {
	nullBool := Bool{true, true}
	require.NotNil(t, nullBool.BoolPtr())
	assert.Equal(t, true, *nullBool.BoolPtr())

	nullBool = Bool{false, true}
	require.NotNil(t, nullBool.BoolPtr())
	assert.Equal(t, false, *nullBool.BoolPtr())

	nullBool = Bool{false, false}
	assert.Nil(t, nullBool.BoolPtr())

	nullBool = Bool{true, false}
	assert.Nil(t, nullBool.BoolPtr())
}

func TestBool_SqlNullBool(t *testing.T) {
	cases := []Bool{
		{Valid: false, Bool: false},
		{Valid: false, Bool: true},
		{Valid: true, Bool: false},
		{Valid: true, Bool: true},
	}

	for _, nullBool := range cases {
		t.Run(nullBool.String(), func(t *testing.T) {
			sqlNullBool := nullBool.SqlNullBool()
			require.IsType(t, sql.NullBool{}, sqlNullBool)
			assert.Equal(t, nullBool.Valid, sqlNullBool.Valid)
			assert.Equal(t, nullBool.Bool, sqlNullBool.Bool)
		})
	}
}

func TestBool_IsDefined(t *testing.T) {
	cases := []struct {
		val     Bool
		defined bool
	}{
		{Bool{Valid: false, Bool: false}, false},
		{Bool{Valid: false, Bool: true}, false},
		{Bool{Valid: true, Bool: false}, true},
		{Bool{Valid: true, Bool: true}, true},
		{NewBool(true), true},
		{NewBool(false), true},
		{True, true},
		{False, true},
	}

	for _, tc := range cases {
		t.Run(tc.val.String(), func(t *testing.T) {
			assert.Equal(t, tc.defined, tc.val.IsDefined())
		})
	}
}

func TestBool_IsTrue(t *testing.T) {
	cases := []struct {
		val    Bool
		isTrue bool
	}{
		{Bool{Valid: false, Bool: false}, false},
		{Bool{Valid: false, Bool: true}, false},
		{Bool{Valid: true, Bool: false}, false},
		{Bool{Valid: true, Bool: true}, true},
	}

	for _, tc := range cases {
		t.Run(tc.val.String(), func(t *testing.T) {
			assert.Equal(t, tc.isTrue, tc.val.IsTrue())
		})
	}
}
