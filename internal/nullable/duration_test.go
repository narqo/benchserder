package nullable

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDuration(t *testing.T) {
	assert.Equal(t, Duration{0, true}, NewDuration(0))
	assert.Equal(t, Duration{10, true}, NewDuration(10))
}

func TestMaxDuration(t *testing.T) {
	maxDuration := MaxDuration(
		Duration{0, true},
		Duration{2, true},
		Duration{1, true},
		Duration{3, false},
	)

	assert.Equal(t, Duration{2, true}, maxDuration)
}

func TestDuration_DurationPtr(t *testing.T) {
	duration := Duration{}
	assert.Nil(t, duration.DurationPtr())

	duration = Duration{
		Duration: time.Hour,
	}
	assert.Nil(t, duration.DurationPtr())

	duration = Duration{
		Duration: time.Hour,
		Valid:    true,
	}
	durationPtr := duration.DurationPtr()
	require.NotNil(t, durationPtr)
	assert.Equal(t, time.Hour, *durationPtr)
}

func TestDuration_IsDefined(t *testing.T) {
	duration := Duration{}
	assert.False(t, duration.IsDefined())

	duration = Duration{
		Duration: time.Hour,
	}
	assert.False(t, duration.IsDefined())

	duration = NewDuration(time.Hour)
	assert.True(t, duration.IsDefined())
}
