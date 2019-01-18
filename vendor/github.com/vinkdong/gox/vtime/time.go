package vtime

import "time"

func ParserTimestampMs(timestampMs int64) time.Time {
	return time.Unix(0, timestampMs*10e5)
}

func ParserTimestampNs(timestampNs int64) time.Time {
	return time.Unix(0, timestampNs)
}

func ParserTimestampS(timestampS int64) time.Time {
	return time.Unix(timestampS, 0)
}
