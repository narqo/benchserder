package nullable

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testTimeLayout = "2006-01-02 15:04:05"

func TestNewTime(t *testing.T) {
	now := time.Now().UTC()
	assert.Equal(t, Time{now, true}, NewTime(now))
	assert.Equal(t, Time{}, NewTime(time.Time{}))
}

func TestTime_IsDefined(t *testing.T) {
	now := time.Now().UTC()

	nullTime := Time{}
	assert.False(t, nullTime.IsDefined())

	nullTime = Time{Time: now}
	assert.False(t, nullTime.IsDefined())

	nullTime = NewTime(now)
	assert.True(t, nullTime.IsDefined())
}

func TestTime_Format_Valid(t *testing.T) {
	now := time.Now().UTC()
	nullTime := NewTime(now)
	nullString := nullTime.Format(testTimeLayout)
	assert.Equal(t, nullString.Valid, true)
	assert.Equal(t, nullString.String, now.Format(testTimeLayout))
}

func TestTime_Format_Invalid(t *testing.T) {
	nullTime := NewTime(time.Time{})
	nullString := nullTime.Format(testTimeLayout)
	assert.Equal(t, nullString.Valid, false)
	assert.Equal(t, nullString.String, "")
}
