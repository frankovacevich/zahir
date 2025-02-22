package player

import "time"

func now() int64 {
	return time.Now().UnixMilli()
}

func waitStep(duration float64, ts int64) int64 {
	sleepMilliseconds := duration*1000 - float64(now()-ts)
	if sleepMilliseconds > 0 {
		time.Sleep(time.Duration(sleepMilliseconds) * time.Millisecond)
	}
	return now()
}
