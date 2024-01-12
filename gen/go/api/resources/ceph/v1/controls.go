package cephv1

import "time"

func (x *OSDScrubbingSchedule) Default() {
	if x.MaxScrubOps == nil {
		maxScrubOps := int64(3)
		x.MaxScrubOps = &maxScrubOps
	}

	zero := int64(0)
	if x.BeginHour == nil {
		x.BeginHour = &zero
	}
	if x.EndHour == nil {
		x.EndHour = &zero
	}
	if x.BeginWeekDay == nil {
		x.BeginWeekDay = &zero
	}
	if x.EndWeekDay == nil {
		x.EndWeekDay = &zero
	}

	if x.MinScrubInterval == nil {
		minScrubInterval := 24 * time.Hour
		minScrubIntervalSecs := int64(minScrubInterval.Seconds())
		x.MinScrubInterval = &minScrubIntervalSecs
	}

	sevenDaysDuration := 7 * 24 * time.Hour
	sevenDaysDurationSecs := int64(sevenDaysDuration.Seconds())
	if x.MaxScrubInterval == nil {
		x.MaxScrubInterval = &sevenDaysDurationSecs
	}
	if x.DeepScrubInterval == nil {
		x.DeepScrubInterval = &sevenDaysDurationSecs
	}

	zeroFloat := float32(0.0)
	if x.ScrubSleepSeconds == nil {
		x.ScrubSleepSeconds = &zeroFloat
	}
}
