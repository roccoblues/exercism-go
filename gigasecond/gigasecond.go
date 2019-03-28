package gigasecond

import "time"

const gigaSecond = time.Second * 1e9

// AddGigasecond adds 10^9 seconds to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}
