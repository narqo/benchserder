package tokens

import (
	"math/rand"
	"time"
)

const (
	minLength = 6
	maxLength = 7

	minIdLen6 int64 = 60466176    // in base36:  100000
	maxInt32  int64 = 2147483647  // in base36:  zik0zj
	maxIdLen6 int64 = 2176782335  // in base36:  zzzzzz
	minIdLen7 int64 = 2176782336  // in base36: 1000000
	maxIdLen7 int64 = 78364164095 // in base36: zzzzzzz
)

const EmptyTrackerToken = TrackerToken("")

type TrackerToken string

func init() {
	// for RandomTrackerToken()
	rand.Seed(time.Now().UnixNano())
}

func (token TrackerToken) String() string {
	return string(token)
}

func (token TrackerToken) IsEmpty() bool {
	return token == EmptyTrackerToken
}

func (token TrackerToken) IsValid() bool {
	l := len(token)
	return l >= minLength && l <= maxLength
}
