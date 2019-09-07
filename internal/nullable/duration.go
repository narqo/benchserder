package nullable

import "time"

type Duration struct {
	Duration time.Duration
	Valid    bool
}

func NewDuration(duration time.Duration) Duration {
	return Duration{
		Duration: duration,
		Valid:    true,
	}
}

func MaxDuration(durations ...Duration) (maxDuration Duration) {
	for _, duration := range durations {
		if duration.Valid && (!maxDuration.Valid || maxDuration.Duration < duration.Duration) {
			maxDuration = duration
		}
	}
	return maxDuration
}

func (v Duration) DurationPtr() *time.Duration {
	if v.Valid {
		return &v.Duration
	}

	return nil
}

func (v Duration) IsDefined() bool {
	return v.Valid
}
