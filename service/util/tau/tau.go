package tau

import (
	"time"
)

//////////
// Time //
//////////

// Timestamp creates a new Unixtime in milliseconds
func Timestamp() uint64 {
	return uint64(time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)))
}
