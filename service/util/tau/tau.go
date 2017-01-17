package tau

import (
	"time"
)

func Timestamp() uint64 {
	return uint64(time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)))
}
